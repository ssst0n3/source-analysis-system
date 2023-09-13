<template>
  <div>
    <span class="nodes-count">{{ count }} nodes</span>
  </div>
</template>

<script>

export default {
  name: "NodesCount",
  props: {
    count: Number,
  },
  data() {
    return {
    }
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