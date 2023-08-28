<template>
  <div id="view" v-html="compiledMarkdown"/>
</template>

<script>
import {marked} from "@/util/marked";

export default {
  name: "MarkdownViewer",
  props: {
    markdown: String,
  },
  data() {
    return {
      edit: this.markdown,
    }
  },
  computed: {
    compiledMarkdown() {
      return marked.parse(this.edit)
    }
  }
}
</script>

<style scoped>

</style>