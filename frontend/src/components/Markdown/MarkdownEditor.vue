<template>
  <div class="h-100">
    <b-card no-body id="previewer" class="h-100">
      <b-card-body class="h-90">
        <b-form-textarea class="h-100 d-inline-block overflow-auto"
                         v-if="true" v-model="edit" debounce="300"/>
        <MarkdownViewer class="d-inline-block overflow-auto text-break"
                        v-if="true" id="preview" :markdown="edit" :size="size"/>
      </b-card-body>
      <template #footer>
        <b-btn v-shortkey="{a: ['ctrl', 'enter'], b: ['meta', 'enter']}" @shortkey="save"
               @click="save" class="float-right">Save</b-btn>
      </template>
    </b-card>
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
  computed: {
    compiledMarkdown() {
      let m = marked.parse(this.edit);
      console.log(m)
      return m
    }
  },
  methods: {
    save() {
      this.$emit('save')
    }
  }
}
</script>

<style scoped>
#previewer {
  margin: 0;
  width: 100%;
  //height: 100% !important;
  font-family: "Helvetica Neue", Arial, sans-serif;
//color: #333;
}

textarea,
#preview {
  width: 49%;
  vertical-align: top;
  box-sizing: border-box;
  padding: 0 20px;
  height: 400px;
}

#preview {
  max-height: 100%;
}

textarea {
  border: none;
  border-right: 1px solid #ccc;
  resize: none;
  outline: none;
  background-color: #f6f6f6;
  font-size: 14px;
  font-family: "Monaco", courier, monospace;
  padding: 20px;
}

.h-90 {
  height: 90% !important;
}
</style>
