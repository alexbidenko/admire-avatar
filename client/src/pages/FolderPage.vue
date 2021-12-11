<script setup lang="ts">
import {ref} from 'vue';
import {
  useLoadingBar,
} from 'naive-ui';
import {FolderType, ImageType} from '~/types/image';
import $axios from '~/api';
import {useRoute} from 'vue-router';
import ImageList from '../components/ImageList.vue';

const route = useRoute();
const loader = useLoadingBar();

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

const clearItem = (id: number) => {
  images.value = images.value.filter((el) => el.id !== id);
};
</script>

<template>
  <div class="container">
    <ImageList :images="images" @clear-item="clearItem" :folders="folders" />
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
