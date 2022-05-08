import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Game from "@/views/GameTable.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/game/:roomNumber/:isOwner?",
    name: "game",
    component: Game,
    meta: {
      title: "game",
    },
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
