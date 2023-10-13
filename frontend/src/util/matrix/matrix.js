import {DFS} from "./dfs";
import {BFS} from "./bfs";
import {parseHeading} from "./toc";

class Matrix {
    constructor(root, nodes, nodeRelations, dfs = false) {
        let fs
        if (dfs) {
            fs = new DFS(root, nodeRelations)
        } else {
            fs = new BFS(root, nodeRelations)
        }
        this.nodes = nodes
        this.nodeRelations = nodeRelations
        this.matrix = fs.matrix
        this.x = fs.x
        this.y = fs.y
        this.toc = {
            headings: [],
            title: null
        }
    }

    shift() {
        for (let x = 2; x <= this.x - 1; x++) {
            let start = -1
            let end = -1
            for (let y = 0; y <= this.y - 1; y++) {
                // find continued lines
                if (this.matrix[x][y] > 0) {
                    if (start !== -1) {
                        end = y
                    } else {
                        start = y
                        end = start
                    }
                }
                if (this.matrix[x][y] === 0 || y === this.y - 1) {
                    if (start !== -1) {
                        // shift continued lines
                        this.shiftLine(x, start, end)
                    }
                    start = -1
                    end = -1
                }
            }
        }
    }

    shiftLine(x, start, end) {
        // calculate shift distance
        let shift = 0
        for (let y = start; y <= end; y++) {
            let col = x
            // find border
            while (col > 0) {
                col--
                if (this.matrix[col][y] !== 0) {
                    col += 1
                    break
                }
            }
            if (col === x) {
                return
            }
            if (shift === 0 || shift > x - col) {
                shift = x - col
            }
        }
        if (shift > 0) {
            for (let y = start; y <= end; y++) {
                this.matrix[x - shift][y] = this.matrix[x][y]
                this.matrix[x][y] = 0
            }
        }
    }

    cleanSuffix() {
        for (let y = 0; y <= this.y - 1; y++) {
            for (let x = this.x - 1; x >= 0; x--) {
                if (this.matrix[x][y] > 0) {
                    break
                }
                this.matrix[x][y] = -1
            }
        }
    }

    print() {
        for (let y = 0; y <= this.y - 1; y++) {
            for (let x = 0; x <= this.x - 1; x++) {
                process.stdout.write(this.matrix[x][y].toString())
            }
            process.stdout.write("\n")
        }
    }

    dumpNode() {
        let matrix = [[]]
        for (let y = 0; y <= this.y - 1; y++) {
            let nodeCol = []
            for (let x = 0; x <= this.x - 1; x++) {
                let id = this.matrix[x][y]
                let node = this.nodes[id]
                if (id > 0) {
                    let nodeRelation = this.nodeRelations[id]
                    node.child = nodeRelation.child
                    node.next = nodeRelation.next
                    if (node.child > 0) {
                        let child = this.nodes[node.child]
                        child.parent = id
                    }
                    if (node.next > 0) {
                        let next = this.nodes[node.next]
                        next.last = id
                    }
                    nodeCol.push(node)
                } else {
                    if (id === 0) {
                        nodeCol.push({ID: id, markdown: ''})
                    }
                }
            }
            matrix.push(nodeCol)
        }
        return matrix
    }

    generateToc(nodeId) {
        if (nodeId === 0) {
            return
        }
        let node = this.nodes[nodeId]
        let parsed = parseHeading(node.markdown)
        if (parsed.title !== null) {
            this.toc.title = {
                nodeName: parsed.title.nodeName,
                innerText: parsed.title.innerText,
                id: parsed.title.id,
            }
        }
        parsed.headings.forEach(heading => {
            let h = {
                nodeName: heading.nodeName,
                innerText: heading.innerText,
                id: heading.id,
            }
            this.toc.headings.push(h)
        })
        let nodeRelation = this.nodeRelations[nodeId]
        if (nodeRelation.child > 0) {
            this.generateToc(nodeRelation.child)
        }
        if (nodeRelation.next > 0) {
            this.generateToc(nodeRelation.next)
        }
    }

    /*
    * return caller's id
    * type1:
    *   caller -> callee
    * type2:
    *   caller -> node1
    *               |
    *              ...
    *               |
    *             callee
    */
    caller(id) {
        let node = this.nodes[id]
        if (id > 0) {
            if (node.parent > 0) {
                return node.parent
            }
            if (node.last > 0) {
                return this.caller(node.last)
            }
        } else {
            return undefined
        }
        return undefined
    }

    count() {
        let c = 0
        for (let i = 0; i < this.x; i++) {
            c += this.matrix[i].filter(x => x > 0).length
        }
        return c
    }
}

export {Matrix}
