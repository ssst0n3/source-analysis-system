<template>
  <span>
    <b-badge pill variant="info" @click.stop="$bvModal.show('node-remove-prompt')">
      <b-link class="text-white">
        <b-icon-x></b-icon-x>
      </b-link>
    </b-badge>
  </span>
</template>

<script>

export default {
  name: "RemoveButton",
}
</script>

<style scoped>

</style>
