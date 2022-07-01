<template>
  <div>
    <div v-for="(col, index) in nodes" :key="`col-${index}`" style="white-space: nowrap" class="mt-5">
      <div :id="'node-'+node.ID" v-for="(node, index2) in col" :key="`col-${index}-row-${index2}`"
           style="display: inline-block; background-color: azure; width: 100px; height: 100px;"
           class="ml-5" :class="{hidden: node.ID === 0}">
        {{ node.ID }}
      </div>
    </div>
  </div>
</template>

<script>
import consts from "@/util/const";
import {jsPlumb} from "jsplumb";
import lightweightRestful from "vue-lightweight_restful";

export default {
  name: "NodeMatrix",
  data() {
    return {
      root: 1,
      nodes: [],
      nodeRelations: []
    }
  },
  async created() {
    await this.nodeMatrix(this.root)
  },

  async mounted() {
    await this.listNodeRelationsByRoot(this.root)
    this.drawLine()
  },
  methods: {
    drawLine() {
      let that = this
      let plumbInstance = jsPlumb.getInstance()
      plumbInstance.ready(function () {
        console.log("node_relations:", that.nodeRelations)
        that.nodeRelations.forEach(nodeRelation => function () {
          if (nodeRelation.child !== 0) {
            plumbInstance.connect({
              source: 'node-' + nodeRelation.node,
              target: 'node-' + nodeRelation.child,
              anchor: ['Right', 'Left'],
              endpoint: 'Blank',
              connector: ['Flowchart'],
              overlays: [['Arrow', {width: 16, length: 16, location: 1}]],
              paintStyle: {stroke: '#909393', strokeWidth: 2}
            })
          }
          if (nodeRelation.next !== 0) {
            plumbInstance.connect({
              source: 'node-' + nodeRelation.node,
              target: 'node-' + nodeRelation.next,
              anchor: ['Bottom', 'Top'],
              endpoint: 'Blank',
              connector: ['Flowchart'],
              overlays: [['Arrow', {width: 16, length: 16, location: 1}]],
              paintStyle: {stroke: '#909393', strokeWidth: 2}
            })
          }
        }())
      })
    },
    async nodeMatrix(id) {
      this.nodes = await lightweightRestful.api.get(consts.api.v1.node.matrix(id), null, {
        caller: this,
      })
    },
    async listNodeRelationsByRoot(id) {
      this.nodeRelations = await lightweightRestful.api.get(consts.api.v1.node_relation.list_by_root(id), null, {
        caller: this,
      })
    },
  },
}
</script>

<style scoped>
.hidden {
  visibility: hidden;
}
</style>
