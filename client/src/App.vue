<script setup lang="ts">
import {
  NConfigProvider, NLayout, darkTheme,
} from 'naive-ui';
</script>

<template>
  <n-config-provider :theme="darkTheme">
    <n-layout>Hello</n-layout>
  </n-config-provider>
</template>

<style></style>
