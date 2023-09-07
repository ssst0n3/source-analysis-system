<template>
  <b-form-select class="size-option" v-model="size" :options="options"/>
</template>

<script>
export default {
  name: "SizeOption",
  data() {
    return {
      size: 'small',
      options: [
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
.size-option {
  width: 8%;
  position: fixed;
  bottom: 20px;
  right: 280px;
  /*background-color: rgba(0, 0, 0, 0.2);*/
  border: none;
}
</style>