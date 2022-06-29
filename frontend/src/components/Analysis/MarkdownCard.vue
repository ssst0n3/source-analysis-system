<template>
  <div>
    <b-card class="col-3">
      <div v-html="markdownToHtml"></div>
    </b-card>
  </div>
</template>

<script>
import {marked} from 'marked';
export default {
  name: "MarkdownCard",
  data() {
    return {
      markdown:  `# Hello World
添加IPC挂载点至mount列表。

https://github.com/moby/moby/blob/v20.10.14/container/container_unix.go#L200

\`\`\`
func WithMounts(daemon *Daemon, c *container.Container) coci.SpecOpts {
return func(ctx context.Context, _ coci.Client, _ *containers.Container, s *coci.Spec) (err error) {
  ...
  if !c.HostConfig.IpcMode.IsPrivate() && !c.HostConfig.IpcMode.IsEmpty() {
    ms = append(ms, c.IpcMounts()...)
  }
  ...
}
\`\`\`
`,
    };
  },
  computed: {
    markdownToHtml(){
      return marked(this.markdown);
    }
  }
}
</script>

<style scoped>

</style>
