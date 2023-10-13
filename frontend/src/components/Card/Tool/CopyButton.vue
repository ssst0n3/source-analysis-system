<template>
  <b-badge pill :variant="copied ? 'success' : 'info'"
           v-clipboard:copy="node.markdown"
           v-clipboard:success="success">
    <b-link class="text-white">
      <b-icon-clipboard-check v-if="copied"/>
      <b-icon-clipboard v-else/>
    </b-link>
  </b-badge>
</template>

<script>
export default {
  name: "CopyButton",
  props: {
    node: Object,
  },
  data() {
    return {
      copied: false,
    }
  },
  methods: {
    success() {
      this.$bvToast.toast("copied " + this.node.markdown.slice(0, 10) + '...', {
            title: "Copy " + this.node.ID,
            variant: 'success'
          }
      )
    }
  }
}
</script>

<style scoped>

</style>
