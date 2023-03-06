<template>
  <div>
    <b-card :border-variant="active ? 'info' : ''" header-bg-variant="light">
      <template #header>
        <div class="text-right">
          <b-badge pill :variant="active ? hasNext?'secondary':'info' : 'light'" @click.stop="next"
                   v-if="!staticView || hasNext">
            <b-link class="text-white">
              <b-icon-arrow-down></b-icon-arrow-down>
            </b-link>
          </b-badge>
          <b-badge pill :variant="active ? hasChild?'secondary':'info' : 'light'" @click.stop="call">
            <b-link class="text-white">
              <b-icon-arrow-right></b-icon-arrow-right>
            </b-link>
          </b-badge>
          <b-badge pill :variant="active ? 'info' : 'light'" @click.stop="insert">
            <b-link class="text-white">
              <b-icon-arrow-up></b-icon-arrow-up>
            </b-link>
          </b-badge>
          <b-badge pill :variant="active ? 'info' : 'light'" @click.stop="down">
            <b-link class="text-white">
              <b-icon-arrow-left></b-icon-arrow-left>
            </b-link>
          </b-badge>
          <b-badge pill :variant="active ? 'info' : 'light'" @click.stop="down">
            <b-link class="text-white">
              <b-icon-x></b-icon-x>
            </b-link>
          </b-badge>
        </div>
      </template>
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
    nextId: Number,
    childId: Number,
    hasParent: Boolean,
    staticView: Boolean,
    active: Boolean,
  },
  data() {
    return {
      hasNext: this.nextId !== 0,
      hasChild: this.childId !== 0,
    };
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
    insert() {
      this.$emit('insert_node')
    }
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
