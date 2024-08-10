import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ArcoVue from '@arco-design/web-vue'
import './tailwind.css'
import '@arco-design/web-vue/dist/arco.css'

document.body.setAttribute('arco-theme', 'dark')

const app = createApp(App)

app.use(router)
app.use(ArcoVue)
app.mount('#app')
