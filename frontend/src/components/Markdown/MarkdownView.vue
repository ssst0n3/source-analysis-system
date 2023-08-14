<template>
  <div id="view" v-html="compiledMarkdown"/>
</template>

<script>
import {Marked} from "marked";
import {markedHighlight} from "marked-highlight";
import hljs from 'highlight.js';

const marked = new Marked(
    markedHighlight({
      langPrefix: 'hljs language-',
      highlight(code, lang) {
        const language = hljs.getLanguage(lang) ? lang : 'plaintext';
        return hljs.highlight(code, {language}).value;
      }
    })
)
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