import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import dayjs  from 'dayjs'

import './plugins/http'
const app = createApp(App)
app.config.globalProperties.$dateFormat = (indate, outdate) => {
  return dayjs(indate).format(outdate)
}

app.use(router)

app.use(vuetify)

app.mount('#app')