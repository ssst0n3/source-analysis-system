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
        <div v-if="node.ID !== -1" @dblclick="edit(node)">
          <MarkdownCard style="white-space: normal" :id="'card-'+node.ID"
                        :markdown="node.markdown.toString()" :nodeId="node.ID.toString()"
                        :next-id="node.next" :child-id="node.child" :last-id="node.last" :parent-id="node.parent"
                        :has-parent="node.parent !== undefined"
                        :static-view="staticView"
                        :active="focus===node.ID"
                        :size="size"
                        v-on:navi="navi" v-on:add="add" v-on:remove="remove"
                        v-if="node.markdown.length>0"/>
          <AnalysisItem v-else/>
        </div>
      </div>
    </div>
    <b-modal id="node-common" hide-footer size="xl" v-if="!staticView">
      <MarkdownEditor ref="markdown_editor_common"
                      :markdown="markdown" :size="size"
                      v-on:save="save"
      />
    </b-modal>
    <b-modal id="delete-prompt" @ok="deleteWithMode">
      <b-form-group label="delete mode" v-slot="{ ariaDescribedby }">
        <b-form-radio v-model="modelDelete.mode" :aria-describedby="ariaDescribedby" :value="1">
          <b-icon-arrow-left-circle-fill variant="dark"/>
          move children nodes left
        </b-form-radio>
        <b-form-radio v-model="modelDelete.mode" :aria-describedby="ariaDescribedby" :value="2">
          <b-icon-arrow-up-circle-fill/>
          move next nodes up
        </b-form-radio>
        <b-form-radio v-model="modelDelete.mode" :aria-describedby="ariaDescribedby" :value="3">
          <b-icon-x-circle-fill/>
          discard all children and next nodes
        </b-form-radio>
      </b-form-group>
    </b-modal>
  </div>
</template>

<script>
import {Matrix} from "@/util/matrix/matrix";
import consts from "@/util/const";
import {jsPlumb} from "jsplumb";
import lightweightRestful from "vue-lightweight_restful";
import MarkdownCard from "@/components/Analysis/MarkdownCard";
import AnalysisItem from "@/components/Analysis/AnalysisItem";
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";
import {anchor} from "@/util/util";
import {update, update_node_relation} from "@/util/nodeRelation";
import {createNode, updateNode} from "@/util/node";
import ToolBar from "@/components/Tool/ToolBar.vue";

export default {
  name: "NodeMatrix",
  components: {
    ToolBar, AnalysisItem, MarkdownCard, MarkdownEditor
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
      mode: 0,
      mode_edit: false,
      baseNode: 0,
      markdown: "",
      plumbInstance: null,
      nodeLoading: false,
      nodeRelationsLoading: false,
      toc: {
        headings: [],
        title: null,
      },
      time_start: 0,
      focus: 0,
      size: 'small',
      modelDelete: {
        node: -1,
        mode: 0,
      }
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
    dfsOptionUpdate(checked) {
      this.dfs = checked
      this.refreshWorld()
    },
    focusNode(id) {
      this.focus = id
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
    edit(node) {
      this.markdown = node.markdown
      this.baseNode = node.ID
      this.mode_edit = true
      this.$bvModal.show('node-common')
    },
    resetMarkdownEditor() {
      this.mode = 0
      this.markdown = ""
      this.mode_edit = false
      this.$bvModal.hide('node-common')
    },
    async save() {
      let content = this.$refs.markdown_editor_common.edit
      if (this.mode_edit) {
        await updateNode(this, this.baseNode, content)
      } else {
        let id = await createNode(this, content)
        await update_node_relation(this, this.mode, this.root, this.baseNode, id, this.nodesMap)
      }
      this.resetMarkdownEditor()
      await this.refreshWorld()
    },
    navi(nodeId, navId, direction) {
      if (direction === consts.directions.left) {
        navId = this.matrix.caller(nodeId)
      }
      anchor('card-' + navId)
      this.focus = navId
    },
    add(nodeId, direction) {
      this.resetMarkdownEditor()
      this.baseNode = parseInt(nodeId)
      this.mode = direction
      this.$bvModal.show('node-common')
    },
    remove(id) {
      this.modelDelete.node = id
      this.$bvModal.show('delete-prompt')
    },
    deleteWithMode() {
      let node = this.nodesMap[this.modelDelete.node]
      console.log("node:", node)
      let data = {}
      let base = node.parent
      let direction = 'child'
      if (base === undefined) {
        base = node.last
        direction = 'next'
      }
      switch (this.modelDelete.mode) {
        case 1:
          alert("TODO: move children nodes left #74")
          break
        case 2:
          if (node.next === 0) {
            alert("There's no next nodes")
            break
          }
          if (node.child === 0) {
            data[direction] = node.next
            update(this, base, data)
            this.hideNode(this.modelDelete.node);
          } else {
            alert("TODO: move next nodes up, discard/preserve child nodes #77")
          }
          break
        case 3:
          this.hideNode(this.modelDelete.node);
      }
      this.modelDelete = {
        node: 0,
        mode: 0,
      }
    },
    mouseover() {
      // console.log("mouseover")
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
              that.plumbInstance.connect({
                source: 'node-' + node.ID,
                target: 'node-' + node.child,
                anchor: ['Right', 'Left'],
                endpoint: 'Blank',
                // connector: ['Flowchart'],
                connector: ['Straight'],
                overlays: [['Arrow', {width: 16, length: 16, location: 1}]],
                paintStyle: {stroke: '#909393', strokeWidth: 2},
                // hoverPaintStyle: {
                //   outlineStroke: 'lightblue'
                // },
              })
            }
            if (node.next !== 0) {
              that.plumbInstance.connect({
                source: 'node-' + node.ID,
                target: 'node-' + node.next,
                anchor: ['Bottom', 'Top'],
                endpoint: 'Blank',
                // connector: ['Flowchart'],
                connector: ['Straight'],
                overlays: [['Arrow', {width: 16, length: 16, location: 1}]],
                paintStyle: {stroke: '#909393', strokeWidth: 2},
              })
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
