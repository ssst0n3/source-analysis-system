<template>
  <b-button class="download-btn" @click="downloadDataSource">
    <b-icon icon="download"/>
  </b-button>
</template>

<script>

export default {
  name: "DownloadData",
  props: {
    nodeMatrix: Array,
    toc: Object,
  },
  data() {
    return {

    }
  },
  methods: {
    downloadDataSource() {
      let data = {
        "matrix": this.nodeMatrix,
        "toc": this.toc,
      }
      let marshaled = JSON.stringify(data)
      const urlBlob = window.URL.createObjectURL(new Blob([marshaled]))
      let fileLink = document.createElement('a');
      fileLink.href = urlBlob;
      fileLink.setAttribute('download', 'data-source.json');
      document.body.appendChild(fileLink);
      fileLink.click();
    },
  }
}
</script>

<style scoped>
.download-btn {
  background-color: rgba(0, 0, 0, 0.2);
  border: none;
}
</style>