import Vue from 'vue'
import VueRouter from 'vue-router'
import HedgeHogs from '../components/pages/HedgeHogs'
import Auth from '../components/pages/Auth'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'HedgeHogs',
    component: HedgeHogs
  },
  {
    path: '/auth',
    name: 'Auth',
    component: Auth
  }
]

const router = new VueRouter({
  routes
})

export default router
