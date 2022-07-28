<template>
  <div>
    <NodeMatrix :root="root"/>
  </div>
</template>

<script>
import NodeMatrix from "@/components/Node/NodeMatrix";

export default {
  name: "DynamicView",
  components: {NodeMatrix},
  data() {
    return {
      root: parseInt(this.$route.params.id),
    }
  }
}
</script>

<style scoped>

</style>
