import Vue from 'vue'
import App from './App.vue'

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

import highlightPlugin from "@highlightjs/vue-plugin";
import 'highlight.js/styles/github.css'
import 'highlight.js/lib/common';
Vue.use(highlightPlugin);

import consts from "@/util/const";
import lightweightRestful from "vue-lightweight_restful";
lightweightRestful.api.initClient(consts.BaseUrl)

import ShortKey from 'vue-shortkey'
Vue.use(ShortKey)

Vue.use(require('vue-cookies'))

import hljs from "highlight.js";
import VMdEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';
VMdEditor.use(githubTheme, {
  Hljs: hljs,
});
Vue.use(VMdEditor)

// copy to clipboard
import VueClipboard from 'vue-clipboard2'
Vue.use(VueClipboard)

import router from '@/router'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
