<template>
  <b-container class="toolbar">
    <b-row align-v="center">
      <b-col>
        <DfsOption v-on:dfs-option-update="dfsOptionUpdate"/>
      </b-col>
      <b-col>
        <SizeOption ref="size_option" v-on:size_update="size_update"/>
      </b-col>
      <b-col>
        <NodesCount :count="nodesCount"/>
      </b-col>
      <b-col>
        <DownloadData v-if="!staticView" :node-matrix="nodeMatrix" :toc="toc"/>
      </b-col>
      <b-col>
        <TableOfContent :toc="toc"/>
      </b-col>
    </b-row>
  </b-container>
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
.toolbar {
  width: 18%;
  position: fixed;
  bottom: 20px;
  right: 360px;
  border: none;
}
</style>