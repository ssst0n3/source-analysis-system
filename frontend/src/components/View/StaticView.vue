<template>
  <div>
    <NodeMatrix static :data-source="`/${dataSource}`"/>
  </div>
</template>

<script>
import NodeMatrix from "@/components/Node/NodeMatrix";
import lightweightRestful from "vue-lightweight_restful";

export default {
  name: "StaticView",
  components: {NodeMatrix},
  data() {
    return {
      dataSource: this.$route.params.data_source,
    }
  },
  created() {
    lightweightRestful.api.initClient('/')
  }
}
</script>

<style scoped>

</style>
