import './assets/main.css'

import { createApp } from 'vue'
// import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'

import AdminFrame from '@/components/frames/AdminFrame.vue'
import UserFrame from '@/components/frames/UserFrame.vue'

import pinia from '@/stores/pinia'



// const pinia = createPinia()
const app = createApp(App)

app.component("AdminFrame", AdminFrame)
app.component("UserFrame", UserFrame)

app.use(pinia)
app.use(router)
app.use(ElementPlus)

app.mount('#app')
