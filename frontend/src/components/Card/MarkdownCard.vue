<template>
  <div>
    <b-card :border-variant="active ? 'info' : ''" header-bg-variant="light">
      <template #header>
        <b-row>
          <b-col class="col-2">
            <span v-show="active"> # {{ node.ID }}</span>
          </b-col>
          <b-col class="text-right">
            <b-badge comment="place holder, do not delete me" pill variant="light">
              <b-link class="text-white">
                <b-icon-arrow-up/>
              </b-link>
            </b-badge>
            <span v-if="active">
              <ToolBar :root="root" :static-view="staticView" :node="node" :size="size"
                       v-on:navi="$emit('navi', $event)"
                       v-on:refresh="$emit('refresh')"
                       v-on:refreshWorld="$emit('refreshWorld')"
              />
            </span>
          </b-col>
        </b-row>
      </template>
      <MarkdownViewer :markdown="node.markdown" :size="size" style="width: 100%"/>
    </b-card>
  </div>
</template>

<script>
import MarkdownViewer from "@/components/Markdown/MarkdownViewer.vue";
import ToolBar from "@/components/Card/Tool/ToolBar";

export default {
  name: "MarkdownCard",
  components: {ToolBar, MarkdownViewer},
  props: {
    root: Number,
    node: Object,
    staticView: Boolean,
    active: Boolean,
    size: String,
  },
  data() {
    return {}
  },
  methods: {}
}
</script>

<style scoped>

</style>
