<template>
  <span>
    <b-badge pill variant="info" @click.stop="trigger">
      <b-link class="text-white">
        <b-icon-x></b-icon-x>
      </b-link>
    </b-badge>
    <b-modal :id="`delete-prompt-${node.ID}`" lazy @ok="remove">
      <b-form-group label="delete mode" v-slot="{ ariaDescribedby }">
        <b-form-radio v-model="mode" :aria-describedby="ariaDescribedby" :value="1">
          <b-icon-arrow-left-circle-fill variant="dark"/>
          move children nodes left
        </b-form-radio>
        <b-form-radio v-model="mode" :aria-describedby="ariaDescribedby" :value="2">
          <b-icon-arrow-up-circle-fill/>
          move next nodes up
        </b-form-radio>
        <b-form-radio v-model="mode" :aria-describedby="ariaDescribedby" :value="3">
          <b-icon-x-circle-fill/>
          discard all children and next nodes
        </b-form-radio>
      </b-form-group>
    </b-modal>
  </span>
</template>

<script>
import {update} from "@/util/nodeRelation";
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/util/const";

export default {
  name: "RemoveButton",
  props: {
    node: Object,
  },
  data() {
    return {
      mode: 0,
    }
  },
  methods: {
    trigger() {
      this.$bvModal.show(`delete-prompt-${this.node.ID}`)
    },
    remove() {
      let data = {}
      let base = this.node.parent
      let direction = 'child'
      if (base === undefined) {
        base = this.node.last
        direction = 'next'
      }
      switch (this.mode) {
        case 1:
          alert("TODO: move children nodes left #74")
          break
        case 2:
          if (this.node.next === 0) {
            alert("There's no next nodes")
            break
          }
          if (this.node.child === 0) {
            data[direction] = this.node.next
            update(this, base, data)
            this.hideNode(this.node.ID);
          } else {
            alert("TODO: move next nodes up, discard/preserve child nodes #77")
          }
          break
        case 3:
          this.hideNode(this.node.ID);
      }
      this.mode = 0
    },
    async hideNode(id) {
      await lightweightRestful.api.delete(consts.api.v1.node_relation.hide_node(id), null, {
        caller: this,
        success_msg: 'delete node successfully'
      })
      await this.$emit('refresh')
    },
  }
}
</script>

<style scoped>

</style>
