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
    },
    {
        path: '/sign-up',
        component: () => import('../views/SignUp.vue'),
    },
    {
        path: '/user/orders',
        component: () => import('../views/Order.vue'),
    },
    {
        path: '/admin/home',
        alias: '/admin',
        component: () => import('../views/admin/Home.vue'),
    },
    {
        path: '/admin/addProduct',
        component: () => import('../views/admin/addProduct.vue'),
    },
    {
        path: '/product-detail/:id',
        name: 'ProductDetail',
        component: () => import('../views/ProductDetail.vue'),
        props: true
    },
    {
        path: '/cart',
        name: 'Cart',
        component: () => import('../views/CartPage.vue'),
        meta: { requiresAuth: true }
    },
    {
        path: '/checkout',
        name: 'Checkout',
        component: () => import('../views/CheckoutView.vue'),
        meta: { requiresAuth: true } // 需要登录验证
    },
    {
        path: '/payment/:orderId',
        name: 'Payment',
        component: () => import('../views/Payment.vue'),
        meta: { requiresAuth: true }
    },
    {
        path: '/orders',
        name: 'Orders',
        component: () => import('../views/Order.vue'),
        meta: { requiresAuth: true }
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router