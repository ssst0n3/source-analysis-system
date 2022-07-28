const api_ = '/api'
const v1 = `${api_}/v1`
const node = 'node'
const node_relation = 'node_relation'

const api = {
    v1: {
        node: `${v1}/${node}`,
        node_relation: `${v1}/${node_relation}`
    }
}

export default {
    // BaseUrl: '/',
    BaseUrl: process.env.NODE_ENV === 'development' ? 'http://127.0.0.1:16081' : '/',
    api: {
        v1: {
            node: {
                node: api.v1.node,
                list: (id) => `${api.v1.node}/list/${id}`,
                matrix: (id) => `${api.v1.node}/matrix/${id}`,
                update_markdown: (id) => `${api.v1.node}/${id}`,
            },
            node_relation: {
                list_by_root: (id) => `${api.v1.node_relation}/${id}`,
                node_relation: api.v1.node_relation,
                update_node_relation_by_node: (id) => `${api.v1.node_relation}/node/${id}`,
            }
        }
    }
}
