import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import config from './config'


async function initApplication() {
  await config.fetchConfig()

  const app = createApp(App)
  app.use(router)
  app.mount('#app')
}

initApplication()
