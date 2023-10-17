<template>
  <b-badge pill variant="info">
    <b-link class="text-white">
      <b-icon-plus v-if="node.collapsed"></b-icon-plus>
      <b-icon-dash v-else @click.stop="collapse"></b-icon-dash>
    </b-link>
  </b-badge>
</template>

<script>
import {updateNode} from "@/util/node";

export default {
  name: "CollapseButton",
  props: {
    node: Object,
  },
  methods: {
    collapse() {
      if (confirm(`Are you sure to collapse node ${this.node.ID} ?`)) {
        updateNode(this, this.node.ID, {collapsed: true})
      }
    }
  }
}
</script>

<style scoped>

</style>
