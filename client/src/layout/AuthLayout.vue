<script setup lang="ts">
import {h, onUnmounted, ref} from 'vue';
import {useRouter} from 'vue-router';
import {useMainStore} from '~/store';
import {logout, getUser} from '~/api/users';
import {
  NLayoutHeader,
  NButton,
  NSpace,
  NLayout,
  NButtonGroup,
  NScrollbar,
  NDropdown,
  NAvatar,
  NBackTop,
  useNotification,
  NText, NDrawer, NDrawerContent, NIcon, NMenu,
} from 'naive-ui';
import {useRoute} from 'vue-router';
import md5 from 'md5';
import {Bars as BarsRegular} from '@vicons/fa';

const screenWidth = ref(0);
const showMenu = ref(false);

const LINKS = [
  {
    label: 'Библиотека',
    key: 'main',
    to: '/',
  },
  {
    label: 'Генератор',
    key: 'Prints',
    to: '/prints',
  },
];

const OPTIONS = [
  {
    label: 'Изменить пароль',
    key: 'password',
    to: '/password',
  },
  {
    label: 'Выйти',
    key: 'logout',
  },
];

const route = useRoute();
const router = useRouter();
const store = useMainStore();
const notification = useNotification();

const onSelectDropdownOption = (key: string) => {
  switch (key) {
  case 'logout':
    logout().then(() => {
      store.commit('logout');
      router.push('/auth');
    });
    break;
  default:
    router.push(OPTIONS.find((el) => el.key === key)?.to || '');
    break;
  }
};

getUser().then(({data}) => {
  store.commit('setUser', data);
  store.commit('setAvatar', store.state.avatar + 1);
});

const ws = new WebSocket(`${location.protocol === 'https:' ? 'wss' : 'ws'}://${location.host}/api/user/channel`);

ws.onmessage = (e: MessageEvent<string>) => {
  const message = JSON.parse(e.data);
  const n = notification.create({
    title: 'С вами поделились изображением',
    content: `Пользователь ${message.user.name} прислал вам изображение. Вы можете посмотреть его в соответствующей директории вашей колекции.`,
    action: () => h(
      NButton,
      {
        text: true,
        type: 'primary',
        onClick: () => {
          router.push(`/folder/${message.image.folder_id}`);
          n.destroy();
        },
      },
      {
        default: () => 'Перейти',
      },
    ),
  });
};

const navigate = (_: string, item: any) => {
  router.push(item.to);
  showMenu.value = false;
};

const onResize = () => {
  screenWidth.value = window.innerWidth;
};
window.addEventListener('resize', onResize);
onUnmounted(() => {
  ws.close();
  window.removeEventListener('resize', onResize);
});
onResize();
</script>

<template>
  <n-layout class="authorizedLayout">
    <n-layout-header position="absolute" bordered>
      <n-space justify="space-between" align="center">
        <n-button quaternary @click="showMenu = !showMenu" v-if="screenWidth < 640" size="large">
          <n-icon>
            <bars-regular />
          </n-icon>
        </n-button>
        <n-button-group v-if="screenWidth >= 640">
          <router-link v-for="link in LINKS" :to="link.to" :key="link.key" v-slot="{ isExactActive }">
            <n-button :bordered="false" :tertiary="isExactActive || (link.key === 'main' && route.fullPath.includes('/generate'))">{{ link.label }}</n-button>
          </router-link>
        </n-button-group>
        <n-dropdown :options="OPTIONS" v-if="store.state.isAuthorized && screenWidth >= 640" :on-select="onSelectDropdownOption" placement="bottom-end">
          <n-button size="large" :bordered="false" style="padding-right: 0">
            <div style="margin-right: 16px; display: flex; flex-direction: column; align-items: end">
              <n-text style="margin-bottom: 4px">{{ store.state.user.name }}</n-text>
              <n-text :depth="3">{{ store.state.user.email }}</n-text>
            </div>
            <n-avatar
                round
                :size="48"
                :src="`/api/admire-avatar/${md5(store.state.user.email)}?key=${store.state.avatar}`"
            />
          </n-button>
        </n-dropdown>
      </n-space>
    </n-layout-header>

    <n-drawer v-model:show="showMenu" :width="320" placement="left">
      <n-drawer-content title="Stoner" body-content-style="padding-left: 0; padding-right: 0">
        <template #header>
          <div style="padding: 8px 0">
            <n-avatar
                round
                :size="80"
                :src="`/api/admire-avatar/${md5(store.state.user.email)}?key=${store.state.avatar}`"
            />
            <n-text style="margin-bottom: 8px; margin-top: 24px; width: 100%; display: block">{{ store.state.user.name }}</n-text>
            <n-text :depth="3">{{ store.state.user.email }}</n-text>
          </div>
        </template>
        <n-menu
            :on-update:value="navigate"
            :options="LINKS"
        />
      </n-drawer-content>
    </n-drawer>

    <n-scrollbar class="authorizedLayout__scrollbar scrollContainer">
      <router-view v-slot="{ Component }">
        <transition name="scale" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>

      <n-back-top />
    </n-scrollbar>
  </n-layout>
</template>

<style lang="scss">
.authorizedLayout {
  padding-top: var(--header-height);

  &__scrollbar {
    & > .n-scrollbar-container {
      box-sizing: border-box;

      & > .n-scrollbar-content {
        padding: 24px 16px;

        @media (min-width: 640px) {
          padding: 24px 10vw;
        }
      }
    }
  }
}

.n-layout-header {
  grid-template-rows: calc(var(--header-height) - 1px);
  display: grid;
  padding: 0 32px;
  align-items: center;
  z-index: 10;
}
</style>
