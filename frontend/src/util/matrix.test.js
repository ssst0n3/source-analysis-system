import Matrix from './matrix'

let MapNodeRelations = {
    2: {
        ID: 2,
        node: 2,
        root: 1,
        next: 4,
        child: 3
    },
    3: {
        ID: 3,
        node: 3,
        root: 1,
        next: 0,
        child: 0,
    },
    4: {
        ID: 4,
        node: 4,
        root: 1,
        next: 0,
        child: 5,
    },
    1: {
        ID: 1,
        node: 1,
        root: 0,
        next: 0,
        child: 2,
    },
    5: {
        ID: 5,
        node: 5,
        root: 1,
        next: 0,
        child: 0
    }
}

// eslint-disable-next-line no-undef
test('next recursive', () => {
    let matrix = new Matrix.Matrix(1, MapNodeRelations)
    matrix.childRecursive(0)
    matrix.print()
});
