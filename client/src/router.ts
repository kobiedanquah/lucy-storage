import { lazy } from "solid-js";

const routes = [
  { path: "/", component: lazy(() => import("./pages/Home")) },
  { path: "/signup", component: lazy(() => import("./pages/Signup")) },
];

export default routes;
