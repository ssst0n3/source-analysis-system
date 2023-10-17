<template>
  <div>
    <b-modal :id="modalId" lazy hide-footer size="xl" v-if="!staticView">
      <MarkdownEditor :ref="editorRef"
                      :markdown="model.update ? node.markdown: ''" :size="size" :mode-update="model.update"
                      v-on:save="save"
      />
    </b-modal>
  </div>
</template>

<script>
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";
import {createNode, updateNode} from "@/util/node";
import {update_node_relation} from "@/util/nodeRelation";

export default {
  name: "NodeEditor",
  components: {MarkdownEditor},
  props: {
    staticView: Boolean,
    node: Object,
    size: String,
    model: Object,
    root: Number,
  },
  data() {
    return {
      modalId: `node-common-modal`,
      editorRef: `node-common-editor`,
    }
  },
  methods: {
    async save() {
      let content = this.$refs[this.editorRef].edit
      this.$bvModal.hide(this.modalId)
      if (this.model.update) {
        await updateNode(this, this.node.ID, {markdown:content})
        this.$emit('refresh')
      } else {
        let newNodeId = await createNode(this, content)
        await update_node_relation(this, this.model.direction, this.root, this.node, newNodeId)
        this.$emit('refreshWorld')
      }
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
