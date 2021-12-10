<script setup lang="ts">
import {
  NConfigProvider, NLayout, darkTheme,
} from 'naive-ui';
</script>

<template>
  <n-config-provider :theme="darkTheme">
    <n-layout>
      <router-view v-slot="{ Component }">
        <transition name="scale" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </n-layout>
  </n-config-provider>
</template>

<style lang="scss">
body {
  overflow: hidden;
}

a {
  text-decoration: none;
}

.n-layout {
  height: 100vh;
}

.scale-enter-active,
.scale-leave-active {
  transition: opacity 0.2s ease;
}

.scale-enter-from,
.scale-leave-to {
  opacity: 0;
}

.n-button-group:not(.n-button-group--vertical) a:last-child:not(:first-child) .n-button {
  margin-left: 0!important;
  border-top-left-radius: 0!important;
  border-bottom-left-radius: 0!important;
}

.n-button-group:not(.n-button-group--vertical) a:not(:first-child):not(:last-child) .n-button {
  margin-left: 0!important;
  margin-right: 0!important;
  border-radius: 0!important;
}

.n-button-group:not(.n-button-group--vertical) a:first-child:not(:last-child) .n-button {
  margin-right: 0!important;
  border-top-right-radius: 0!important;
  border-bottom-right-radius: 0!important;
}
</style>
