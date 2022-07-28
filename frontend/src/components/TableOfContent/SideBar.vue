<template>
  <div>
    <b-button class="toc-btn" v-b-toggle.sidebar-toc>
      <b-icon icon="justify"/>
    </b-button>
    <b-sidebar id="sidebar-toc" title="Toc" right shadow>
      <div class="px-3 py-2">
        <ul>
          <li :class="'toc-'+h.nodeName.toLowerCase()" v-for="(h, index) in toc" :key="`header-${index}`">
            <a :href="'#'+h.id">{{ h.innerText }}</a>
          </li>
        </ul>
      </div>
    </b-sidebar>
  </div>
</template>

<script>
export default {
  name: "SideBar",
  props: {
    toc: Array,
  },
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
