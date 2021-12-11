<script setup lang="ts">
import {ref} from 'vue';
import {
  NButton,
  NCard,
  NButtonGroup,
  NImage,
  NH3,
  NGrid,
  NGridItem,
} from 'naive-ui';
import {getImages, deleteImage, createAvatar} from '~/api/images';
import {ImageType} from '~/types/image';
import {TrashAlt} from '@vicons/fa';

const images = ref<ImageType[]>([]);

getImages().then(({data}) => {
  images.value = data;
});

const deleteCurrentImage = (id: number) => {
  deleteImage(id);
};

const selectedAvatar = (image: ImageType) => {
  createAvatar(image.id).then(() => {
    image.main = true;
  });
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
              @click="selectedAvatar(image)"
              :class="{'selected': image.main}"
              preview-disabled
              :src="`/api/files/images/${image.source}`"
          />
          <trash-alt @click="deleteCurrentImage(image.id)" class="close"/>
          <a download="avatar.png" :href="`/api/files/images/${image.source}`">
            <n-button type="success">Скачать</n-button>
          </a>
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
  border: 4px solid red;
  border-radius: 5px;
}

.n-image > img {
  width: 100%;
  cursor: pointer;
}
</style>
