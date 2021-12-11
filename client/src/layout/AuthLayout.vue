<script setup lang="ts">
import {useRouter} from 'vue-router';
import {useMainStore} from '~/store';
import {logout, getUser} from '~/api/users';
import {
  NLayoutHeader, NButton, NSpace, NLayout, NButtonGroup, NScrollbar, NDropdown, NAvatar,
} from 'naive-ui';
import {useRoute} from 'vue-router';
import {getAvatar} from '~/api/images';

const LINKS = [
  {
    label: 'Аватарки',
    key: 'main',
    to: '/',
  },
  {
    label: 'Принты',
    key: 'Prints',
    to: '/prints',
  },
];

const OPTIONS = [
  {
    label: 'Выйти',
    key: 'logout',
  },
];

const route = useRoute();
const router = useRouter();
const store = useMainStore();

const onSelectDropdownOption = (key: string) => {
  switch (key) {
  case 'logout':
    logout().then(() => {
      store.commit('logout');
      router.push('/auth');
    });
    break;
  }
};

getUser().then(({data}) => {
  store.commit('setUser', data);
});

getAvatar().then(({data}) => {
  store.commit('setAvatar', data);
});
</script>

<template>
  <n-layout class="authorizedLayout">
    <n-layout-header position="absolute" bordered>
      <n-space justify="space-between" align="center">
        <n-button-group>
          <router-link v-for="link in LINKS" :to="link.to" :key="link.key" v-slot="{ isExactActive }">
            <n-button :bordered="false" :tertiary="isExactActive || (link.key === 'main' && route.fullPath.includes('/generate'))">{{ link.label }}</n-button>
          </router-link>
        </n-button-group>
        <n-dropdown :options="OPTIONS" v-if="store.state.isAuthorized" :on-select="onSelectDropdownOption" placement="bottom-end">
          <n-button size="large" :bordered="false" style="padding-right: 0">
            <span style="margin-right: 16px">
            {{ store.state.user.name }}
            </span>
            <n-avatar
              round
              :size="48"
              :src="`/api/files/images/${store.state.avatar?.source}`"
            />
          </n-button>
        </n-dropdown>
      </n-space>
    </n-layout-header>
    <n-scrollbar class="authorizedLayout__scrollbar">
      <router-view />
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
