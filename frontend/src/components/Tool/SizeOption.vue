<template>
  <b-form-select v-model="size" :options="options"/>
</template>

<script>
export default {
  name: "SizeOption",
  data() {
    return {
      size: 'small',
      options: [
        { value: 'x-small', text: 'x-small' },
        { value: 'small', text: 'small' },
        { value: 'medium', text: 'medium'},
        { value: 'large', text: 'large' },
      ]
    }
  },
  watch: {
    size() {
      this.$emit('size_update', this.size)
    }
  }
}
</script>

<style scoped>
</style>