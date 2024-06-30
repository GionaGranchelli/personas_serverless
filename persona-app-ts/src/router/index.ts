import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import AboutView from '@/views/AboutView.vue'
import CreatePersona from '@/views/CreatePersona.vue'
import ListPersonas from '@/views/ListPersonas.vue'
import SearchPersona from '@/views/SearchPersona.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView
    },
    {
      path: '/create-persona',
      name: 'Create-Persona',
      component: CreatePersona
    },
    {
      path: '/list-personas',
      name: 'List-Personas',
      component: ListPersonas
    },
    {
      path: '/search-persona',
      name: 'Search-Personas',
      component: SearchPersona
    }

  ]
})

export default router
