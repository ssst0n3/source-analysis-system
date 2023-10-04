<template>
  <span class="align-middle">{{ count }} nodes</span>
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
</style>