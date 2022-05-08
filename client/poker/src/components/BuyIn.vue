<template>
  <div class="buy-in" v-show="showBuyIn">
    <div class="shadow" @click="closeBuyIn"></div>
    <div class="buy-in-body">
      <div class="input-bd">
        <div class="input-name">
          <span>buy in: </span><input type="number" v-model="buyInSize" />
        </div>
        <range
          :max="max"
          :min="min"
          v-model="buyInSize"
          @change="getBuyInSize"
        ></range>
      </div>
      <div class="btn"><span @click="buyIn">buy in</span></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import range from "@/components/BuyRange.vue";
import { watch } from "vue";
const props = defineProps({
  showBuyIn: { type: Boolean, required: true },
  min: { type: Number, required: true },
  max: { type: Number, required: true },
});
const emit = defineEmits(["update:showBuyIn", "buyIn"]);

const buyInSize = 0;

watch(
  () => buyInSize,
  (val) => {
    this.buyInSize =
      val > this.max ? this.max : val < this.min ? this.min : val;
  }
);

function getBuyInSize(val: string) {
  this.buyInSize = Number(val);
}

function closeBuyIn() {
  emit("update:showBuyIn", false);
}

async function buyIn() {
  this.closeBuyIn();
  this.emit("buyIn", this.buyInSize);
}
function mounted() {
  this.buyInSize = this.min;
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.buy-in {
  position: fixed;
  z-index: 99;

  .shadow {
    position: fixed;
    z-index: 9;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.2);
  }

  .buy-in-body {
    z-index: 99;
    position: fixed;
    left: 50%;
    top: 50%;
    margin: -100px -150px;
    width: 300px;
    border-radius: 12px;
    box-sizing: border-box;
    background: #fff;
    padding: 20px;
  }

  .input-text {
    input {
      width: 100px;
    }
  }
  .input-name {
    margin-bottom: 15px;
    font-size: 20px;
    text-align: center;
    input {
      width: 70px;
      font-size: 20px;
    }
  }
  .btn {
    margin-top: 20px;
  }
}
</style>
