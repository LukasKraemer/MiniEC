import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import ItemPage from "../views/ItemPage.vue";
import Admin from "../views/admin.vue";

const routes: Array<RouteRecordRaw> = [
  { path: "/", name: "Home", component: Home },
  { path: "/login", name: "Login", component: Login },
  { path: "/admin", name: "admin", component: Admin },
  { path: "/article/:id", name: "article", component: ItemPage },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
