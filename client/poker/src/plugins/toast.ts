import { defineComponent } from "vue";
import Toast from "@/components/ToastPlugin.vue";

import { App } from "vue";

interface Options {
  text?: string;
  timeOut?: number;
}

const ToastConstructor = defineComponent({ extends: Toast });

export class ToastExtendConstructor extends ToastConstructor {
  public close() {
    this.$props.show = false;
    this.$off("update:show");
  }
}

let instance: ToastExtendConstructor;
let defaultOptions: Options;

const getInstance = () => {
  if (instance) {
    return instance;
  }

  instance = new ToastExtendConstructor({
    el: document.createElement("div"),
  });

  return instance;
};

const toast = (options: string | Options) => {
  const vm = getInstance();

  if (!defaultOptions) {
    defaultOptions = { ...vm.$props };
  }

  let opts: Options;
  if (typeof options === "string") {
    opts = { ...defaultOptions, text: options };
  } else {
    opts = { ...defaultOptions, ...options };
  }

  Object.keys(opts).forEach((key) => {
    vm.$props[key] = opts[key as keyof Options];
  });
  console.log(vm.$props);
  vm.$props.show = true;
  vm.$off("update:show");
  vm.$on("update:show", (val: boolean) => {
    vm.$props.show = val;
  });
  document.body.appendChild(vm.$el);

  return vm;
};

export default {
  install: (app: App) => {
    app.config.globalProperties.$plugin.toast = toast;
  },
};
