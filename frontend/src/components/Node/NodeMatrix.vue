<template>
  <div>
    <b-button class="download-btn" @click="downloadDataSource" v-if="!staticView">
      <b-icon icon="download"/>
    </b-button>
    <TableOfContent :toc="toc"/>
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
          <b-tooltip v-if="false" offset="-200" boundary="document" placement="top" :target="'node-'+node.ID"
                     variant="light"
                     triggers="hover">
            <div>
              <!--              <b-btn variant="light" v-if="node.next !== 0">-->
              <!--                <a @click.prevent="anchor('card-'+node.next)" :href="'#card-'+node.next">Next</a>-->
              <!--              </b-btn>-->
              <!--              <b-btn @click="next(node.ID)" v-else-if="!staticView">Next</b-btn>-->
              <!--              <b-btn variant="light" class="ml-2" v-if="node.child !==0">-->
              <!--                <a @click.prevent="anchor('card-'+node.child)" :href="'#card-'+node.child">Call</a>-->
              <!--              </b-btn>-->
              <!--              <b-btn @click="call(node.ID)" class="ml-2" v-else-if="!staticView">Call</b-btn>-->
              <!--              <b-btn variant="light" class="ml-2" v-if="node.parent !==undefined">-->
              <!--                <a @click.prevent="anchor('card-'+node.parent)" :href="'#card-'+node.parent">Parent</a>-->
              <!--              </b-btn>-->
              <b-btn v-if="!staticView" class="ml-2">
                <a @click="unlinkNodeFromParent(node.ID)">Delete</a>
              </b-btn>
            </div>
          </b-tooltip>
          <!--          <b-tooltip boundary="document" placement="bottom" :target="'node-'+node.ID" variant="light"-->
          <!--                     triggers="hover">-->
          <!--            <div>-->
          <!--              <b-btn size="sm" variant="light" v-if="node.next !== 0">-->
          <!--                <a @click.prevent="anchor('card-'+node.next)" :href="'#card-'+node.next">Next</a>-->
          <!--              </b-btn>-->
          <!--              <b-btn size="sm" @click="next(node.ID)" v-else-if="!staticView">Next</b-btn>-->
          <!--            </div>-->
          <!--          </b-tooltip>-->
          <!--          <b-tooltip boundary="document" placement="righttop" :target="'node-'+node.ID" variant="light"-->
          <!--                     triggers="hover">-->
          <!--            <div>-->
          <!--              <b-btn size="sm" variant="light" class="ml-2" v-if="node.child !== 0">-->
          <!--                <a @click.prevent="anchor('card-'+node.child)" :href="'#card-'+node.child">Call</a>-->
          <!--              </b-btn>-->
          <!--              <b-btn size="sm" @click="call(node.ID)" class="ml-2" v-else-if="!staticView">Call</b-btn>-->
          <!--            </div>-->
          <!--          </b-tooltip>-->
          <b-tooltip v-if="false" boundary="document" placement="lefttop" :target="'node-'+node.ID" variant="light"
                     triggers="hover">
            <div>
              <b-btn size="sm" variant="light" class="ml-2" v-if="node.parent !==undefined">
                <a @click="anchor('card-'+node.parent)" :href="'#card-'+node.parent">Parent</a>
              </b-btn>
            </div>
          </b-tooltip>
          <b-tooltip v-if="false" boundary="document" placement="top" :target="'node-'+node.ID" variant="light"
                     triggers="hover">
            <div>
              <b-btn size="sm">
                INSERT
              </b-btn>
            </div>
          </b-tooltip>
          <MarkdownCard style="white-space: normal" :id="'card-'+node.ID"
                        :markdown="node.markdown.toString()" :nodeId="node.ID.toString()"
                        :next-id="node.next" :child-id="node.child" :last-id="node.last" :parent-id="node.parent"
                        :has-parent="node.parent !== undefined"
                        :static-view="staticView"
                        :active="focus===node.ID"
                        v-on:update_node="refreshWorld"
                        v-on:navi="navi" v-on:remove="remove"
                        v-on:next="next" v-on:call="call" v-on:insert="insert" v-on:caller="caller"
                        v-if="node.markdown.length>0"/>
          <AnalysisItem v-else/>
        </div>
      </div>
    </div>
    <b-modal id="node-common" hide-footer size="xl" v-if="!staticView">
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
const modeInsertLast = 3
const modeInsertCaller = 4

import consts from "@/util/const";
import {jsPlumb} from "jsplumb";
import lightweightRestful from "vue-lightweight_restful";
import MarkdownCard from "@/components/Analysis/MarkdownCard";
import AnalysisItem from "@/components/Analysis/AnalysisItem";
import MarkdownEditor from "@/components/Markdown/MarkdownEditor";
import TableOfContent from "@/components/TableOfContent/TableOfContent";
import {anchor} from "@/util/util";

