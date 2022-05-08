import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import toastPlugin from "@/plugins/toast";

createApp(App).use(toastPlugin).use(store).use(router).mount("#app");
