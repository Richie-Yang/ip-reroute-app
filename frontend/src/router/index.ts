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
      path: '/en',
      name: 'en',
      component: () => import('../views/en/EnHomeView.vue'),
    },
    {
      path: '/tw',
      name: 'tw',
      component: () => import('../views/tw/TwHomeView.vue'),
    },
    {
      path: '/en/product',
      name: 'en-product',
      component: () => import('../views/en/EnProductView.vue'),
    },
    {
      path: '/tw/product',
      name: 'tw-product',
      component: () => import('../views/tw/TwProductView.vue'),
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/'
    }
  ],
})

export default router
