class Matrix {
    constructor(root, nodeRelations) {
        this.x = 1;
        this.y = 1;
        this.matrix = [[root]]
        this.nodeRelations = nodeRelations
        this.nextRecursive(root)
    }

    nextRecursive(id) {
        if (id === 0) {
            return
        }
        let node = this.nodeRelations[id]
        if (node.next === 0) {
            for (let y = this.matrix[this.x - 1].length; y < this.y; y++) {
                this.matrix[this.x - 1].push(0)
            }
            return
        }
        this.matrix[this.matrix.length - 1].push(node.next)
        let newLen = this.matrix[this.matrix.length - 1].length
        if (newLen > this.y) {
            for (let x = 0; x < this.x - 1; x++) {
                this.matrix[x].push(0)
            }
            this.y = newLen
        }
        this.nextRecursive(node.next)
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

    dumpNode(nodes) {
        let matrix = [[]]
        for (let y = 0; y <= this.y - 1; y++) {
            let nodeCol = []
            for (let x = 0; x <= this.x - 1; x++) {
                let id = this.matrix[x][y]
                let node = nodes[id]
                if (id > 0) {
                    let nodeRelation = this.nodeRelations[id]
                    node.child = nodeRelation.child
                    node.next = nodeRelation.next
                } else {
                    node = {ID:0, markdown: ''}
                }
                nodeCol.push(node)
            }
            matrix.push(nodeCol)
        }
        return matrix
    }
}

export default {
    Matrix: Matrix
}
