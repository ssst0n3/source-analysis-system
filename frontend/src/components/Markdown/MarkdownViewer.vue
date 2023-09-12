<template>
  <div class="markdown-viewer overflow-auto" v-html="compiledMarkdown" :style="{'font-size': size}"/>
</template>

<script>
import {marked} from "@/util/marked";

export default {
  name: "MarkdownViewer",
  props: {
    markdown: String,
    size: String,
  },
  data() {
    return {
      // edit: this.markdown,
    }
  },
  computed: {
    compiledMarkdown() {
      return marked.parse(this.markdown)
    }
  }
}
</script>

<style scoped>
.markdown-viewer{
  max-height: 500px;
}
</style>