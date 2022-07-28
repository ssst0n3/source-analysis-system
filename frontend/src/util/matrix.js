import {marked} from 'marked'

function find_headings(doc) {
    let headings = []
    let tag_names = {
        h1: 1,
        h2: 1,
        h3: 1,
        h4: 1,
        h5: 1,
        h6: 1
    }

    function walk(root) {
        if (root.nodeType === 1 && root.nodeName !== 'script') {
            if (Object.prototype.hasOwnProperty.call(tag_names, root.nodeName.toLowerCase())) {
                if (root.getAttribute("class") === null) {
                    headings.push(root);
                }
            } else {
                for (let i = 0; i < root.childNodes.length; i++) {
                    walk(root.childNodes[i]);
                }
            }
        }
    }

    walk(doc);
    return headings
}


function parseHeading(markdown) {
    let rendered = marked(markdown)
    let div = document.createElement('div')
    div.innerHTML = rendered
    return find_headings(div)
}

class Matrix {
    constructor(root, nodes, nodeRelations) {
        this.x = 1
        this.y = 1
        this.toc = []
        this.matrix = [[root]]
        this.nodes = nodes
        this.nodeRelations = nodeRelations
        this.nextRecursive(root)
    }

    nextRecursive(id) {
        if (id === 0) {
            return
        }
        let nodeRelation = this.nodeRelations[id]
        if (nodeRelation.next === 0) {
            for (let y = this.matrix[this.x - 1].length; y < this.y; y++) {
                this.matrix[this.x - 1].push(0)
            }
            return
        }
        this.matrix[this.matrix.length - 1].push(nodeRelation.next)
        let newLen = this.matrix[this.matrix.length - 1].length
        if (newLen > this.y) {
            for (let x = 0; x < this.x - 1; x++) {
                this.matrix[x].push(0)
            }
            this.y = newLen
        }
        this.nextRecursive(nodeRelation.next)
    }

    childRecursive(x) {
        for (let y = this.y - 1; y >= 0; y--) {
            let id = this.matrix[x][y]
            if (id === 0) {
                continue
            }
            let node = this.nodeRelations[id]
            if (node.child === 0) {
                continue
            }
            let col = new Array(y).fill(0)
            col.push(node.child)
            this.matrix.push(col)
            this.x += 1
            this.nextRecursive(node.child)
            this.childRecursive(this.x - 1)
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
                } else {
                    node = {ID: 0, markdown: ''}
                }
                nodeCol.push(node)
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
        let headings = parseHeading(node.markdown)
        headings.forEach(heading => {
            let h = {
                nodeName: heading.nodeName,
                innerText: heading.innerText,
                id: heading.id,
            }
            this.toc.push(h)
        })
        let nodeRelation = this.nodeRelations[nodeId]
        if (nodeRelation.child > 0) {
            this.generateToc(nodeRelation.child)
        }
        if (nodeRelation.next > 0) {
            this.generateToc(nodeRelation.next)
        }
    }
}

export default {
    Matrix: Matrix
}
