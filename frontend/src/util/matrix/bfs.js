class BFS {
    constructor(root, nodeRelations) {
        this.x = 1
        this.y = 1
        this.matrix = [[root]]
        this.nodeRelations = nodeRelations
        this.nextRecursive(root)
        this.childRecursive(0)
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
}

export {BFS}