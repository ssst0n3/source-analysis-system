import Vue from 'vue'
import Router from 'vue-router'
import NodeMatrix from "@/components/Node/NodeMatrix";
import IndexView from "@/components/IndexView";


Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'IndexView',
            component: IndexView
        },
        {
            path: '/node/:id',
            name: 'NodeMatrix',
            component: NodeMatrix,
        }
    ]
})
