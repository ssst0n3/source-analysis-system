<template>
  <div>
    <b-btn @click="$bvModal.show('modal-create-root-node')">
      <b-icon icon="plus"/>
    </b-btn>
    <div v-for="node in nodes" :key="node.id">
      <router-link :to="`/node/${node.ID}`">{{ node.markdown }}</router-link>
    </div>
    <b-modal id="modal-create-root-node" hide-footer size="xl">
      <b-form-group
          id="fieldset-horizontal"
          label-cols-sm="4"
          label-cols-lg="3"
          content-cols-sm
          content-cols-lg="7"
          description="display in summary mode"
          label="Title"
          label-for="input-horizontal"
      >
        <b-form-input v-model="title"/>
      </b-form-group>
      <p>markdown</p>
      <MarkdownEditor ref="create_root_node_markdown_editor" markdown=""/>
      <div id="panel" class="float-right">
        <b-btn @click="createRootNode">save</b-btn>
      </div>
    </b-modal>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/util/const";
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";

export default {
  name: "RootNodes",
  components: {MarkdownEditor},
  data() {
    return {
      nodes: [],
      title: "",
    }
  },
  created() {
    this.listNodesByRoot()
  },
  methods: {
    async listNodesByRoot() {
      this.nodes = await lightweightRestful.api.get(consts.api.v1.node.list(0), null, {
        caller: this,
      })
    },
    async createRootNode() {
      let response = await lightweightRestful.api.post(consts.api.v1.node.node, null, {
        title: this.title,
        markdown: this.$refs.create_root_node_markdown_editor.edit
      }, {
        caller: this,
      })
      let id = response.id
      await lightweightRestful.api.post(consts.api.v1.node_relation.node_relation, null, {
        root: this.root,
        node: id,
      }, {
        caller: this,
      })
      await lightweightRestful.api.post(consts.api.v1.node_relation.node_relation, null, {
        root: this.root,
        node: id,
      }, {
        caller: this,
      })
      this.$bvModal.hide('modal-create-root-node')
    }
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
