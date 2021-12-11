<script setup lang="ts">
import {computed, ref} from 'vue';
import {
  NButton,
  NCard,
  NButtonGroup,
  NImage,
  NGrid,
  NGridItem, NIcon, useLoadingBar, NDropdown,
} from 'naive-ui';
import {deleteImage, createAvatar} from '~/api/images';
import {FolderType, ImageType} from '~/types/image';
import {Download as DownloadRegular, UserCircle, TrashAlt, Share as ShareRegular} from '@vicons/fa';
import {useMainStore} from '~/store';
import $axios from '~/api';
import {useRoute} from 'vue-router';

const route = useRoute();
const loader = useLoadingBar();
const store = useMainStore();

const folders = ref<FolderType[]>([]);
const images = ref<ImageType[]>([]);

loader.start();
Promise.all([
  $axios.get(`images/folder/${route.params.folderId}`).then(({data}) => {
    images.value = data;
  }),
  $axios.get<FolderType[]>('folders').then(({data}) => {
    folders.value = data.sort((a, b) => a.name > b.name ? 1 : -1);
  }),
]).catch(loader.error).finally(loader.finish);

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

const options = computed(() => [
  {
    label: 'Переместить в директорию',
    key: 'title',
    disabled: true,
  },
  {
    label: 'Корневая директория',
    key: 0,
  },
  ...folders.value.filter((el) => el.id !== +route.params.folderId).map((el) => ({
    label: el.name,
    key: el.id,
  })),
]);

const handleSelect = (key: number, image: ImageType) => {
  loader.start();
  $axios.put(`images/${image.id}/folder/${key}`).then(() => {
    images.value = images.value.filter((el) => el.id !== image.id);
  }).catch(loader.error).finally(loader.finish);
};
</script>

<template>
  <div class="container">
    <n-grid cols="2 s:3 m:4 l:5 xl:6 2xl:8" responsive="screen" x-gap="16" y-gap="16">
      <n-grid-item
        v-for="image in images"
        :key="image.id"
      >
        <n-card content-style="padding: 0">
          <div class="squareImageCard">
            <n-image
                :class="{'selected': image.main}"
                :src="`/api/files/images/${image.source}`"
            />
          </div>
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
                <user-circle />
              </n-icon>
            </n-button>
            <n-button type="error" @click="deleteCurrentImage(image.id)">
              <n-icon>
                <trash-alt />
              </n-icon>
            </n-button>
            <n-dropdown trigger="hover" @select="handleSelect($event, image)" :options="options">
              <n-button>
                <n-icon>
                  <share-regular />
                </n-icon>
              </n-button>
            </n-dropdown>
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

.mainPage__folder {
  .n-button {
    opacity: 0;
    transition: opacity 0.3s ease, color .3s var(--bezier), background-color .3s var(--bezier), opacity .3s var(--bezier), border-color .3s var(--bezier);
  }

  &:hover .n-button {
    opacity: 1;
  }
}
</style>
