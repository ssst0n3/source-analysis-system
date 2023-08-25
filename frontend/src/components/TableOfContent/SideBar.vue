<template>
  <div>
    <b-button class="toc-btn" v-b-toggle.sidebar-toc>
      <b-icon icon="justify"/>
    </b-button>
    <b-sidebar id="sidebar-toc" title="Toc" right visible shadow>
      <div class="px-3 py-2">
        <a @click.prevent="anchor(toc.title.id)" :href="'#'+toc.title.id" v-if="toc.title !== null">
          <strong>{{ toc.title.innerText }}</strong>
        </a>
        <hr>
        <ul>
          <li :class="'toc-'+h.nodeName.toLowerCase()" v-for="(h, index) in toc.headings" :key="`header-${index}`">
            <a @click.prevent="anchor(h.id)" :href="'#'+h.id">{{ h.innerText }}</a>
          </li>
        </ul>
      </div>
    </b-sidebar>
  </div>
</template>

<script>
import {anchor} from "@/util/util";

export default {
  name: "SideBar",
  props: {
    toc: Object,
  },
  methods: {
    anchor: anchor
  }
}
</script>

<style scoped>
.toc-btn {
  position: fixed;
  bottom: 20px;
  right: 20px;
  background-color: rgba(0, 0, 0, 0.2);
  border: none;
}

.toc-h3 {
  margin-left: 10px;
}

.toc-h4 {
  margin-left: 20px;
}

.toc-h5 {
  margin-left: 30px;
}

.toc-h6 {
  margin-left: 40px;
}
</style>
