<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {
  NButton,
  NCard,
  NButtonGroup,
  NImage,
  NH3,
  NGrid,
  NGridItem,
} from 'naive-ui';
import {getImages, deleteImage, downloadImage, createAvatar} from '~/api/images';
import {ImageType} from '~/types/image';
import {CloseSharp} from '@vicons/material';

const images = ref<ImageType[]>([]);
const isSelectedImage = ref(false);

onMounted(() => {
  getImages().then(({data}) => {
    images.value = data;
  });
});

const deleteCurrentImage = (id: number) => {
  deleteImage(id);
};

const downloadSelectedImage = (id: number) => {
  downloadImage(id);
};

const selectedAvatar = (id: number) => {
  isSelectedImage.value = true;
  createAvatar(id);
};
</script>

<template>
  <div class="container">
    <n-card>
      <n-h3 v-if="images.length === 0">Добавьте первое изображение</n-h3>
      <n-grid cols="2 375:3 600:4 700:5 1200:6">
        <n-grid-item
          v-for="image in images"
          :key="image.id"
        >
          <n-image
              @click="selectedAvatar(image.id)"
              :class="{'selected': isSelectedImage}"
              preview-disabled
              src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
          />
          <CloseSharp @click="deleteCurrentImage(image.id)" class="close"/>
          <n-button type="success" @click="downloadSelectedImage(image.id)">Скачать</n-button>
        </n-grid-item>
        <n-grid-item>
          <n-image
              @click="selectedAvatar(image.id)"
              :class="{'selected': isSelectedImage}"
              preview-disabled
              src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
          />
          <CloseSharp @click="deleteCurrentImage(image.id)" class="close"/>
          <n-button type="success" @click="downloadSelectedImage(image.id)">Скачать</n-button>
        </n-grid-item>
      </n-grid>
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

.n-grid > div {
  position: relative;
  display: flex;
  flex-direction: column;

  .n-button {
    margin-top: 10px;
  }
}

.close {
  position: absolute;
  top: 0;
  right: 0;
  width: 20%;
  cursor: pointer;
}

.selected {
  border: 1px solid red;
  border-radius: 5px;
}

.n-image > img {
  width: 100%;
  cursor: pointer;
}
</style>
