// 导入 Vue 框架
import Vue from 'vue'

// 导入 Vuetify 库，用于实现 Material Design 风格的 UI 组件
import Vuetify from 'vuetify/lib/framework'

// 导入第三方插件 vue-multiple-message，用于显示全局消息通知
import Message from 'vue-multiple-message'

// 配置 Message 插件的全局设置
Message.config({
  top: 60,         // 设置消息的显示位置距离顶部 60 像素
  duration: 3000,  // 设置消息自动消失的时间为 3000 毫秒（3秒）
  zIndex: 2000     // 设置消息的 z-index，为 2000，以确保消息显示在较前层
})

// 将 Message 作为 Vue 实例的一个全局属性，使用时可以通过 `this.$message` 来调用
Vue.prototype.$message = Message

// 使用 Vue.use() 注册 Vuetify 作为 Vue 插件
Vue.use(Vuetify)

// 导出一个 Vuetify 实例，配置项目中的 Vuetify 设置
export default new Vuetify({
  breakpoint: {
    mobileBreakpoint: 'sm' // 定义移动端断点，‘sm’ 代表宽度小于等于 960px 时视为移动端
  }
})
