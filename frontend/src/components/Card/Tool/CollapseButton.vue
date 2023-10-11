<template>
  <b-badge pill variant="info" @click.stop="collapse">
    <b-link class="text-white">
      <b-icon-dash v-if="node.collapsed"></b-icon-dash>
      <b-icon-dash v-else></b-icon-dash>
    </b-link>
  </b-badge>
</template>

<script>
export default {
  name: "CollapseButton",
  props: {
    node: Object,
  },
  methods: {
    collapse() {
      if (confirm(`Are you sure to collapse node ${this.node.ID} ?`)) {
        alert("TODO: add an collapsed property for each node #9")
      }
    }
  }
}
</script>

<style scoped>

</style>
