<template>
  <b-form-checkbox class="dfs-option" v-model="checked" name="check-button" size="lg" switch>
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
.dfs-option {
  position: fixed;
  bottom: 20px;
  right: 180px;
  /*background-color: rgba(0, 0, 0, 0.2);*/
  border: none;
}
</style>