class DFS {
    constructor(root, nodeRelations) {
        this.x = 1
        this.y = 1
        this.matrix = [[root]]
        this.nodeRelations = nodeRelations
        this.childRecursive(root)
        this.nextRecursive(0)
    }

    childRecursive(id) {
        if (id === 0) {
            return
        }
        let nodeRelation = this.nodeRelations[id]
        if (nodeRelation.child === 0) {
            return
        }
        let col = new Array(this.y).fill(0)
        col[this.y-1] = nodeRelation.child
        this.x += 1
        this.matrix.push(col)
        this.childRecursive(nodeRelation.child)
    }

    nextRecursive(y) {
        for (let x = this.x - 1; x >= 0; x--) {
            let id = this.matrix[x][y]
            if (id === 0) {
                continue
            }
            let node = this.nodeRelations[id]
            if (node.next === 0) {
                continue
            }
            for (let i = 0; i < this.x; i ++) {
                this.matrix[i][this.y] = 0
            }
            this.matrix[x][this.y] = node.next
            this.y += 1
            this.childRecursive(node.next)
            this.nextRecursive(this.y - 1)
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

export {
    DFS
}
