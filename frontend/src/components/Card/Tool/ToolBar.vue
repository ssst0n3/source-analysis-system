<template>
  <span>
    <NaviButton :static-view="staticView" :node="node"
                v-on:save="save" v-on:navi="$emit('navi', arguments)"/>
    <AddButton class="ml-1"
               :directions="directions" :save="save"/>
    <EditButton class="ml-3"
                :save="save"/>
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
import consts from "@/util/const";

export default {
  name: "ToolBar",
  components: {EditButton, RemoveButton, CollapseButton, NaviButton, AddButton},
  props: {
    node: Object,
    staticView: Boolean,
  },
  data() {
    return {
      directions: consts.directions,
    }
  },
  methods: {
    save($event) {
      this.$emit('save', $event)
    }
  }
}
</script>

<style scoped>

</style>
