<template>
  <div>
    <label class="collapse btn btn-secondary toc-btn" for="_1">
        <b-icon icon="justify"/>
    </label>
    <input id="_1" type="radio" name="c1">
    <b-sidebar id="sidebar-toc" title="Toc" right shadow visible>
      <div class="px-3 py-2">
        <ul>
          <li :class="'toc-'+h.nodeName.toLowerCase()" v-for="(h, index) in toc" :key="`header-${index}`">
            <a :href="'#'+h.id">{{ h.innerText }}</a>
          </li>
        </ul>
      </div>
      <template #header>
        <label class="collapse" for="_2">
          <b-icon icon="x" variant="secondary"></b-icon>
        </label>
        <input id="_2" type="radio" name="c1">
      </template>

    </b-sidebar>
  </div>
</template>

<script>
export default {
  name: "SideBarStatic",
  props: {
    toc: Array,
  },
}
</script>

<style scoped>
.collapse{
  cursor: pointer;
  display: block;
}
.collapse + input{
  display: none; /* hide the checkboxes */
}
.collapse + input + div{
  display:none;
}
.collapse + input:checked + div{
  display:block;
}

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
