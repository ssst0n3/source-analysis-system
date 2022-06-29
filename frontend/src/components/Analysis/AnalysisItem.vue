<template>
  <div>
    <b-card
        tag="article"
        class="mb-2 col-3 "
    >
      <b-card-title>
        WithMounts()
      </b-card-title>
      <b-card-sub-title>
        <b-link href="#" class="card-link">
          <b-icon-github/>
        </b-link>
      </b-card-sub-title>
      <b-card-body>
        <b-card-text>
          {{ note }}
        </b-card-text>
        <pre> {{ code }} </pre>
      </b-card-body>
    </b-card>
  </div>
</template>

<script>
export default {
  name: "AnalysisItem",
  data() {
    return {
      note: "添加IPC挂载点至mount列表。",
      link: "https://github.com/moby/moby/blob/v20.10.14/container/container_unix.go#L200",
      code: `func WithMounts(daemon *Daemon, c *container.Container) coci.SpecOpts {
  return func(ctx context.Context, _ coci.Client, _ *containers.Container, s *coci.Spec) (err error) {
    ...
    if !c.HostConfig.IpcMode.IsPrivate() && !c.HostConfig.IpcMode.IsEmpty() {
    ms = append(ms, c.IpcMounts()...)
  }
  ...
}`,
    }
  }
}
</script>

<style scoped>

</style>
