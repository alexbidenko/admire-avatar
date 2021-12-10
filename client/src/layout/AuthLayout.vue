<script setup lang="ts">
import {onMounted} from 'vue';
import {useRouter} from 'vue-router';
import {useMainStore} from '~/store';
import {logout, getUser} from '~/api/users';
import {
  NLayoutHeader, NButton, NSpace, NLayout, NH2, NButtonGroup,
} from 'naive-ui';
import {useRoute} from 'vue-router';

const route = useRoute();
const router = useRouter();
const store = useMainStore();

const logoutUser = () => {
  logout().then(() => {
    router.push('/auth');
  });
};

onMounted(() => {
  getUser().then(({data}) => {
    store.commit('setUser', data);
  });
});
</script>

<template>
  <n-layout class="authorizedLayout">
    <n-layout-header position="absolute" bordered>
      <n-space justify="space-between">
        <n-h2>{{store.state.user.name}}</n-h2>
        <n-button-group>
          <router-link to="/">
            <n-button v-if="route.path === '/generate'" type="info">Вернуться</n-button>
          </router-link>
        </n-button-group>
        <n-button type="warning" @click="logoutUser">Выйти</n-button>
      </n-space>
    </n-layout-header>
    <router-view />
  </n-layout>
</template>

<style lang="scss">
.authorizedLayout {
  padding-top: var(--header-height);
}

.n-layout-header {
  grid-template-rows: calc(var(--header-height) - 1px);
  display: grid;
  padding: 0 32px;
  align-items: center;
  z-index: 10;
}
</style>
