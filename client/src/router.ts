import { createWebHistory, createRouter } from "vue-router";
import Home from "./pages/Home.vue";
import Login from "./pages/Login.vue";
import Signup from "./pages/Signup.vue";
import VerifyEmail from "./pages/VerifyEmail.vue";
import NotFound from "./pages/NotFound.vue";

const routes = [
  { path: "/:pathMatch(.*)*", name: "NotFound", component: NotFound },
  { path: "/", component: Home },
  { path: "/signup", component: Signup },
  { path: "/login", component: Login },
  { path: "/verification", component: VerifyEmail },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
