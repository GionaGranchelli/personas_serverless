import './assets/main.css'

import { createApp, provide } from 'vue'
import App from './App.vue'
import router from './router'
import config from './config'


async function initApplication() {
  const app = createApp(App)
  const cnf = await config.fetchConfig()
  app.use(router)
  app.provide('apiEndpoint', cnf?.apiEndpoint || "http://localhost:3000")
  app.mount('#app')
}

initApplication()
