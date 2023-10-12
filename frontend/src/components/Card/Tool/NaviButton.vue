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
  </span>
</template>

<script>
import consts from "@/util/const";

export default {
  name: "NaviButton",
  props: {
    staticView: Boolean,
    node: Object,
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
    exists(nodeId) {
      return nodeId !== undefined && nodeId !== 0
    },
    navi(direction) {
      let navId = this.directions2Id[direction]
      if (this.exists(navId) || direction === consts.directions.left) {
        this.$emit('navi', this.node.ID, navId, direction)
      } else {
        this.$emit('save', {update: false, direction: direction})
      }
    },
  }
}
</script>

<style scoped>

</style>
