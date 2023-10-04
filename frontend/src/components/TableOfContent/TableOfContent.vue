<template>
  <div>
    <SideBar :toc="toc"/>
<!--    <SideBarStatic :toc="toc"/>-->
  </div>
</template>

<script>
import SideBar from "@/components/TableOfContent/SideBar";

export default {
  name: "TableOfContent",
  components: {SideBar},
  props: {
    toc: Object,
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
