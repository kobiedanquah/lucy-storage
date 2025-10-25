import { lazy } from "solid-js";

const routes = [
  { path: "/", component: lazy(() => import("./pages/Home")) },
  { path: "/signup", component: lazy(() => import("./pages/Signup")) },
  {
    path: "/verification",
    component: lazy(() => import("./pages/VerifyEmail")),
  },
  {path: "*404", component: lazy(()=>import("./pages/NotFound"))}
];

export default routes;