export default {
  name: "NodeMatrix",
  components: {TableOfContent, AnalysisItem, MarkdownCard, MarkdownEditor},
  props: {
    root: Number,
    staticView: Boolean,
    dataSource: String,
  },
  data() {
    return {
      nodesMap: {},
      nodeRelationsMap: {},
      nodeMatrix: [],
      mode: 0,
      baseNode: 0,
      plumbInstance: null,
      nodeLoading: false,
      nodeRelationsLoading: false,
      toc: {
        headings: [],
        title: null,
      },
      time_start: 0,
      focus: 0,
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
    }
  },
  methods: {
    focusNode(id) {
      this.focus = id
    },
    // async unlinkNodeFromParent(id) {
    //   await lightweightRestful.api.post(consts.api.v1.node_relation.unlink(id), null, null, {
    //     caller: this,
    //   })
    // },
    async hideNode(id) {
      await lightweightRestful.api.delete(consts.api.v1.node_relation.hide_node(id), null, {
        caller: this,
        success_msg: 'delete node successfully'
      })
      await this.refreshData()
    },
    anchor: anchor,
    downloadDataSource() {
      let data = {
        "matrix": this.nodeMatrix,
        "toc": this.toc,
      }
      let marshaled = JSON.stringify(data)
      const urlBlob = window.URL.createObjectURL(new Blob([marshaled]))
      let fileLink = document.createElement('a');
      fileLink.href = urlBlob;
      fileLink.setAttribute('download', 'data-source.json');
      document.body.appendChild(fileLink);
      fileLink.click();
    },
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
      let matrix = new Matrix.Matrix(this.root, this.nodesMap, this.nodeRelationsMap)
      matrix.childRecursive(0)
      matrix.shift()
      matrix.cleanSuffix()
      this.nodeMatrix = matrix.dumpNode()
      matrix.generateToc(this.root)
      this.toc = matrix.toc
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
      await this.update_node_relation(id)
      this.$bvModal.hide('node-common')
      await this.refreshWorld()
    },
    async update_node_relation(id) {
      switch (this.mode) {
        case modeCall:
        case modeNext:
          await this.update_node_relation_for_next_and_call(id)
          break
        case modeInsertLast:
          await this.update_node_relation_for_insert(id)
          break
      }
    },
    async update_node_relation_for_insert(id) {
      await lightweightRestful.api.post(consts.api.v1.node_relation.node_relation, null, {
        root: this.root,
        node: id,
        next: this.baseNode
      }, {
        caller: this,
      })

      let baseNode = this.nodesMap[this.baseNode]
      let parent = baseNode.parent
      let last = baseNode.last
      if (parent !== undefined) {
        await lightweightRestful.api.put(consts.api.v1.node_relation.update_node_relation_by_node(parent), null, {
          child: id,
        }, {
          caller: this,
        })
      }
      if (last !== undefined) {
        await lightweightRestful.api.put(consts.api.v1.node_relation.update_node_relation_by_node(last), null, {
          next: id,
        }, {
          caller: this,
        })
      }
    },
    async update_node_relation_for_next_and_call(id) {
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
    },
    next(id, nextId) {
      if (nextId !== 0) {
        anchor('card-' + nextId)
        this.focus = nextId
      } else {
        this.baseNode = parseInt(id)
        this.mode = modeNext
        this.$bvModal.show('node-common')
      }
    },
    call(id, childId) {
      if (childId !== 0) {
        anchor('card-' + childId)
        this.focus = childId
      } else {
        this.baseNode = parseInt(id)
        this.mode = modeCall
        this.$bvModal.show('node-common')
      }
    },
    navi(id) {
      anchor('card-' + id)
      this.focus = id
    },
    insert(id) {
      this.baseNode = parseInt(id)
      this.mode = modeInsertLast
      this.$bvModal.show('node-common')
    },
    caller(id) {
      this.baseNode = parseInt(id)
      this.mode = modeInsertCaller
      this.$bvModal.show('node-common')
    },
    remove(id) {
      if (confirm(`are you sure to delete #${id}? all of it's children will be deleted`)) {
        this.hideNode(id)
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
      let nodeRelations = await lightweightRestful.api.get(consts.api.v1.node_relation.item(id), null, {
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

/*.node-anchor {*/
/*  display: flex;*/
/*  position: absolute;*/
/*  width: 20px;*/
/*  height: 20px;*/
/*  align-items: center;*/
/*  justify-content: center;*/
/*  border-radius: 10px;*/
/*  cursor: crosshair;*/
/*  z-index: 9999;*/
/*  background: -webkit-radial-gradient(sandybrown 10%, white 30%, #9a54ff 60%);*/
/*}*/

/*.anchor-top {*/
/*  top: 20px;*/
/*  left: 50%;*/
/*  margin-left: 20px;*/
/*}*/

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

.download-btn {
  position: fixed;
  bottom: 20px;
  right: 100px;
  background-color: rgba(0, 0, 0, 0.2);
  border: none;
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
