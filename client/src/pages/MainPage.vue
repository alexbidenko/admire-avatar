<script setup lang="ts">
import {onMounted, onUnmounted, ref} from 'vue';
import {
  NButton,
  NCard,
  NButtonGroup,
  NImage,
  NH3,
  NGrid,
  NGridItem, NIcon, useLoadingBar, NH1,
} from 'naive-ui';
import {getImages, deleteImage, createAvatar} from '~/api/images';
import {ImageType} from '~/types/image';
import {Download as DownloadRegular, SaveRegular, TrashAlt} from '@vicons/fa';
import {useMainStore} from '~/store';

const COUNT = 40;

const loader = useLoadingBar();
const store = useMainStore();
const page = ref(0);
const isRequest = ref(true);
const isFinish = ref(false);

const images = ref<ImageType[]>([]);

loader.start();
getImages(0, COUNT).then(({data}) => {
  isRequest.value = false;
  images.value = data;
  if (data.length < COUNT) isFinish.value = true;
}).catch(loader.error).finally(loader.finish);

onMounted(() => {
  const scrollContainer = document.querySelector('.scrollContainer > .n-scrollbar-container')!;
  const onScroll = () => {
    const bottomOfWindow = scrollContainer.scrollTop + scrollContainer.clientHeight > scrollContainer.scrollHeight - 300;
    if (bottomOfWindow && !isRequest.value && !isFinish.value) {
      isRequest.value = true;
      page.value++;
      getImages(page.value * COUNT, COUNT).then(({data}) => {
        isRequest.value = false;
        images.value = images.value.concat(data);
        if (data.length < COUNT) isFinish.value = true;
      });
    }
  };

  scrollContainer.addEventListener('scroll', onScroll);
  onUnmounted(() => {
    scrollContainer.removeEventListener('scroll', onScroll);
  });
});

const deleteCurrentImage = (id: number) => {
  deleteImage(id).then(() => {
    images.value = images.value.filter((el) => el.id !== id);
  });
};

const selectAvatar = (image: ImageType) => {
  createAvatar(image.id).then(() => {
    images.value.forEach((el) => {
      el.main = false;
    });
    image.main = true;
    store.commit('setAvatar', image);
  });
};
</script>

<template>
  <div class="container">
    <n-h1>Генерация аватарок</n-h1>
    <n-card style="margin-bottom: 32px">
      <router-link to="/generate">
        <n-button type="warning">Сгенерировать</n-button>
      </router-link>
    </n-card>
    <n-card v-if="images.length === 0">Добавьте первое изображение</n-card>
    <n-grid v-else cols="2 s:3 m:4 l:5 xl:6 2xl:8" responsive="screen" x-gap="16" y-gap="16">
      <n-grid-item
          v-for="image in images"
          :key="image.id"
      >
        <n-card content-style="padding: 0">
          <n-image
              :class="{'selected': image.main}"
              :src="`/api/files/images/${image.source}`"
          />
          <n-button-group style="padding-bottom: 10px; width: 100%; justify-content: center">
            <a download="avatar.png" :href="`/api/files/images/${image.source}`">
              <n-button type="info">
                <n-icon>
                  <download-regular />
                </n-icon>
              </n-button>
            </a>
            <n-button type="success" @click="selectAvatar(image)">
              <n-icon>
                <save-regular />
              </n-icon>
            </n-button>
            <n-button type="error" @click="deleteCurrentImage(image.id)">
              <n-icon>
                <trash-alt />
              </n-icon>
            </n-button>
          </n-button-group>
        </n-card>
      </n-grid-item>
    </n-grid>
  </div>
</template>

<style lang="scss">
.n-grid > div {
  position: relative;
  display: flex;
  flex-direction: column;
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
