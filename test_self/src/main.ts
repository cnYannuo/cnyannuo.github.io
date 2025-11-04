import './assets/main.css'

import { createApp, inject, ref } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

export const activeSection = ref('home');
export function setActiveSection(section: string) {
    activeSection.value = section;
    window.scrollTo(0, 0);
}


const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')




