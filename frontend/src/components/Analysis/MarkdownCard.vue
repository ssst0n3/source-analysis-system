<template>
  <div>
    <b-card :border-variant="active ? 'info' : ''" header-bg-variant="light"
            @dblclick="$bvModal.show('modal-node-'+nodeId)">
      <template #header>
        <div class="text-right">
          <b-badge comment="just a place holder" pill variant="light">
            <b-link class="text-white">
              <b-icon-arrow-up/>
            </b-link>
          </b-badge>
          <span v-if="active" >
            <b-badge pill :variant="hasNext?'secondary':'info'" @click.stop="navi(directions.down)"
                     v-if="!staticView || hasNext">
              <b-link class="text-white">
                <b-icon-arrow-down/>
              </b-link>
            </b-badge>
            <b-badge pill :variant="hasChild?'secondary':'info'" @click.stop="navi(directions.right)">
              <b-link class="text-white">
                <b-icon-arrow-right/>
              </b-link>
            </b-badge>
            <b-badge v-if="hasLast" pill variant="secondary" @click.stop="navi(directions.up)">
              <b-link class="text-white">
                <b-icon-arrow-up/>
              </b-link>
            </b-badge>
            <b-badge pill variant="secondary" @click.stop="navi(directions.left)">
              <b-link class="text-white">
                <b-icon-arrow-left/>
              </b-link>
            </b-badge>
            <b-badge class="ml-1" pill :variant="active ? 'info' : 'light'" @click.stop="add(directions.up)">
              <b-link class="text-white">
                <b-icon-arrow-bar-up/>
              </b-link>
            </b-badge>
            <b-badge pill :variant="active ? 'info' : 'light'" @click.stop="add(directions.left)">
              <b-link class="text-white">
                <b-icon-arrow-bar-left/>
              </b-link>
            </b-badge>
            <b-badge pill :variant="active ? 'info' : 'light'" @click.stop="remove">
              <b-link class="text-white">
                <b-icon-x></b-icon-x>
              </b-link>
            </b-badge>
          </span>
        </div>
      </template>
      <MarkdownViewer :markdown="markdown" style="width: 100%"/>
    </b-card>
  </div>
</template>

<script>
import consts from "@/util/const";
import MarkdownViewer from "@/components/Markdown/MarkdownView.vue";

export default {
  name: "MarkdownCard",
  components: {MarkdownViewer},
  props: {
    nodeId: String,
    markdown: String,
    nextId: Number,
    childId: Number,
    lastId: Number,
    parentId: Number,
    hasParent: Boolean,
    staticView: Boolean,
    active: Boolean,
  },
  data() {
    let directions2Id = {}
    directions2Id[consts.directions.up] = this.lastId
    directions2Id[consts.directions.down] = this.nextId
    directions2Id[consts.directions.left] = this.parentId
    directions2Id[consts.directions.right] = this.childId
    return {
      hasNext: this.nextId !== 0,
      hasChild: this.childId !== 0,
      hasLast: this.lastId !== undefined,
      directions: consts.directions,
      directions2Id: directions2Id,
    }
  },
  methods: {
    exists(nodeId) {
      return nodeId !== undefined  && nodeId !== 0
    },
    add(direction) {
      this.$emit('add', this.nodeId, direction)
    },
    /*
    * navi or add node
    *
    */
    navi(direction) {
      let navId = this.directions2Id[direction]
      if (this.exists(navId) || direction === consts.directions.left) {
        this.$emit('navi', this.nodeId, navId, direction)
      } else {
        this.add(direction)
      }
    },
    remove() {
      this.$emit('remove', this.nodeId)
    },
  }
}
</script>

<style scoped>
/*noinspection CssUnusedSymbol*/
/deep/ .modal-dialog {
  width: 90%;
  max-width: 100%;
  height: 100%;
}

/*noinspection CssUnusedSymbol*/
/deep/ .modal-content {
  height: 100%;
}
</style>
