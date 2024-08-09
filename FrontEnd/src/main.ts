import { createApp } from 'vue'
import App from './views/App.vue'
import router from './router'
import ArcoVue from '@arco-design/web-vue'
import './styles/tailwind.css'
import '@arco-design/web-vue/dist/arco.css'

const app = createApp(App)

app.use(router)
app.use(ArcoVue)
app.mount('#app')