<template>
  <div>
    <b-btn @click="$bvModal.show('modal-create-root-node')">
      <b-icon icon="plus"/>
    </b-btn>
    <div v-for="item in nodes" :key="item.id">
      <router-link :to="`/node/${item.node.ID}`">{{ item.node.title }}</router-link>
      <NodesCount class="ml-5" :count="item.count"/>
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
      <MarkdownEditor ref="create_root_node_markdown_editor" markdown="" v-on:save="createRootNode"/>
    </b-modal>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/util/const";
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";
import NodesCount from "@/components/Tool/NodesCount.vue";

export default {
  name: "RootNodes",
  components: {NodesCount, MarkdownEditor},
  data() {
    return {
      nodes: [],
      title: "",
      counts: [],
    }
  },
  created() {
    this.listNodesByRoot()
  },
  methods: {
    async listNodesByRoot() {
      let nodes = await lightweightRestful.api.get(consts.api.v1.node.list(0), null, {
        caller: this,
      })
      for (let i = 0; i < nodes.length; i++) {
        let node = nodes[i]
        let count = await lightweightRestful.api.get(consts.api.v1.node_relation.count(node.ID), null, {
          caller: this,
        })
        this.nodes.push({node: node, count: count})
      }
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
