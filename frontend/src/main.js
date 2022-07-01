import Vue from 'vue'
import App from './App.vue'

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

import consts from "@/util/const";
import lightweightRestful from "vue-lightweight_restful";
lightweightRestful.api.initClient(consts.BaseUrl)


Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
