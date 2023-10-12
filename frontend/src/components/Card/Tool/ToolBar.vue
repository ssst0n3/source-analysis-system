<template>
  <span>
    <NaviButton :static-view="staticView" :node="node"
                v-on:save="$emit('save', $event)" v-on:navi="$emit('navi', arguments)"/>
    <AddButton class="ml-1"
               v-on:save="$emit('save', $event)"/>
    <EditButton class="ml-3"
                v-on:save="$emit('save', $event)"/>
    <CollapseButton class="ml-3"
                    :node="node"/>
    <RemoveButton :node="node"/>
  </span>
</template>

<script>
import AddButton from "@/components/Card/Tool/AddButton";
import NaviButton from "@/components/Card/Tool/NaviButton";
import CollapseButton from "@/components/Card/Tool/CollapseButton";
import RemoveButton from "@/components/Card/Tool/RemoveButton";
import EditButton from "@/components/Card/Tool/EditButton";

export default {
  name: "ToolBar",
  components: {EditButton, RemoveButton, CollapseButton, NaviButton, AddButton},
  props: {
    node: Object,
    staticView: Boolean,
  },
}
</script>

<style scoped>

</style>
