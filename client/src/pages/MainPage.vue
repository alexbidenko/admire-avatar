<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {
  NButton,
  NCard,
  NButtonGroup,
  NImage,
  NH3,
} from 'naive-ui';
import {getImages, deleteImage} from '~/api/images';
import {ImageType} from '~/types/image';

const images = ref<ImageType[]>([]);

onMounted(() => {
  getImages().then(({data}) => {
    images.value = data;
  });
});

const deleteCurrentImage = (id: number) => {
  deleteImage(id);
};
</script>

<template>
  <div class="container">
    <n-card>
      <n-h3 v-if="images.length === 0">Добавьте первое изображение</n-h3>
      <n-image
          v-for="image in images"
          @click="deleteCurrentImage(image.id)"
          :key="image.id"
          width="100"
          src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
      />
    </n-card>
    <n-button-group class="button">
      <router-link to="/generate">
        <n-button type="warning">Добавить новое изображение</n-button>
      </router-link>
    </n-button-group>
  </div>
</template>

<style lang="scss">
.button {
  padding-top: 24px;
}

.imageContainer {
  position: relative;
}

.close__icon {
  position: absolute;
  right: 0;
  top: 0;
}
</style>
