import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import './assets/css/global.css'
import axios from 'axios'
import { createPinia } from 'pinia'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import Particles from 'particles.vue3'
import 'animate.css/animate.min.css'
import VueTilt from 'vue-tilt.js'

import router from './router/index'

const app = createApp(App)
app.config.globalProperties.$http = axios

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

// app.config.productionTip = false

const pinia = createPinia()

app
    .use(router)
    .use(ElementPlus)
    .use(Particles)
    .use(pinia)
    .use(VueTilt)
    .mount('#app')

export default app
