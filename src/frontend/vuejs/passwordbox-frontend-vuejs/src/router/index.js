import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },


    { path: '/login', name: 'sign-in', component: () => import('../views/SignInView.vue') },
    { path: '/signup', name: 'sign-up', component: () => import('../views/SignUpView.vue') },
    { path: '/guide', name: 'guide-to-add-user', component: () => import('../views/GuideToAddUser.vue') },
    { path: '/fast-gen', name: 'fast-gen', component: () => import('../views/FastGenView.vue') },

  ],
})

export default router
