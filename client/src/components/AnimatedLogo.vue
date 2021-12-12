<script setup lang="ts">
import {ref} from 'vue';
import {NCard, NH1} from 'naive-ui';
const show = ref(true);
setTimeout(() => {
  show.value = false;
}, 1000);
</script>

<template>
  <transition name="fade">
    <n-card v-if="show" tile class="logoContainer">
      <img class="logo" src="/loader.png" alt="Логотип проекта" />
      <n-h1>PICART</n-h1>
    </n-card>
  </transition>
</template>

<style lang="scss">
.logoContainer {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1000000;
  & > div {
    display: flex;
    align-content: center;
    justify-content: center;
    flex-direction: column;
  }
  h1 {
    text-align: center;
    font-size: 52px;
    letter-spacing: 0.2em;
  }
}
.logo {
  animation: animated-logo 1.5s ease forwards;
  width: 240px;
  height: 240px;
  margin: 0 auto 32px;
  @media (min-width: 1024px) {
    width: 320px;
    height: 320px;
  }
  @keyframes animated-logo {
    0% {
      opacity: 0;
      transform: scale(0.7) rotate(45deg);
    }
    30% {
      opacity: 1;
      transform: rotate(45deg);
    }
    70% {
      transform: rotate(-15deg);
    }
    80% {
      transform: rotate(0);
    }
    90% {
      transform: scale(1.3) rotate(0);
      opacity: 1;
    }
    100% {
      transform: scale(1.1) rotate(0);
      opacity: 0;
    }
  }
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.4s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
