<template>
  <div>
    <SideBar :toc="toc"/>
<!--    <SideBarStatic :toc="toc"/>-->
  </div>
</template>

<script>
import SideBar from "@/components/TableOfContent/SideBar";
import SideBarStatic from "@/components/TableOfContent/SideBarStatic";

export default {
  name: "TableOfContent",
  // eslint-disable-next-line vue/no-unused-components
  components: {SideBarStatic, SideBar},
  props: {
    toc: Array,
  },
  data() {
    return {
      headings: [
        {
          nodeName: "h1",
          innerText: "heading"
        }
      ]
    }
  },
  mounted() {
    // this.$nextTick(() => {
    //   this.find_headings()
    // })
  },
  methods: {}
}
</script>

<style scoped>

</style>
