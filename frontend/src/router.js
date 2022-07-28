import Vue from 'vue'
import Router from 'vue-router'
import IndexView from "@/components/View/IndexView";
import StaticView from "@/components/View/StaticView";
import DynamicView from "@/components/View/DynamicView";


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
            component: DynamicView,
        },
        {
            path: '/static/:data_source',
            name: 'StaticView',
            component: StaticView,
        }
    ]
})
