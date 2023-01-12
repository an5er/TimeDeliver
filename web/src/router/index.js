import { createRouter, createWebHistory } from 'vue-router'
import {defineAsyncComponent} from 'vue'

const routes = [
{
    path: '/login',
    name: 'Login',
    component: defineAsyncComponent(() => import('../view/login/index.vue'))
},
{
    path: '/home',
    name: 'home',
    component: defineAsyncComponent(() => import('../view/home/index.vue'))
},
{
    path: '/test',
    name: 'test',
    component: defineAsyncComponent(() => import('../view/test/index.vue'))
}
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

// router.beforeEach((to, from, next) => {
//     if(to.path === "/"){
//         return next("/login")
//     }
//     if(to.path === "/home"){
//         if (window.sessionStorage.getItem("token") != null){
//             if (window.sessionStorage.getItem("token").length > 20) {
//                 return next()
//             }else {
//                 return next("/login")
//             }
//         }else {
//             return next("/login")
//         }
//     }
//     if(to.path === "/login"){
//         if (window.sessionStorage.getItem("token") != null){
//             if (window.sessionStorage.getItem("token").length > 20){
//                 return next("/home")
//             }else {
//                 next()
//             }
//         }else {
//             next()
//         }
//     }
//     next()
// })

export default router
