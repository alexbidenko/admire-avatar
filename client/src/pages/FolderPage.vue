<script setup lang="ts">
import {computed, ref} from 'vue';
import {
  NBreadcrumb, NBreadcrumbItem,
  NCard, NIcon,
  useLoadingBar,
} from 'naive-ui';
import {FolderType, ImageType} from '~/types/image';
import $axios from '~/api';
import {useRoute, useRouter} from 'vue-router';
import ImageList from '../components/ImageList.vue';
import {FolderRegular} from '@vicons/fa';
import {useMainStore} from '~/store';

const store = useMainStore();
const route = useRoute();
const router = useRouter();
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

const currentFolder = computed(() => folders.value.find((el) => el.id === +route.params.folderId));
</script>

<template>
  <div class="container">
    <n-card>
      <n-breadcrumb>
        <n-breadcrumb-item href="/" @click.prevent="router.push('/')">
          <n-icon>
            <folder-regular />
          </n-icon>
          Моя коллекция
        </n-breadcrumb-item>
        <n-breadcrumb-item>
          <n-icon>
            <folder-regular />
          </n-icon>
          {{ currentFolder?.name }}
        </n-breadcrumb-item>
      </n-breadcrumb>
    </n-card>
    <ImageList :images="images" @clear-item="clearItem" :folders="folders" :key="folders.length" :has-access="currentFolder?.userId === store.state.user.id" />
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

.n-image > img {
  width: 100%;
  cursor: pointer;
}

.mainPage__folder {
  .n-card__content {
    padding: 8px 8px 8px 16px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  p {
    margin: 0;
  }

  .n-button {
    opacity: 0;
    transition: opacity 0.3s ease, color .3s var(--bezier), background-color .3s var(--bezier), opacity .3s var(--bezier), border-color .3s var(--bezier);
  }

  &:hover .n-button {
    opacity: 1;
  }
}
</style>
