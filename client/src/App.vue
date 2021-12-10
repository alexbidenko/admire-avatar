<script setup lang="ts">
import {
  NConfigProvider, NLayout, darkTheme,
} from 'naive-ui';
import Main from './pages/Main.vue';
</script>

<template>
  <n-config-provider :theme="darkTheme">
    <n-layout>
      <Main />
    </n-layout>
  </n-config-provider>
</template>

<style></style>
