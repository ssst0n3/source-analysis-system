<template>
  <span :class="toolbar ? 'nodes-count' : ''">{{ count }} nodes</span>
</template>

<script>

export default {
  name: "NodesCount",
  props: {
    count: Number,
    toolbar: Boolean,
  },
  data() {
    return {}
  }
}
</script>

<style scoped>
.nodes-count {
  width: 8%;
  position: fixed;
  bottom: 20px;
  right: 360px;
  border: none;
}
</style>