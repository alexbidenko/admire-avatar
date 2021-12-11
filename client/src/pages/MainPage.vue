<script setup lang="ts">
import {onMounted, onUnmounted, ref} from 'vue';
import {
  NButton,
  NCard,
  NButtonGroup,
  NGrid,
  NGridItem,
  NIcon,
  useLoadingBar,
  NH1,
  useMessage,
  NModal,
  NInput,
  NSpace,
  useDialog,
} from 'naive-ui';
import {getImages} from '~/api/images';
import {FolderType, ImageType} from '~/types/image';
import {
  TrashAlt,
  FolderPlus,
} from '@vicons/fa';
import $axios from '~/api';
import {UserType} from '~/types/user';
import ImageList from '~/components/ImageList.vue';

const COUNT = 40;

const dialog = useDialog();
const message = useMessage();
const loader = useLoadingBar();
const page = ref(0);
const isRequest = ref(true);
const isFinish = ref(false);

const showModal = ref(false);
const showShare = ref(false);
const folders = ref<FolderType[]>([]);
const images = ref<ImageType[]>([]);
const selectedUser = ref<UserType>();
const directoryName = ref('');

loader.start();
Promise.all([
  $axios.get('folders').then(({data}) => {
    folders.value = data;
  }),
  getImages(0, COUNT).then(({data}) => {
    isRequest.value = false;
    images.value = data;
    if (data.length < COUNT) isFinish.value = true;
  }),
]).catch(loader.error).finally(loader.finish);

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

const cancel = () => {
  showShare.value = false;
  showModal.value = false;
  directoryName.value = '';
  selectedUser.value = undefined;
};

const deleteFolder = (folder: FolderType) => {
  dialog.warning({
    title: 'Подтверждение удаления',
    content: `Вы уверены что хотите удалить директорию ${folder.name}?`,
    negativeText: 'Отмена',
    positiveText: 'Удалить',
    onPositiveClick: () => {
      loader.start();
      $axios.delete(`folders/${folder.id}`).then(() => {
        folders.value = folders.value.filter((el) => el.id !== folder.id);
      }).catch(loader.error).finally(loader.finish);
    },
  });
};

const createFolder = () => {
  loader.start();
  $axios.post('folders', {
    name: directoryName.value,
  }).then(({data}) => {
    folders.value = folders.value.concat([data]).sort((a, b) => a.name > b.name ? 1 : -1);
    showModal.value = false;
    directoryName.value = '';
  }).catch(() => {
    message.info('Во время создании директории произошла ошибка');
    loader.error();
  }).finally(loader.finish);
};

const clearItem = (id: number) => {
  images.value = images.value.filter((el) => el.id !== id);
};
</script>

<template>
  <div class="container">
    <n-h1>Мои изображения</n-h1>
    <n-card>
      <n-space justify="space-between">
        <router-link to="/generate">
          <n-button type="warning">Сгенерировать</n-button>
        </router-link>
        <n-button type="success" @click="showModal = !showModal">
          <n-icon>
            <folder-plus />
          </n-icon>
        </n-button>
      </n-space>
    </n-card>
    <n-card style="margin-top: 32px" v-if="!images.length && !folders.length">Добавьте первое изображение</n-card>
    <n-grid style="margin-top: 32px" v-if="folders.length" cols="2 s:3 m:4 l:5 xl:6 2xl:8" responsive="screen" x-gap="16" y-gap="16">
      <n-grid-item v-for="folder in folders" :key="folder.id">
        <router-link :to="`/folder/${folder.id}`">
          <n-card hovered class="mainPage__folder">
            <n-space justify="space-between" align="center">
              {{ folder.name }}
              <n-button size="small" type="error" @click.stop.prevent="deleteFolder(folder)" secondary>
                <n-icon>
                  <trash-alt />
                </n-icon>
              </n-button>
            </n-space>
          </n-card>
        </router-link>
      </n-grid-item>
    </n-grid>

    <ImageList :images="images" @clear-item="clearItem" :folders="folders" />

    <n-modal v-model:show="showModal">
      <n-card style="width: 600px;" title="Создание директории" preset="card">
        <n-input v-model:value="directoryName" type="text" placeholder="Название директории" />
        <template #footer>
          <n-button-group style="display: flex; justify-content: end">
            <n-button type="warning" @click="cancel">
              Отменить
            </n-button>
            <n-button type="success" @click="createFolder">
              Создать
            </n-button>
          </n-button-group>
        </template>
      </n-card>
    </n-modal>
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

.mainPage {
  &__folder {
    .n-button {
      opacity: 0;
      transition: opacity 0.3s ease, color .3s var(--bezier), background-color .3s var(--bezier), opacity .3s var(--bezier), border-color .3s var(--bezier);
    }

    .n-card__content {
      padding-top: 8px !important;
      padding-bottom: 8px !important;
    }

    &:hover .n-button {
      opacity: 1;
    }
  }

  &__actions {
    padding-bottom: 10px;
    width: 100%;
    justify-content: center;

    .n-button {
      width: 34px;
    }
  }
}
</style>
