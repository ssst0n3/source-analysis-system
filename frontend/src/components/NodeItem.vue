<template>
  <div>
    <div ref="div1" style="background-color: #42b983; width: 300px; height: 100px" class="ml-5"></div>
    <svg>
    <line ref="line" x1="0" y1="0" x2="200" y2="50" stroke="#666666" stroke-width="2" marker-end="url(#arrow)" stroke-dasharray="10,10"></line>
    </svg>
    <div ref="div2" style="background-color: azure; width: 300px; height: 100px" class="ml-5 mt-5"></div>
  </div>
</template>

<script>
import * as d3 from 'd3';

export default {
  name: "NodeItem",
  data() {
    return {
      data: [99, 71, 78]
    }
  },
  mounted() {
    // const svg = d3.select(this.$el)
    //     .append('svg')
    //     .attr('width', 500)
    //     .attr('height', 270)
    //     .append('g')
    //     .attr('transform', 'translate(0, 10)');
    d3.select(this.$el).append('div').html('<p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>');
    d3.select(this.$el).append("line")
        .attr("x1", this.$refs.div1.getBoundingClientRect().left+this.$refs.div1.getBoundingClientRect().width/2)
        .attr("y1", this.$refs.div1.getBoundingClientRect().bottom)
        .attr("x2", this.$refs.div2.getBoundingClientRect().left+this.$refs.div2.getBoundingClientRect().width/2)
        .attr("y2", this.$refs.div2.getBoundingClientRect().top)
        .attr("stroke", "black")
        .attr("stroke-width", "2px");
    this.$refs.line.setAttribute("x1", this.$refs.div1.getBoundingClientRect().left+this.$refs.div1.getBoundingClientRect().width/2)
    this.$refs.line.setAttribute("y1", this.$refs.div1.getBoundingClientRect().bottom)
  },
}
</script>

<style scoped>

</style>
