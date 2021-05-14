import Vue from 'vue'
import Router from 'vue-router'
import VueRouter from 'vue-router'

Vue.use(Router)

const home = r => require.ensure([], () => r(require('../page/home/home')), 'home')
const manager_login = r => require.ensure([], () => r(require('../page/manager/login')), 'home')

const router = new VueRouter({
    mode: 'history',
    routes: [
        {
            path: '/',
            redirect: '/home',
        },
        {
            path: '/home',
            component: home
        },
        {
            path: '/login',
            component: manager_login
        }
    ]
})

export default router