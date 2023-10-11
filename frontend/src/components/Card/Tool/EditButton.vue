<template>
<span>
    <b-badge pill variant="info" @click.stop="edit">
      <b-link class="text-white">
        <b-icon-pencil/>
      </b-link>
    </b-badge>
    <b-modal :id="modalId" lazy hide-footer size="xl" v-if="!staticView">
      <MarkdownEditor :ref="editorRef"
                      :markdown="node.markdown" :size="size"
                      v-on:save="save"
      />
    </b-modal>
</span>
</template>

<script>
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";
import {updateNode} from "@/util/node";

export default {
  name: "EditButton",
  components: {MarkdownEditor},
  props: {
    staticView: Boolean,
    node: Object,
    size: String,
  },
  data() {
    return {
      modalId: `node-edit-${this.node.ID}`,
      editorRef: `node-edit-markdown-editor-common-${this.node.ID}`,
    }
  },
  methods: {
    edit() {
      this.$bvModal.show(this.modalId)
    },
    async save() {
      let content = this.$refs[this.editorRef].edit
      await updateNode(this, this.node.ID, content)
      await this.$emit('refresh')
      this.$bvModal.hide(this.modalId)
    },
  }
}
</script>

<style scoped>
/*noinspection CssUnusedSymbol*/
/deep/ .modal-dialog {
  width: 90%;
  max-width: 100%;
  height: 90%;
}

/*noinspection CssUnusedSymbol*/
/deep/ .modal-content {
  height: 100%;
}
</style>
