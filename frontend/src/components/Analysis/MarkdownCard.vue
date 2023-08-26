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
            <b-badge pill :variant="hasNext?'secondary':'info'" @click.stop="next"
                     v-if="!staticView || hasNext">
              <b-link class="text-white">
                <b-icon-arrow-down/>
              </b-link>
            </b-badge>
            <b-badge pill :variant="hasChild?'secondary':'info'" @click.stop="call">
              <b-link class="text-white">
                <b-icon-arrow-right/>
              </b-link>
            </b-badge>
            <b-badge v-if="hasLast" pill variant="secondary" @click.stop="navi(directions.up)">
              <b-link class="text-white">
                <b-icon-arrow-up/>
              </b-link>
            </b-badge>
            <b-badge v-if="hasParent" pill variant="secondary" @click.stop="navi(directions.left)">
              <b-link class="text-white">
                <b-icon-arrow-left/>
              </b-link>
            </b-badge>
            <b-badge class="ml-1" pill :variant="active ? 'info' : 'light'" @click.stop="insertUp">
              <b-link class="text-white">
                <b-icon-arrow-bar-up/>
              </b-link>
            </b-badge>
            <b-badge pill :variant="active ? 'info' : 'light'" @click.stop="insertLeft">
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
    <b-modal :id="'modal-node-'+nodeId" hide-footer size="xl">
      <MarkdownEditor ref="markdown_editor" :markdown="markdown"/>
      <div id="panel" class="float-right">
        <b-btn @click="save">save</b-btn>
      </div>
    </b-modal>
  </div>
</template>

<script>
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/util/const";
import MarkdownViewer from "@/components/Markdown/MarkdownView.vue";

export default {
  name: "MarkdownCard",
  components: {MarkdownViewer, MarkdownEditor},
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
    return {
      hasNext: this.nextId !== 0,
      hasChild: this.childId !== 0,
      hasLast: this.lastId !== undefined,
      directions: consts.directions,
    };
  },
  methods: {
    async save() {
      await lightweightRestful.api.put(consts.api.v1.node.item(this.nodeId), null, {
        markdown: this.$refs.markdown_editor.edit
      }, {
        caller: this,
        success_msg: 'update successfully'
      })
      this.$bvModal.hide('modal-node-' + this.nodeId)
      this.$emit('update_node')
    },
    down() {
      alert(1)
    },
    next() {
      this.$emit('next', this.nodeId, this.nextId)
    },
    call() {
      this.$emit('call', this.nodeId, this.childId)
    },
    navi(direction) {
      switch (direction) {
        case this.directions.up:
          this.$emit('navi', this.lastId)
          break
        case this.directions.left:
          this.$emit('navi', this.parentId)
          break
        case this.directions.down:
          this.$emit('navi', this.nextId)
          break
        case this.directions.right:
          this.$emit('navi', this.childId)
          break
      }
    },
    insert() {
      this.$emit('insert', this.nodeId, this.lastId)
    },
    caller() {
      this.$emit('caller', this.nodeId, this.parentId)
    },
    insertUp() {
      this.$emit('insertUp', this.nodeId, this.lastId)
    },
    insertLeft() {
      this.$emit('insertLeft', this.nodeId, this.parentId)
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
