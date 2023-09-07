<template>
  <div id="previewer">
    <b-form-textarea v-model="edit" debounce="300"/>
<!--    <div id="preview" v-html="compiledMarkdown"/>-->
    <MarkdownViewer id="preview" :markdown="edit" :size="size"/>
  </div>
</template>

<script>
import {Marked} from "marked";
import {markedHighlight} from "marked-highlight";
import hljs from 'highlight.js';
import MarkdownViewer from "@/components/Markdown/MarkdownViewer.vue";

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
  name: "MarkdownEditor",
  components: {MarkdownViewer},
  props: {
    markdown: String,
    size: String,
  },
  data() {
    return {
      edit: this.markdown,
    }
  },
  mounted() {

  },
  computed: {
    compiledMarkdown() {
      let m = marked.parse(this.edit);
      console.log(m)
      return m
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
  //color: #333;
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
  //color: #f66;
}

</style>
