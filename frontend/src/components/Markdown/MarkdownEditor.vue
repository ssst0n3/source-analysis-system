<template>
  <div id="previewer">
    <b-form-textarea v-model="edit" debounce="300"/>
    <div id="preview" v-html="compiledMarkdown"/>
  </div>
</template>

<script>
import {marked} from 'marked'

export default {
  name: "MarkdownEditor",
  props: {
    markdown: String,
  },
  data() {
    return {
      edit: this.markdown
    }
  },
  computed: {
    compiledMarkdown() {
      return marked(this.edit, {sanitize: true});
    }
  }
}
</script>

<style scoped>
#previewer {
  margin: 0;
  width: 100%;
  height: 85%;
  display: inline-block;
  overflow-wrap: break-word;
  font-family: "Helvetica Neue", Arial, sans-serif;
  color: #333;
}

textarea,
#preview {
  display: inline-block;
  width: 49%;
  height: 100%;
  vertical-align: top;
  box-sizing: border-box;
  padding: 0 20px;
}

textarea {
  border: none;
  border-right: 1px solid #ccc;
  resize: none;
  outline: none;
  background-color: #f6f6f6;
  font-size: 14px;
  font-family: "Monaco", courier, monospace;
  height: 100% !important;
  padding: 20px;
}

code {
  color: #f66;
}

</style>
