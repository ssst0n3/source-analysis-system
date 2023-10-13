<template functional>
  <span>
    <b-badge pill variant="info" @click.stop="props.save({update: false, direction: props.directions.up})">
      <b-link class="text-white">
        <b-icon-arrow-bar-up/>
      </b-link>
    </b-badge>
    <b-badge pill variant="info" @click.stop="props.save({update: false, direction: props.directions.left})">
      <b-link class="text-white">
        <b-icon-arrow-bar-left/>
      </b-link>
    </b-badge>
  </span>
</template>
