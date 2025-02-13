import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/',
        alias: '/home',
        component: () => import('../views/Home.vue'),
    },
    {
        path: '/login',
        component: () => import('../views/LoginIn.vue'),
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router