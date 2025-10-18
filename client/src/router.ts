import { createWebHistory, createRouter } from "vue-router";

import Home from "./views/Home.vue";
import Login from "./views/Login.vue";
import Signup from "./views/Signup.vue";
import VerifyEmail from "./views/VerifyEmail.vue";

const routes = [
    { path: "/", component: Home },
    { path: "/signup", component: Signup },
    { path: "/login", component: Login },
    { path: "/signup/verify-email", component: VerifyEmail }
]

export const router = createRouter({
    history: createWebHistory(),
    routes
})