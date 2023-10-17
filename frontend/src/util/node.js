import lightweightRestful from "vue-lightweight_restful";
import consts from "@/util/const";

async function updateNode(caller, id, node){
    await lightweightRestful.api.put(consts.api.v1.node.item(id), null, node, {
        caller: caller,
        success_msg: 'update successfully'
    })
}

async function createNode(caller, markdown) {
    let response = await lightweightRestful.api.post(consts.api.v1.node.node, null, {
        markdown: markdown
    }, {
        caller: caller,
    })
    return response.id
}

export {
    updateNode,
    createNode,
}
