import Vue from 'vue'
import Router from 'vue-router'
import App from '../App'

//安装路由
Vue.use(Router)

const home = r => require.ensure([], () => r(require('../page/home/home')), 'home')
const login = r => require.ensure([], () => r(require('../page/login/login')), 'login')
const collect = r => require.ensure([], () => r(require('../page/collect/collect')), 'collect')


//路由配置
export default new Router({
  mode:'history',
  routes:[
    {
      path:'/',
      name:"home",
      component:home,
    },
    {
      path:'/collect',
      name:'collect',
      component:collect
    },
    {
      path:'login',
      name:'login',
      component:login
    }
  ]
})