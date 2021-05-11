import Vue from 'vue'
import VueRouter from 'vue-router'
import Router from 'vue-router'
import App from '../App'

//1. 安装路由插件
Vue.use(Router)

const home = r => require.ensure([], () => r(require('../page/home/home')), 'home')
const login = r => require.ensure([], () => r(require('../page/login/login')), 'login')
const collect = r => require.ensure([], () => r(require('../page/collect/collect')), 'collect')

//2. 创建 VueRouter 对象
const router = new VueRouter({
  routes: [
    {
      path: '/',
      component: home,
    },
    {
      path: '/collect',
      component: collect
    },
    {
      path: '/login',
      component: login
    }
  ]
})

//3. 将 router 对象传入到 Vue 实例中
export default router