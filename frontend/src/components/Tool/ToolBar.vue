<template>
  <div>
    <DfsOption v-on:dfs-option-update="dfsOptionUpdate"/>
    <DownloadData v-if="!staticView" :node-matrix="nodeMatrix" :toc="toc"/>
    <TableOfContent :toc="toc"/>
    <SizeOption ref="size_option" v-on:size_update="size_update"/>
    <NodesCount toolbar :count="nodesCount"/>
  </div>
</template>

<script>

import DfsOption from "@/components/Tool/DfsOption.vue";
import DownloadData from "@/components/Tool/DownloadData.vue";
import TableOfContent from "@/components/TableOfContent/TableOfContent.vue";
import SizeOption from "@/components/Tool/SizeOption.vue";
import NodesCount from "@/components/Tool/NodesCount.vue";

export default {
  components: {NodesCount, SizeOption, TableOfContent, DownloadData, DfsOption},
  props: {
    toc: Object,
    staticView: Boolean,
    nodeMatrix: Array,
    nodesCount: Number,
  },
  methods: {
    dfsOptionUpdate(checked) {
      // TODO: store in cookie
      this.$emit('dfs-option-update', checked)
    },
    size_update(s) {
      // TODO: store in cookie
      this.$emit('size_update', s)
    }
  }
}
</script>

<style scoped>

</style>