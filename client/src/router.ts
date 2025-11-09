import { createWebHistory, createRouter } from "vue-router";

import Home from "./views/Home.vue";
import Login from "./views/Login.vue";
import Signup from "./views/Signup.vue";
import VerifyEmail from "./views/VerifyEmail.vue";
import NotFound from "./views/NotFound.vue";

const routes = [
  { path: "/", component: Home },
  { path: "/signup", component: Signup },
  { path: "/verification", component: VerifyEmail },
  { path: "/login", component: Login },
  { path: "/:pathMatch(.*)*", name: "NotFound", component: NotFound },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
