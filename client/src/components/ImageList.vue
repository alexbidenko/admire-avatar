<script setup lang="ts">
import {computed, h, ref} from 'vue';
import {
  NButton,
  NCard,
  NButtonGroup,
  NImage,
  NGrid,
  NGridItem,
  NIcon,
  useLoadingBar,
  useMessage,
  NModal,
  NDropdown,
  NSelect,
  NAvatar,
  NText,
} from 'naive-ui';
import {deleteImage, createAvatar} from '~/api/images';
import {FolderType, ImageType} from '~/types/image';
import {
  Download as DownloadRegular,
  UserCircle,
  TrashAlt,
  Share as ShareRegular,
  ShareAlt,
} from '@vicons/fa';
import {useMainStore} from '~/store';
import $axios from '~/api';
import {UserType} from '~/types/user';
import {SelectBaseOption} from 'naive-ui/es/select/src/interface';
import md5 from 'md5';

const {images, folders, hasAccess} = defineProps<{
  images: ImageType[];
  folders: FolderType[];
  hasAccess: boolean;
}>();

// eslint-disable-next-line func-call-spacing
const emit = defineEmits<{
  (event: 'clear-item', value: number): void
}>();

const message = useMessage();
const loader = useLoadingBar();
const store = useMainStore();

const isRequestUser = ref(false);
const isRequestShare = ref(false);
const showModal = ref(false);
const showShare = ref(false);
const users = ref<(UserType & {label: string; value: number})[]>([]);
const selectedImage = ref<ImageType>();
const selectedUser = ref<number>();
const directoryName = ref('');

const deleteCurrentImage = (id: number) => {
  loader.start();
  deleteImage(id).then(() => {
    emit('clear-item', id);
  }).catch(loader.error).finally(loader.finish);
};

const selectAvatar = (image: ImageType) => {
  createAvatar(image.id).then(() => {
    images.forEach((el) => {
      el.main = false;
    });
    image.main = true;
    store.commit('setAvatar', store.state.avatar + 1);
  });
};

const cancel = () => {
  showShare.value = false;
  showModal.value = false;
  directoryName.value = '';
  selectedUser.value = undefined;
};

const options = computed(() => [
  {
    label: 'Переместить в директорию',
    key: 'title',
    disabled: true,
  },
  ...folders.map((el) => ({
    label: el.name,
    key: el.id,
  })),
]);

const handleSelect = (key: number, image: ImageType) => {
  loader.start();
  $axios.put(`images/${image.id}/folder/${key}`).then(() => {
    emit('clear-item', image.id);
  }).catch(loader.error).finally(loader.finish);
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

const shareImage = (image: ImageType) => {
  handleSearch('');
  showShare.value = true;
  selectedImage.value = image;
};

const shareImageRequest = () => {
  if (!selectedUser.value || !selectedImage.value || isRequestShare.value) return;
  isRequestShare.value = true;
  loader.start();
  $axios.post('/images/share', {
    userId: selectedUser.value,
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

const renderSingleSelectTag: any = (data: { option: SelectBaseOption<UserType> }) => (h as any)(
  'div',
  {
    style: {
      display: 'flex',
      alignItems: 'center',
    },
  },
  [
    h(NAvatar, {
      src: `/api/admire-avatar/${(md5 as any)(data.option.email)}`,
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
  <n-grid style="margin-top: 32px" v-if="images.length" cols="1 s:2 m:3 l:4 xl:5 2xl:6" responsive="screen" x-gap="16" y-gap="16">
    <n-grid-item
        v-for="image in images"
        :key="image.id"
    >
      <n-card content-style="padding: 0">
        <div class="squareImageCard">
          <n-image
              class="imageList__image"
              :class="{'imageList__image_selected': image.main}"
              :src="`/api/files/images/${image.source}`"
          />
        </div>
        <n-button-group class="imageList__actions" v-if="hasAccess">
          <n-button type="info" @click="shareImage(image)">
            <n-icon>
              <share-alt />
            </n-icon>
          </n-button>
          <n-button type="success" @click="selectAvatar(image)">
            <n-icon>
              <user-circle />
            </n-icon>
          </n-button>
          <a :download="image.source" :href="`/api/files/images/${image.source}`">
            <n-button type="info">
              <n-icon>
                <download-regular />
              </n-icon>
            </n-button>
          </a>
          <n-dropdown trigger="hover" @select="handleSelect($event, image)" :options="options">
            <n-button>
              <n-icon>
                <share-regular />
              </n-icon>
            </n-button>
          </n-dropdown>
          <n-button type="error" @click="deleteCurrentImage(image.id)">
            <n-icon>
              <trash-alt />
            </n-icon>
          </n-button>
        </n-button-group>
      </n-card>
    </n-grid-item>
  </n-grid>

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

.imageList {
  &__actions {
    padding-bottom: 10px;
    width: 100%;
    justify-content: center;
  }

  &__image {
    position: relative;

    &::after {
      content: "";
      border: 4px solid #63e2b7;
      border-radius: 5px;
      box-sizing: border-box;
      opacity: 0;
      pointer-events: none;
      transition: opacity 0.3s ease;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      display: block;
      position: absolute;
    }

    &_selected::after {
      opacity: 1;
    }
  }
}
</style>
