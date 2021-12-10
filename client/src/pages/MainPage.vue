<script setup lang="ts">
import {
  NButton,
  NCard,
  NButtonGroup,
  NImage,
} from 'naive-ui';
import {useRouter} from 'vue-router';

const router = useRouter();
</script>

<template>
  <div class="container">
    <n-card>
      <n-image
          width="100"
          src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
      />
    </n-card>
    <n-button-group class="button">
      <router-link to="/generateImages">
        <n-button type="warning">Добавить новое изображение</n-button>
      </router-link>
    </n-button-group>
  </div>
</template>

<style lang="scss">
.button {
  padding-top: 24px;
}
</style>
