<template>
  <div>
    <b-card>
      <div style="width: 100%" v-html="markdownToHtml" @dblclick="$bvModal.show('modal-node-'+nodeId)"/>
    </b-card>
    <b-modal :id="'modal-node-'+nodeId" hide-footer size="xl">
      <MarkdownEditor ref="markdown_editor" :markdown="markdown"/>
      <div id="panel" class="float-right">
        <b-btn @click="save">save</b-btn>
      </div>
    </b-modal>
  </div>
</template>

<script>
import {marked} from 'marked';
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/util/const";

export default {
  name: "MarkdownCard",
  components: {MarkdownEditor},
  props: {
    nodeId: String,
    markdown: String,
  },
  data() {
    return {};
  },
  computed: {
    markdownToHtml() {
      return marked(this.markdown);
    }
  },
  methods: {
    async save() {
      await lightweightRestful.api.put(consts.api.v1.node.item(this.nodeId), null, {
        markdown: this.$refs.markdown_editor.edit
      }, {
        caller: this,
        success_msg: 'update successfully'
      })
      this.$bvModal.hide('modal-node-'+this.nodeId)
      this.$emit('update_node')
    }
  }
}
</script>

<style scoped>
/deep/ .modal-dialog {
  width: 90%;
  max-width: 100%;
  height: 100%;
}

/deep/ .modal-content {
  height: 100%;
}
</style>
