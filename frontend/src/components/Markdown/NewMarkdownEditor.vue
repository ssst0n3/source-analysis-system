<template>
  <v-md-editor right-toolbar="preview toc sync-scroll fullscreen esc" :toolbar="toolbar"
               v-model="edit" height="600px"></v-md-editor>
</template>

<script>
export default {
  name: "NewMarkdownEditor",
  props: {
    markdown: String,
    size: String,
  },
  data() {
    return {
      edit: this.markdown,
      toolbar: {
        esc: {
          icon: 'v-md-icon-x'
        }
      }
    }
  },
}
</script>

<style scoped>

</style>