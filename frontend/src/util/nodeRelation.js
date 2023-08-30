import consts from "@/util/const"
import lightweightRestful from "vue-lightweight_restful";

async function create_node_relation(caller, root, node, next, child) {
    next = next === undefined ? 0 : next
    child = child === undefined ? 0 : child
    await lightweightRestful.api.post(consts.api.v1.node_relation.node_relation, null, {
        root: root,
        node: node,
        next: next,
        child: child,
    }, {
        caller: caller,
    })
}

async function update(caller, base, data) {
    await lightweightRestful.api.put(
        consts.api.v1.node_relation.update_node_relation_by_node(base),
        null,
        data,
        {
            caller: caller,
        }
    )
}

async function update_node_relation_insert(caller, root, base, node, nodesMap) {
    await create_node_relation(caller, root, node, base)
    let baseNode = nodesMap[base]
    let parent = baseNode.parent
    let last = baseNode.last
    if (parent !== undefined) {
        await update(caller, parent, {child: node})
    }
    if (last !== undefined) {
        await update(caller, last, {next: node})
    }
}

async function update_node_relation(caller, direction, root, base, node, nodesMap) {
    switch (direction) {
        case consts.directions.right:
            await create_node_relation(caller, root, node)
            await update(caller, base, {child: node})
            break
        case consts.directions.down:
            await create_node_relation(caller, root, node)
            await update(caller, base, {next: node})
            break
        case consts.directions.up:
            await update_node_relation_insert(caller, root, base, node, nodesMap)
            break
        case consts.directions.left:
            alert("left")
            break
    }
}

export {update_node_relation}