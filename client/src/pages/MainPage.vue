<script setup lang="ts">
import {computed, h, onMounted, onUnmounted, ref} from 'vue';
import {
  NButton,
  NCard,
  NButtonGroup,
  NImage,
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
  NDropdown,
  NSelect,
  NAvatar,
  NText,
} from 'naive-ui';
import {getImages, deleteImage, createAvatar} from '~/api/images';
import {FolderType, ImageType} from '~/types/image';
import {
  Download as DownloadRegular,
  UserCircle,
  TrashAlt,
  Share as ShareRegular,
  FolderPlus,
  ShareAlt,
} from '@vicons/fa';
import {useMainStore} from '~/store';
import $axios from '~/api';
import {UserType} from '~/types/user';
import {SelectBaseOption} from 'naive-ui/es/select/src/interface';
import md5 from 'md5';

const COUNT = 40;

const dialog = useDialog();
const message = useMessage();
const loader = useLoadingBar();
const store = useMainStore();
const page = ref(0);
const isRequest = ref(true);
const isFinish = ref(false);

const isRequestUser = ref(false);
const isRequestShare = ref(false);
const showModal = ref(false);
const showShare = ref(false);
const folders = ref<FolderType[]>([]);
const images = ref<ImageType[]>([]);
const users = ref<(UserType & {label: string; value: number})[]>([]);
const selectedImage = ref<ImageType>();
const selectedUser = ref<UserType>();
const directoryName = ref('');

loader.start();
Promise.all([
  getImages(0, COUNT).then(({data}) => {
    isRequest.value = false;
    images.value = data;
    if (data.length < COUNT) isFinish.value = true;
  }),
  $axios.get('folders').then(({data}) => {
    isRequest.value = false;
    folders.value = data;
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

const deleteCurrentImage = (id: number) => {
  loader.start();
  deleteImage(id).then(() => {
    images.value = images.value.filter((el) => el.id !== id);
  }).catch(loader.error).finally(loader.finish);
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

const options = computed(() => [
  {
    label: 'Переместить в директорию',
    key: 'title',
    disabled: true,
  },
  ...folders.value.map((el) => ({
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

const shareImage = (image: ImageType) => {
  showShare.value = true;
  selectedImage.value = image;
};

const shareImageRequest = () => {
  if (!selectedUser.value || !selectedImage.value || isRequestShare.value) return;
  isRequestShare.value = true;
  loader.start();
  $axios.post('/images/share', {
    userId: selectedUser.value.id,
    imageId: selectedImage.value.id,
  }).then(() => {
    selectedUser.value = undefined;
    showShare.value = false;
    message.success('Сообщение успешно отправлено');
  }).catch(() => {
    message.error('Во время отправки изображения произошла ошибка');
    loader.error();
  }).finally(() => {
    loader.finish();
    isRequestShare.value = false;
  });
};

const handleSearch = (query: string) => {
  isRequestUser.value = true;
  $axios.get<UserType[]>('users', {
    params: {
      email: query,
    },
  }).then(({data}) => {
    users.value = data.map((el) => ({
      ...el,
      label: el.name,
      value: el.id,
    }));
  }).finally(() => {
    isRequestUser.value = false;
  });
};

handleSearch('');

const renderSingleSelectTag = (data: { option: SelectBaseOption<UserType> }) => (h as any)(
  'div',
  {
    style: {
      display: 'flex',
      alignItems: 'center',
    },
  },
  [
    h(NAvatar, {
      src: `/api/admire-avatar/${md5(data.option.email)}`,
      round: true,
      size: 24,
      style: {
        marginRight: '12px',
      },
    }),
    data.option.name,
  ],
);

const renderLabel = (option: UserType) => (h as any)(
  'div',
  {
    style: {
      display: 'flex',
      alignItems: 'center',
    },
  },
  [
    h(NAvatar, {
      src: `/api/admire-avatar/${md5(option.email)}`,
      round: true,
      size: 'small',
    }),
    h(
      'div',
      {
        style: {
          marginLeft: '12px',
          padding: '4px 0',
        },
      },
      [
        (h as any)('div', null, [option.name]),
        h(
          NText,
          {depth: 3, tag: 'div'},
          {
            default: () => option.email,
          },
        ),
      ],
    ),
  ],
);
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
    <n-grid style="margin-top: 32px" v-if="images.length" cols="2 s:3 m:4 l:5 xl:6 2xl:8" responsive="screen" x-gap="16" y-gap="16">
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
          <n-button-group class="mainPage__actions">
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
            <n-button type="info" @click="shareImage(image)">
              <n-icon>
                <share-alt />
              </n-icon>
            </n-button>
          </n-button-group>
        </n-card>
      </n-grid-item>
    </n-grid>

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

    <n-modal v-model:show="showShare">
      <n-card style="width: 600px;" title="Поделиться изображением" preset="card">
        <n-select
            :render-label="renderLabel"
            :render-tag="renderSingleSelectTag"
            filterable
            placeholder="Введите Email пользователя"
            :options="users"
            :loading="isRequestUser"
            clearable
            remote
            @search="handleSearch"
            v-model:value="selectedUser"
        />
        <template #footer>
          <n-button-group style="display: flex; justify-content: end">
            <n-button type="warning" @click="cancel">
              Отменить
            </n-button>
            <n-button type="success" @click="shareImageRequest" :loading="isRequestShare" :disabled="isRequestShare">
              Поделиться
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
