import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView2.vue'
import BlogView from '../views/Blog.vue'
import ContactView from '../views/Contact.vue'
import ResumeView from '../views/Resume.vue'
// import AboutView from '../views/AboutView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HomeView,
    },
    {
      path: '/blog',
      name: 'Blog',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: BlogView,
    },
    {
      path: '/contact',
      name: 'Contact',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: ContactView,
    },
    {
      path: '/resume',
      name: 'Resume',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: ResumeView,
    },
  ],
})

export default router
