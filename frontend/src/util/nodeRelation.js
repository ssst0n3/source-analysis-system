import consts from "@/util/const"
import lightweightRestful from "vue-lightweight_restful";

async function create_node_relation(toaster, rootId, nodeId, nextId, childId) {
    nextId = nextId === undefined ? 0 : nextId
    childId = childId === undefined ? 0 : childId
    await lightweightRestful.api.post(consts.api.v1.node_relation.node_relation, null, {
        root: rootId,
        node: nodeId,
        next: nextId,
        child: childId,
    }, {
        caller: toaster,
    })
}

async function update(toaster, baseId, data) {
    await lightweightRestful.api.put(
        consts.api.v1.node_relation.update_node_relation_by_node(baseId),
        null,
        data,
        {
            caller: toaster,
        }
    )
}

async function update_node_relation_insert(toaster, baseNode, nodeId) {
    let parent = baseNode.parent
    let last = baseNode.last
    if (parent !== undefined) {
        await update(toaster, parent, {child: nodeId})
    }
    if (last !== undefined) {
        await update(toaster, last, {next: nodeId})
    }
}

async function update_node_relation(toaster, direction, rootId, baseNode, nodeId) {
    switch (direction) {
        case consts.directions.right:
            await create_node_relation(toaster, rootId, nodeId)
            await update(toaster, baseNode.ID, {child: nodeId})
            break
        case consts.directions.down:
            await create_node_relation(toaster, rootId, nodeId)
            await update(toaster, baseNode.ID, {next: nodeId})
            break
        case consts.directions.up:
            await create_node_relation(toaster, rootId, nodeId, baseNode.ID)
            await update_node_relation_insert(toaster, baseNode, nodeId)
            break
        case consts.directions.left:
            await create_node_relation(toaster, rootId, nodeId, 0, baseNode.ID)
            await update_node_relation_insert(toaster, baseNode, nodeId)
            break
    }
}

export {update, update_node_relation}
