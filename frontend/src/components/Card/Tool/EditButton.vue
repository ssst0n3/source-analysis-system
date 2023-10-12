<template>
<span>
    <b-badge pill variant="info" @click.stop="$emit('save', {update: true})">
      <b-link class="text-white">
        <b-icon-pencil/>
      </b-link>
    </b-badge>
</span>
</template>

<script>

export default {
  name: "EditButton",
  data() {
    return {}
  },
  methods: {}
}
</script>

<style scoped>
</style>
