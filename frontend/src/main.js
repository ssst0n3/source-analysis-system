import Vue from 'vue'
import App from './App.vue'

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

// import hljs from "highlight.js";
import highlightPlugin from "@highlightjs/vue-plugin";
import 'highlight.js/styles/github.css'
import 'highlight.js/lib/common';
Vue.use(highlightPlugin);

import consts from "@/util/const";
import lightweightRestful from "vue-lightweight_restful";
lightweightRestful.api.initClient(consts.BaseUrl)

import router from '@/router'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
