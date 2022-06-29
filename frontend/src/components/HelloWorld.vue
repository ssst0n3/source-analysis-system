<template>
  <div class="hello">
    <AnalysisItem/>
    <MarkdownCard/>
    <NodeItem/>
  </div>
</template>

<script>
import AnalysisItem from "@/components/Analysis/AnalysisItem";
import MarkdownCard from "@/components/Analysis/MarkdownCard";
import NodeItem from "@/components/NodeItem";
export default {
  name: 'HelloWorld',
  components: {NodeItem, MarkdownCard, AnalysisItem},
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
