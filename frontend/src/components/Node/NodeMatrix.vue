<template>
  <div>
    <b-toast id="my-toast" variant="warning" solid no-auto-hide :visible="loading">
      <template #toast-title>
        <div class="d-flex flex-grow-1 align-items-baseline">
          <b-img blank blank-color="#ff5555" class="mr-2" width="12" height="12"></b-img>
          <strong class="mr-auto">Loading</strong>
          <small class="text-muted mr-2">42 seconds ago</small>
        </div>
      </template>
      <b-spinner/>
    </b-toast>
    <div v-for="(col, index) in nodeMatrix" :key="`col-${index}`" style="white-space: nowrap;" class="mt-5">
      <div @mouseover="mouseover"
           :ref="'node-'+node.ID" :id="'node-'+node.ID" v-for="(node, index2) in col"
           :key="`col-${index}-row-${index2}`"
           style="display: inline-block; width: 500px; vertical-align:middle"
           class="ml-5" :class="{hidden: node.ID === 0}">
        <b-btn @click="next(node.ID)">Next</b-btn>
        <b-btn @click="call(node.ID)" class="ml-2">Call</b-btn>
        <MarkdownCard style="white-space: normal" :id="node.ID.toString()" :markdown="node.markdown.toString()"
                      v-on:update_node="refreshWorld"
                      v-if="node.markdown.length>0"/>
        <AnalysisItem v-else/>
      </div>
    </div>
    <b-modal id="node-common" hide-footer size="xl">
      <MarkdownEditor ref="markdown_editor_common" markdown=""/>
      <div id="panel" class="float-right">
        <b-btn @click="save">save</b-btn>
      </div>
    </b-modal>
  </div>
</template>

<script>
import Matrix from "@/util/matrix";

const modeNext = 1
const modeCall = 2
import consts from "@/util/const";
import {jsPlumb} from "jsplumb";
import lightweightRestful from "vue-lightweight_restful";
import MarkdownCard from "@/components/Analysis/MarkdownCard";
import AnalysisItem from "@/components/Analysis/AnalysisItem";
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";

export default {
  name: "NodeMatrix",
  components: {AnalysisItem, MarkdownCard, MarkdownEditor},
  data() {
    return {
      root: parseInt(this.$route.params.id),
      nodesMap: {},
      nodeMatrix: [],
      nodeRelations: [],
      nodeRelationsMap: {},
      mode: 0,
      baseNode: 0,
      plumbInstance: null,
      nodeLoading: false,
      nodeRelationsLoading: false,
    }
  },
  computed: {
    loading() {
      return this.nodeLoading || this.nodeRelationsLoading
    }
  },
  async created() {
    await this.refreshData()
  },

  mounted() {
    this.plumbInstance = jsPlumb.getInstance()
    // await this.listNodeRelationsByRoot(this.root)
  },
  watch: {
    nodeMatrix: async function () {
      this.$nextTick(() => {
        this.drawLine()
      })
    }
  },
  methods: {
    async refreshData() {
      await this.listNodes(this.root)
      await this.listNodeRelationsByRoot(this.root)
      let matrix = new Matrix.Matrix(this.root, this.nodeRelationsMap)
      matrix.childRecursive(0)
      this.nodeMatrix = matrix.dumpNode(this.nodesMap)
    },
    async refreshWorld() {
      this.plumbInstance.deleteEveryConnection()
      this.plumbInstance.deleteEveryEndpoint()
      // this.renderComponent = false
      await this.refreshData()
      // this.$nextTick(() => {
      //   this.renderComponent = true;
      // });
      // this.$nextTick(() => {
      //   this.drawLine()
      // })
    },
    async save() {
      let response = await lightweightRestful.api.post(consts.api.v1.node.node, null, {
        markdown: this.$refs.markdown_editor_common.edit
      }, {
        caller: this,
      })
      let id = response.id
      let data = {}
      if (this.mode === modeCall) {
        data.child = id
      } else if (this.mode === modeNext) {
        data.next = id
      }
      await lightweightRestful.api.post(consts.api.v1.node_relation.node_relation, null, {
        root: this.root,
        node: id,
      }, {
        caller: this,
      })
      await lightweightRestful.api.put(consts.api.v1.node_relation.update_node_relation_by_node(this.baseNode), null, data, {
        caller: this,
      })
      this.$bvModal.hide('node-common')
      await this.refreshWorld()
    },
    next(id) {
      this.baseNode = id
      this.mode = modeNext
      this.$bvModal.show('node-common')
    },
    call(id) {
      this.baseNode = id
      this.mode = modeCall
      this.$bvModal.show('node-common')
    },
    mouseover() {
      console.log("mouseover")
    },
    drawLine() {
      let that = this
      // let plumbInstance = jsPlumb.getInstance()
      that.plumbInstance.ready(function () {
        that.nodeRelations.forEach(nodeRelation => function () {
          if (nodeRelation.child !== 0) {
            that.plumbInstance.connect({
              source: 'node-' + nodeRelation.node,
              target: 'node-' + nodeRelation.child,
              anchor: ['Right', 'Left'],
              endpoint: 'Blank',
              // connector: ['Flowchart'],
              connector: ['Straight'],
              overlays: [['Arrow', {width: 16, length: 16, location: 1}]],
              paintStyle: {stroke: '#909393', strokeWidth: 2}
            })
          }
          if (nodeRelation.next !== 0) {
            that.plumbInstance.connect({
              source: 'node-' + nodeRelation.node,
              target: 'node-' + nodeRelation.next,
              anchor: ['Bottom', 'Top'],
              endpoint: 'Blank',
              // connector: ['Flowchart'],
              connector: ['Straight'],
              overlays: [['Arrow', {width: 16, length: 16, location: 1}]],
              paintStyle: {stroke: '#909393', strokeWidth: 2}
            })
          }
        }())
      })
    },
    async listNodes(id) {
      this.nodeLoading = true
      let nodes = await lightweightRestful.api.get(consts.api.v1.node.list(id), null, {
        caller: this,
        success_msg: 'list node matrix successfully'
      })
      nodes.forEach(node => {
        this.nodesMap[node.ID]=node
      })
      this.nodeLoading = false
    },
    async listNodeRelationsByRoot(id) {
      this.nodeRelationsLoading = true
      this.nodeRelations = await lightweightRestful.api.get(consts.api.v1.node_relation.list_by_root(id), null, {
        caller: this,
        success_msg: 'list node_relation successfully'
      })
      this.nodeRelations.forEach(nodeRelation => {
        this.nodeRelationsMap[nodeRelation.node]=nodeRelation
      })
      this.nodeRelationsLoading = false
    },
  },
}
</script>

<style scoped>
.hidden {
  visibility: hidden;
}

.node-anchor {
  display: flex;
  position: absolute;
  width: 20px;
  height: 20px;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  cursor: crosshair;
  z-index: 9999;
  background: -webkit-radial-gradient(sandybrown 10%, white 30%, #9a54ff 60%);
}

.anchor-top {
  top: 20px;
  left: 50%;
  margin-left: 20px;
}

/deep/ .modal-dialog {
  width: 90%;
  max-width: 100%;
  height: 90%;
}

/deep/ .modal-content {
  height: 100%;
}

/*/deep/ code {*/
/*  display: none;*/
/*}*/
</style>
