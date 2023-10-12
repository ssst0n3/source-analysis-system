<template>
  <div>
    <ToolBar :toc="toc" :static-view="staticView" :node-matrix="nodeMatrix" :nodes-count="nodesCount"
             v-on:size_update="(s)=>{size=s}"
             v-on:dfs-option-update="dfsOptionUpdate"/>
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
           @click="focusNode(node.ID)"
           :ref="'node-'+node.ID" :id="'node-'+node.ID" v-for="(node, index2) in col"
           :key="`col-${index}-row-${index2}`"
           style="display: inline-block; width: 500px; vertical-align:middle"
           class="ml-5" :class="{hidden: node.ID === 0}">
        <div v-if="node.ID !== -1">
          <MarkdownCard style="white-space: normal" :id="'card-'+node.ID"
                        :static-view="staticView" :active="focus===node.ID" :size="size" :node="node"
                        v-on:save="save" v-on:navi="navi"
          />
        </div>
      </div>
    </div>
    <NodeEditor :static-view="staticView" :node="activeNode" :size="size" :model="editorModel" :root="root"
                v-on:refresh="refreshData" v-on:refreshWorld="refreshWorld"/>
    <NodeRemove :node="activeNode"
                v-on:refresh="refreshData"/>
  </div>
</template>

<script>
import {Matrix} from "@/util/matrix/matrix";
import consts from "@/util/const";
import {jsPlumb} from "jsplumb";
import lightweightRestful from "vue-lightweight_restful";
import MarkdownCard from "@/components/Card/MarkdownCard";
import {anchor} from "@/util/util";
import ToolBar from "@/components/Tool/ToolBar.vue";
import NodeEditor from "@/components/Node/NodeEditor";
import NodeRemove from "@/components/Node/NodeRemove";

