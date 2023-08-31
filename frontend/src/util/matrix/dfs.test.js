import {MapNodeRelations} from './bfs.test'
import {DFS} from "./dfs";

test('dfs', () => {
    let matrix = new DFS(1,  MapNodeRelations)
    matrix.print()
});