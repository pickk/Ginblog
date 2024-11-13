<template>
  <v-footer padless color="indigo darken-2">
    <v-row justify="center" no-gutters>
      <v-col class="py-4 text-center white--text" cols="12">{{ new Date().getFullYear() }} - GinBlog</v-col>

      <div class="text-center white--text">
        <a class="text-center white--text" href="https://beian.miit.gov.cn/">{{icp_record}}</a>
      </div>
    </v-row>
  </v-footer>
</template>

<script>
import { eventBus } from '@/eventBus'

export default {
  name: 'AFooter',
  data() {
    return {
      icp_record: '',
    }
  },
  mounted() {
    // 使用 eventBus.on 来监听 'msg' 事件
    eventBus.on('msg', (msg) => {
      this.icp_record = msg
    })
  },
  beforeUnmount() {
    // 在组件卸载之前移除监听器，防止内存泄漏
    eventBus.off('msg')
  },
}
</script>

<style>
</style>