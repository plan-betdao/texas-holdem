<template>
  <div class="toast-container">
    <div class="toast-body" v-show="showValue">
      {{ text }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = {
  text: { type: String, required: true },
  show: { type: Boolean, default: false, required: true },
  timeOut: { type: Number, default: 3000, required: true },
};

let Time: number;

const emit = defineEmits(["update:show", "close"]);
let showValue = computed(() => ({
  get() {
    console.log("come in1111", props.show);
    if (props.show) {
      close();
    }
    return props.show;
  },
  set(val: boolean) {
    emit("update:show", val);
  },
}));

function close() {
  console.log("come in");
  clearTimeout(Time);
  Time = setTimeout(() => {
    showValue = false;
    emit("close");
  }, props.timeOut || 0);
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.toast-container {
  .toast-body {
    padding: 4px 10px;
    background: rgba(0, 0, 0, 0.6);
    text-align: center;
    color: #fff;
    font-size: 12px;
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate3d(-50%, -50%, 0);
    border-radius: 4px;
    line-height: 16px;
  }
}
</style>
