<template>
  <span>
    <b-badge pill
             :variant="neighbour.hasNext?'secondary':'info'"
             @click.stop="navi(directions.down)"
             v-if="!staticView || neighbour.hasNext">
      <b-link class="text-white">
        <b-icon-arrow-down/>
      </b-link>
    </b-badge>
    <b-badge pill
             :variant="neighbour.hasChild?'secondary':'info'"
             @click.stop="navi(directions.right)">
      <b-link class="text-white">
        <b-icon-arrow-right/>
      </b-link>
    </b-badge>
    <b-badge v-if="neighbour.hasLast" pill variant="secondary" @click.stop="navi(directions.up)">
      <b-link class="text-white">
        <b-icon-arrow-up/>
      </b-link>
    </b-badge>
    <b-badge pill variant="secondary" @click.stop="navi(directions.left)">
      <b-link class="text-white">
        <b-icon-arrow-left/>
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
import consts from "@/util/const";
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";
import {createNode} from "@/util/node";
import {update_node_relation} from "@/util/nodeRelation";

export default {
  name: "NaviButton",
  components: {MarkdownEditor},
  props: {
    staticView: Boolean,
    node: Object,
    size: String,
  },
  data() {
    let directions2Id = {}
    directions2Id[consts.directions.up] = this.node.last
    directions2Id[consts.directions.down] = this.node.next
    directions2Id[consts.directions.left] = this.node.parent
    directions2Id[consts.directions.right] = this.node.child
    let neighbour = {
      hasNext: this.node.next !== 0,
      hasChild: this.node.child !== 0,
      hasLast: this.node.last !== undefined,
      hasParent: this.node.parent !== undefined,
    }
    return {
      modalId: `node-navi-${this.node.ID}`,
      editorRef: `navi-markdown-editor-common-${this.node.ID}`,
      neighbour: neighbour,
      directions: consts.directions,
      directions2Id: directions2Id,
    }
  },
  methods: {
    reset() {
      this.direction = 0
      this.markdown = ""
      this.$bvModal.hide(this.modalId)
    },
    exists(nodeId) {
      return nodeId !== undefined && nodeId !== 0
    },
    navi(direction) {
      let navId = this.directions2Id[direction]
      if (this.exists(navId) || direction === consts.directions.left) {
        this.$emit('navi', this.node.ID, navId, direction)
      } else {
        this.add(direction)
      }
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

</style>
