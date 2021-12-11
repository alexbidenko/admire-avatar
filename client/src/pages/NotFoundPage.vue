<script setup lang="ts">
import {NResult} from 'naive-ui';
</script>

<template>
  <div class="notFoundPage">
    <n-result
      status="404"
      title="404 Not Found"
    />
  </div>
</template>

<style lang="scss">
.notFoundPage {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
