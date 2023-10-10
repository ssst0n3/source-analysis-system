<template>
  <span>
    <NaviButton :node="node"
                v-on:navi="$emit('navi', arguments)" v-on:refresh="$emit('refreshWorld')"/>
    <AddButton class="ml-1"
               :root="root" :static-view="staticView" :node="node"
               v-on:refresh="$emit('refreshWorld')"/>
    <CollapseButton class="ml-3"
                    :node="node"/>
    <RemoveButton :node="node"
                  v-on:refresh="$emit('refresh')"/>
  </span>
</template>

<script>
import AddButton from "@/components/Card/Tool/AddButton";
import NaviButton from "@/components/Card/Tool/NaviButton";
import CollapseButton from "@/components/Card/Tool/CollapseButton";
import RemoveButton from "@/components/Card/Tool/RemoveButton";

export default {
  name: "ToolBar",
  components: {RemoveButton, CollapseButton, NaviButton, AddButton},
  props: {
    root: Number,
    node: Object,
    staticView: Boolean,
  },
  methods: {}
}
</script>

<style scoped>

</style>
