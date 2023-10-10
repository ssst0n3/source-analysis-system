<template>
  <span>
    <b-badge pill variant="info" @click.stop="add(directions.up)">
      <b-link class="text-white">
        <b-icon-arrow-bar-up/>
      </b-link>
    </b-badge>
    <b-badge pill variant="info" @click.stop="add(directions.left)">
      <b-link class="text-white">
        <b-icon-arrow-bar-left/>
      </b-link>
    </b-badge>
    <b-modal :id="modalId" lazy hide-footer size="xl" v-if="!staticView">
      <MarkdownEditor :ref="editorRef" :size="size"
                      v-on:save="save"
      />
    </b-modal>
  </span>
</template>

<script>
import consts from "@/util/const";
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";
import {createNode} from "@/util/node";
import {update_node_relation} from "@/util/nodeRelation";

export default {
  name: "AddButton",
  components: {MarkdownEditor},
  props: {
    root: Number,
    staticView: Boolean,
    node: Object,
    size: String,
  },
  data() {
    return {
      modalId: `node-add-${this.node.ID}`,
      editorRef: `add-markdown-editor-common-${this.node.ID}`,
      directions: consts.directions,
      direction: 0,
    }
  },
  methods: {
    reset() {
      this.direction = 0
      this.$bvModal.hide(this.modalId)
    },
    add(direction) {
      this.reset()
      this.direction = direction
      this.$bvModal.show(this.modalId)
    },
    async save() {
      let content = this.$refs[this.editorRef].edit
      let newNodeId = await createNode(this, content)
      await update_node_relation(this, this.direction, this.root, this.node, newNodeId)
      this.reset()
      await this.$emit('refresh')
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
