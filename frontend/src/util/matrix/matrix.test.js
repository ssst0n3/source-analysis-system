import {Matrix} from "./matrix";
import {MapNodeRelations} from "./bfs.test";

test('union', () => {
    let matrix = new Matrix(1, [], MapNodeRelations, true)
    matrix.print()
    matrix.shift()
    matrix.print()
    matrix.cleanSuffix()
    matrix.print()
});