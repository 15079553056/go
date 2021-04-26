import Vue from 'vue'
import Router from 'vue-router'



Vue.use(Router)

const login = r => require.ensure([], () => r(require('../page/login/login')), 'login')
const collect = r => require.ensure([], () => r(require('../page/collect/collect')), 'collect')

const routes= [
  {
    path: '/',
    component: login
  },
  {
    path: '/collect',
    component: collect
  }
]

export default new Router({
  routes,
  strict: process.env.NODE_ENV !== 'production',
})