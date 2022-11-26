import Vue from 'vue'
import VueRouter from 'vue-router'
import { ifValidToken } from '../store/api_auth'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import('@/components/Home.vue'),
        meta: {
            requiresAuth: true
        }
    },
    {
        path: '/about',
        name: 'About',
        component: () => import('@/components/About.vue'),
        meta: {
            requiresAuth: true
        }
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/components/Login.vue'),
        meta: {
            requiresAuth: false
        }
    },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

router.beforeEach((to, from, next) => {
    if (to.path === '/login' && ifValidToken()) {
        next({
            name: 'Home',
            params: { redirect: to.fullPath }
        })
    }

    if (to.matched.some(record => !record.meta.requiresAuth)) {
        return next()
    }

    if (!ifValidToken()) {
        next({
            name: 'Login',
            query: { redirect: to.fullPath }
        })
    } else {
        next()
    }
})

export default router


