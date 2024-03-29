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
    BaseUrl: process.env.NODE_ENV === 'development' ? 'http://127.0.0.1:16080' : '/',
    api: {
        v1: {
            node: {
                node: api.v1.node,
                list: (id) => `${api.v1.node}/list/${id}`,
                matrix: (id) => `${api.v1.node}/matrix/${id}`,
                // update_markdown: (id) => `${api.v1.node}/${id}`,
                item: (id) => `${api.v1.node}/${id}`,
            },
            node_relation: {
                node_relation: api.v1.node_relation,
                list: (id) => `${api.v1.node_relation}/list/${id}`,
                count: (id) => `${api.v1.node_relation}/count/${id}`,
                item: (id) => `${api.v1.node_relation}/${id}`,
                update_node_relation_by_node: (id) => `${api.v1.node_relation}/node/${id}`,
                // unlink: (id) => `${api.v1.node_relation}/unlink/${id}`,
                hide_node: (id) => `${api.v1.node_relation}/hide/${id}`,
            }
        }
    },
    directions: {
        up: 1,
        left: 2,
        down: 3,
        right: 4,
    }
}
