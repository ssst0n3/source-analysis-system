import consts from "@/util/const"
import lightweightRestful from "vue-lightweight_restful";

async function create_node_relation(toaster, root, node, next, child) {
    next = next === undefined ? 0 : next
    child = child === undefined ? 0 : child
    await lightweightRestful.api.post(consts.api.v1.node_relation.node_relation, null, {
        root: root,
        node: node,
        next: next,
        child: child,
    }, {
        caller: toaster,
    })
}

async function update(toaster, base, data) {
    await lightweightRestful.api.put(
        consts.api.v1.node_relation.update_node_relation_by_node(base),
        null,
        data,
        {
            caller: toaster,
        }
    )
}

async function update_node_relation_insert(toaster, root, base, node, nodesMap) {
    let baseNode = nodesMap[base]
    let parent = baseNode.parent
    let last = baseNode.last
    if (parent !== undefined) {
        await update(toaster, parent, {child: node})
    }
    if (last !== undefined) {
        await update(toaster, last, {next: node})
    }
}

async function update_node_relation(toaster, direction, root, base, node, nodesMap) {
    switch (direction) {
        case consts.directions.right:
            await create_node_relation(toaster, root, node)
            await update(toaster, base, {child: node})
            break
        case consts.directions.down:
            await create_node_relation(toaster, root, node)
            await update(toaster, base, {next: node})
            break
        case consts.directions.up:
            await create_node_relation(toaster, root, node, base)
            await update_node_relation_insert(toaster, root, base, node, nodesMap)
            break
        case consts.directions.left:
            await create_node_relation(toaster, root, node, 0, base)
            await update_node_relation_insert(toaster, root, base, node, nodesMap)
            break
    }
}

export {update, update_node_relation}