export default {
  name: "NodeMatrix",
  components: {
    NodeRemove,
    NodeEditor,
    ToolBar, MarkdownCard
  },
  props: {
    root: Number,
    staticView: Boolean,
    dataSource: String,
  },
  data() {
    return {
      dfs: false,
      nodesMap: {},
      nodeRelationsMap: {},
      matrix: {},
      nodeMatrix: [],
      nodesCount: 0,
      plumbInstance: null,
      nodeLoading: false,
      nodeRelationsLoading: false,
      toc: {
        headings: [],
        title: null,
      },
      time_start: 0,
      focus: 0,
      activeNode: {markdown: ''},
      editorModel: {
        update: false,
        direction: 0,
      },
      size: 'small',
    }
  },
  computed: {
    loading() {
      return this.nodeLoading || this.nodeRelationsLoading
    }
  },
  async created() {
    this.time_start = new Date().getTime()
    await this.refreshData()
  },
  mounted() {
    this.plumbInstance = jsPlumb.getInstance()
    this.plumbInstance.bind('click', function (conn) {
      // console.log(conn, event)
      window.scrollTo(conn.target.offsetLeft - 100, conn.target.offsetTop - 100)
    })
    // this.$nextTick(() => {
    //   console.log(document.getElementsByTagName('h1'))
    // })
    // await this.listNodeRelationsByRoot(this.root)
  },
  destroyed() {
    this.plumbInstance.deleteEveryConnection()
    this.plumbInstance.deleteEveryEndpoint()
  },
  watch: {
    nodeMatrix: async function () {
      this.$nextTick(() => {
        this.drawLine()
      })
    },
  },
  methods: {
    save(model) {
      this.editorModel = model
      this.$bvModal.show('node-common-modal')
    },
    dfsOptionUpdate(checked) {
      this.dfs = checked
      this.refreshWorld()
    },
    focusNode(id) {
      this.focus = id
      this.activeNode = this.nodesMap[id]
    },
    async hideNode(id) {
      await lightweightRestful.api.delete(consts.api.v1.node_relation.hide_node(id), null, {
        caller: this,
        success_msg: 'delete node successfully'
      })
      await this.refreshData()
    },
    anchor: anchor,
    async refreshData() {
      if (this.staticView) {
        let dataSource = await lightweightRestful.api.get(this.dataSource, null, {
          caller: this,
          success_msg: 'list node matrix successfully'
        })
        this.nodeMatrix = dataSource["matrix"]
        this.toc = dataSource["toc"]
        return
      }
      console.log('time before data pulled is', `${new Date().getTime() - this.time_start}ms`)
      await this.listNodes(this.root)
      await this.listNodeRelationsByRoot(this.root)
      console.log('time after data pulled is', `${new Date().getTime() - this.time_start}ms`)
      this.matrix = new Matrix(this.root, this.nodesMap, this.nodeRelationsMap, this.dfs)
      // this.matrix.childRecursive(0)
      this.matrix.shift()
      this.nodesCount = this.matrix.count()
      // this.matrix.cleanSuffix()
      this.nodeMatrix = this.matrix.dumpNode()
      this.matrix.generateToc(this.root)
      this.toc = this.matrix.toc
      console.log('time after generate matrix', `${new Date().getTime() - this.time_start}ms`)
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
    navi(params) {
      let nodeId = params[0]
      let navId = params[1]
      let direction = params[2]
      if (direction === consts.directions.left) {
        navId = this.matrix.caller(nodeId)
      }
      anchor('card-' + navId)
      this.focus = navId
    },
    mouseover() {
      // console.log("mouseover")
    },
    async doDrawLine(source, target, anchor) {
      this.plumbInstance.connect({
        source: 'node-' + source,
        target: 'node-' + target,
        anchor: anchor,
        endpoint: 'Blank',
        // connector: ['Flowchart'],
        connector: ['Straight'],
        overlays: [['Arrow', {width: 16, length: 16, location: 1}]],
        paintStyle: {stroke: '#909393', strokeWidth: 2},
        // hoverPaintStyle: {
        //   outlineStroke: 'lightblue'
        // },
      })
    },
    drawLine() {
      let that = this
      // let plumbInstance = jsPlumb.getInstance()
      that.plumbInstance.ready(function () {
        for (let x = 0; x < that.nodeMatrix.length; x++) {
          for (let y = 0; y < that.nodeMatrix[x].length; y++) {
            let node = that.nodeMatrix[x][y]
            if (node.ID <= 0) {
              continue
            }
            if (node.child !== 0) {
              that.doDrawLine(node.ID, node.child, ['Right', 'Left'])
            }
            if (node.next !== 0) {
              that.doDrawLine(node.ID, node.next, ['Bottom', 'Top'])
            }
          }
        }
      })
      console.log('time after line drawn cost is', `${new Date().getTime() - this.time_start}ms`)
    },
    async listNodes(id) {
      this.nodeLoading = true
      let nodes = await lightweightRestful.api.get(consts.api.v1.node.list(id), null, {
        caller: this,
        success_msg: 'list node successfully'
      })
      nodes.forEach(node => {
        this.nodesMap[node.ID] = node
      })
      this.nodeLoading = false
    },
    async listNodeRelationsByRoot(id) {
      this.nodeRelationsLoading = true
      let nodeRelations = await lightweightRestful.api.get(consts.api.v1.node_relation.list(id), null, {
        caller: this,
        success_msg: 'list node_relation successfully'
      })
      nodeRelations.forEach(nodeRelation => {
        this.nodeRelationsMap[nodeRelation.node] = nodeRelation
      })
      this.nodeRelationsLoading = false
    },
  }
  ,
}
</script>

<style scoped>
.hidden {
  visibility: hidden;
}

/*noinspection CssUnusedSymbol*/
/deep/ #node-common .modal-dialog {
  width: 90%;
  max-width: 100%;
  height: 90%;
}

/*noinspection CssUnusedSymbol*/
/deep/ #node-common .modal-content {
  height: 100%;
}

/*/deep/ code {*/
/*  display: none;*/
/*}*/


/*noinspection CssUnusedSymbol*/
/deep/ .tooltip-inner {
  max-width: unset;
  /*background-color: rgba(0,0,0,0);*/
}

</style>
