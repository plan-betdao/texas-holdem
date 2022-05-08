<template>
  <div class="card-style-container">
    <div class="svg">
      <svg
        viewBox="0 0 100 100"
        :width="`${sizeValue}px`"
        :height="`${sizeValue}px`"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path :d="path" :fill="color" />
      </svg>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
  type: any;
  size: string;
}>();

const pathObj = {
  h: "M 0,25 A 25,25 0,0,1 50,25 A 25,25 0,0,1 100,25 C 100,60 70,60 50,100 C 30,60 0,60 0,25",
  c: "M25,25 A25,25 0,0,1 75,25 A25,25 0,1,1 52,65 A35,35 0,0,0 75,100 L25,100 A35,35 0,0,0 48,65 A25,25 0,1,1 25,25",
  s: "M 4,65 A 22,22 0,0,0 48,65 Q 50,90 30,100 L 70,100 Q 50,90 52,65 A 22,22 0,0,0 96,65 C 96,40 70,40 50,0 C 30,40 4,40 4,65",
  d: "M 50,0 Q50,15 15,50 Q50,85 50,100 Q50,85 85,50 Q50,15 50,0",
};
const sizeObj = {
  big: 22,
  small: 12,
  default: 20,
};
const sizeValue = computed(() => {
  sizeObj[props.size];
});

const path = computed(() => {
  return this.pathObj[this.type];
});

const color = computed(() => {
  return this.type === "h" || this.type === "d" ? "red" : "black";
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
.card-style-container {
  .svg {
    position: relative;
  }
}
</style>
