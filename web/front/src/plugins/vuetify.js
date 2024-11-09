// 导入 Vuetify 库和相关样式
import 'vuetify/styles'
import { createVuetify } from 'vuetify'

// 创建并导出 Vuetify 实例
export default createVuetify({
  theme: {
    defaultTheme: 'light',
  },
  breakpoint: {
    mobileBreakpoint: 'sm' // 定义移动端断点，‘sm’ 代表宽度小于等于 960px 时视为移动端
  },
})
