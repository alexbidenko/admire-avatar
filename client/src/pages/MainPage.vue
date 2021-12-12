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
  NH2,
  useMessage,
  NModal,
  NInput,
  NSpace,
  useDialog, NSwitch,
} from 'naive-ui';
import {getImages} from '~/api/images';
import {FolderType, ImageType} from '~/types/image';
import {
  TrashAlt,
  FolderPlus, EditRegular, EyeRegular,
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
const isRequestEdit = ref(false);

const showModal = ref(false);
const showShare = ref(false);
const editedFolder = ref<FolderType>();
const folders = ref<FolderType[]>([]);
const pubicFolders = ref<FolderType[]>([]);
const images = ref<ImageType[]>([]);
const selectedUser = ref<UserType>();
const directoryName = ref('');
const directoryPublic = ref(false);

loader.start();
Promise.all([
  $axios.get('folders/public').then(({data}) => {
    pubicFolders.value = data;
  }),
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
        if (folder.public) pubicFolders.value = pubicFolders.value.filter((el) => el.id !== folder.id);
      }).catch(loader.error).finally(loader.finish);
    },
  });
};

const editFolder = (folder: FolderType) => {
  showModal.value = true;
  editedFolder.value = folder;
  directoryName.value = folder.name;
  directoryPublic.value = folder.public;
};

const toCreateDirectory = () => {
  showModal.value = true;
  editedFolder.value = undefined;
};

const saveFolder = () => {
  if (isRequestEdit.value) return;
  isRequestEdit.value = true;
  loader.start();
  if (editedFolder.value) {
    $axios.put(`folders/${editedFolder.value.id}`, {
      name: directoryName.value,
      public: directoryPublic.value,
    }).then(({data}) => {
      folders.value = folders.value.map((el) => el.id === data.id ? data : el).sort((a, b) => a.name > b.name ? 1 : -1);
      if (editedFolder.value?.public && !data.public) pubicFolders.value = pubicFolders.value.filter((el) => el.id !== data.id);
      else if (!editedFolder.value?.public && data.public) pubicFolders.value = pubicFolders.value.concat([data]).sort((a, b) => a.name > b.name ? 1 : -1);
      else pubicFolders.value = pubicFolders.value.map((el) => el.id === data.id ? data : el).sort((a, b) => a.name > b.name ? 1 : -1);
      showModal.value = false;
      directoryName.value = '';
    }).catch(() => {
      message.info('Во время изменения директории произошла ошибка');
      loader.error();
    }).finally(() => {
      loader.finish();
      isRequestEdit.value = false;
    });
  } else {
    $axios.post('folders', {
      name: directoryName.value,
      public: directoryPublic.value,
    }).then(({data}) => {
      folders.value = folders.value.concat([data]).sort((a, b) => a.name > b.name ? 1 : -1);
      if (data.public) pubicFolders.value = pubicFolders.value.concat([data]).sort((a, b) => a.name > b.name ? 1 : -1);
      showModal.value = false;
      directoryName.value = '';
    }).catch(() => {
      message.info('Во время создании директории произошла ошибка');
      loader.error();
    }).finally(() => {
      loader.finish();
      isRequestEdit.value = false;
    });
  }
};

const clearItem = (id: number) => {
  images.value = images.value.filter((el) => el.id !== id);
};
</script>

<template>
  <div class="container">
    <template v-if="pubicFolders.length">
      <n-h2>Публичный контент</n-h2>
      <n-grid cols="1 s:2 m:3 l:4 xl:5 2xl:6" responsive="screen" x-gap="16" y-gap="16">
        <n-grid-item v-for="folder in pubicFolders" :key="folder.id">
          <router-link :to="`/folder/${folder.id}`">
            <n-card hovered class="mainPage__folder">
              <p style="white-space: nowrap; overflow: hidden; text-overflow: ellipsis">{{ folder.name }}</p>
              <n-icon>
                <eye-regular />
              </n-icon>
            </n-card>
          </router-link>
        </n-grid-item>
      </n-grid>
    </template>

    <n-h2>Моя коллекция</n-h2>
    <n-card>
      <n-space justify="space-between">
        <router-link to="/generate">
          <n-button type="warning">Подобрать</n-button>
        </router-link>
        <n-button type="success" @click="toCreateDirectory">
          <n-icon>
            <folder-plus />
          </n-icon>
        </n-button>
      </n-space>
    </n-card>
    <n-card style="margin-top: 32px" v-if="!images.length && !folders.length">Добавьте первое изображение</n-card>
    <n-grid style="margin-top: 32px" v-if="folders.length" cols="1 s:2 m:3 l:4 xl:5 2xl:6" responsive="screen" x-gap="16" y-gap="16">
      <n-grid-item v-for="folder in folders" :key="folder.id">
        <router-link :to="`/folder/${folder.id}`">
          <n-card hovered class="mainPage__folder">
            <p style="white-space: nowrap; overflow: hidden; text-overflow: ellipsis">{{ folder.name }}</p>

            <n-button-group>
              <n-button size="small" type="info" @click.stop.prevent="editFolder(folder)" secondary>
                <n-icon>
                  <edit-regular />
                </n-icon>
              </n-button>
              <n-button size="small" type="error" @click.stop.prevent="deleteFolder(folder)" secondary>
                <n-icon>
                  <trash-alt />
                </n-icon>
              </n-button>
            </n-button-group>
          </n-card>
        </router-link>
      </n-grid-item>
    </n-grid>

    <ImageList v-if="images.length" :images="images" @clear-item="clearItem" :folders="folders" :key="folders.length" hasAccess />

    <n-modal v-model:show="showModal">
      <n-card style="width: 600px;" :title="`${editedFolder ? 'Редактирование' : 'Создание'} директории`" preset="card">
        <n-input v-model:value="directoryName" type="text" placeholder="Название директории" />
        <n-space style="margin-top: 16px">
          <n-switch v-model:value="directoryPublic" />
          Публичная папка
        </n-space>
        <template #footer>
          <n-button-group style="display: flex; justify-content: end">
            <n-button type="warning" @click="cancel">
              Отменить
            </n-button>
            <n-button type="success" @click="saveFolder" :loading="isRequestEdit" :disabled="isRequestEdit">
              {{ editedFolder ? 'Сохранить' : 'Создать' }}
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
}
</style>
