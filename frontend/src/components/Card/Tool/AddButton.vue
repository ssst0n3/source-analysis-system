<template>
  <span>
    <b-badge pill variant="info" @click.stop="add(directions.up)">
      <b-link class="text-white">
        <b-icon-arrow-bar-up/>
      </b-link>
    </b-badge>
    <b-badge pill variant="info" @click.stop="add(directions.left)">
      <b-link class="text-white">
        <b-icon-arrow-bar-left/>
      </b-link>
    </b-badge>
  </span>
</template>

<script>
import consts from "@/util/const";

export default {
  name: "AddButton",
  data() {
    return {
      directions: consts.directions,
    }
  },
  methods: {
    add(direction) {
      this.$emit('save', {update: false, direction: direction})
    },
  }
}
</script>

<style scoped>
</style>
