<template>
  <div>
    <div style="white-space: nowrap">
      <div class="" style="display: inline-block;" id="div1">
        <AnalysisItem id="card1-1" style="width: 1000px; white-space: normal"/>
      </div>
      <div class="ml-5" style="display: inline-block;" id="div2">
        <MarkdownCard id="card1-2" style="width: 1000px; white-space: normal"/>
      </div>
      <div class="ml-5" style="display: inline-block;" id="div2">
        <MarkdownCard id="card1-3" style="width: 1000px; white-space: normal"/>
      </div>
    </div>
    <div class="row mt-5">
      <AnalysisItem id="card3" style="width: 1000px"/>
      <MarkdownCard id="card4" class="ml-5" style="width: 1000px"/>
    </div>
  </div>
</template>

<script>
import * as d3 from 'd3';
import {jsPlumb} from "jsplumb"
import AnalysisItem from "@/components/Analysis/AnalysisItem";
import MarkdownCard from "@/components/Analysis/MarkdownCard";

export default {
  name: "NodeItem",
  components: {MarkdownCard, AnalysisItem},
  data() {
    return {
      data: [99, 71, 78]
    }
  },
  mounted() {
    let plumbInstance = jsPlumb.getInstance()
    plumbInstance.ready(function () {
      plumbInstance.connect({
        source: 'card1-1',
        target: 'card1-2',
        anchor: ['Right', 'Left'],
        // anchor: ['Left', 'Right', 'Top', 'Bottom', [0.3, 0, 0, -1], [0.7, 0, 0, -1], [0.3, 1, 0, 1], [0.7, 1, 0, 1]],
        endpoint: 'Blank',
        connector: ['Flowchart'],
        overlays: [['Arrow', {width: 16, length: 16, location: 1}]],
        paintStyle: {stroke: '#909393', strokeWidth: 5}
      })
      plumbInstance.connect({
        source: 'card1-1',
        target: 'card3',
        anchor: ['Bottom', 'Top'],
        endpoint: '',
        connector: ['Straight'],
        overlays: [['Arrow', {width: 16, length: 16, location: 1}]],
        paintStyle: {stroke: '#909393', strokeWidth: 5}
      })
    })
    // const svg = d3.select(this.$el)
    //     .append('svg')
    //     .attr('width', 500)
    //     .attr('height', 270)
    //     .append('g')
    //     .attr('transform', 'translate(0, 10)');
    // d3.select(this.$el).append('div').html('<p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>');
    d3.select(this.$el).append("line")
        .attr("x1", this.$refs.div1.getBoundingClientRect().left + this.$refs.div1.getBoundingClientRect().width / 2)
        .attr("y1", this.$refs.div1.getBoundingClientRect().bottom)
        .attr("x2", this.$refs.div2.getBoundingClientRect().left + this.$refs.div2.getBoundingClientRect().width / 2)
        .attr("y2", this.$refs.div2.getBoundingClientRect().top)
        .attr("stroke", "black")
        .attr("stroke-width", "2px");
    this.$refs.line.setAttribute("x1", this.$refs.div1.getBoundingClientRect().left + this.$refs.div1.getBoundingClientRect().width / 2)
    this.$refs.line.setAttribute("y1", this.$refs.div1.getBoundingClientRect().bottom)
  },
}
</script>

<style scoped>

</style>
