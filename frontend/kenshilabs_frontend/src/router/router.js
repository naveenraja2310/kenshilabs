import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../components/home.vue'
import Login from '../components/login.vue'
import register from '../components/register.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/login',
            name: 'login',
            component: Login
        },
        {
            path: '/register',
            name: 'register',
            component: register
        }
    ]
})

export default router
