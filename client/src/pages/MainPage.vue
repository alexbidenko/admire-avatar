<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {
  NButton,
  NCard,
  NButtonGroup,
  NImage,
} from 'naive-ui';
import {getImages} from '~/api/images';
import {ImageType} from '~/types/image';

const images = ref<ImageType[]>([]);

onMounted(() => {
  getImages().then(({data}) => {
    images.value = data;
  });
});
</script>

<template>
  <div class="container">
    <n-card>
      <n-image
          v-for="image in images"
          :key="image.id"
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
