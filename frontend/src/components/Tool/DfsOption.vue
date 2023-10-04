<template>
  <b-form-checkbox v-model="checked" name="check-button" size="lg" switch>
    {{status}}
  </b-form-checkbox>
</template>

<script>
export default {
  name: "DfsOption",
  data() {
    return {
      checked: false,
    }
  },
  computed: {
    status() {
      return this.checked ? "DFS" : "BFS"
    }
  },
  watch: {
    checked() {
      this.$emit('dfs-option-update', this.checked)
    }
  }
}
</script>

<style scoped>
</style>