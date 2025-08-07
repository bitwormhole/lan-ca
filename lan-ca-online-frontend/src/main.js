import './assets/main.css'

import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import App from './App.vue'
import router from './router'
import pinia from '@/stores/pinia'


import AdminFrame from '@/components/frames/AdminFrame.vue'
import UserFrame from '@/components/frames/UserFrame.vue'




const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.component("AdminFrame", AdminFrame)
app.component("UserFrame", UserFrame)

app.use(pinia)
app.use(router)
app.use(ElementPlus)

app.mount('#app')
