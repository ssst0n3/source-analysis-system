<template functional>
  <b-badge pill variant="info" @click.stop="props.save({update: true})">
    <b-link class="text-white">
      <b-icon-pencil/>
    </b-link>
  </b-badge>
</template>
