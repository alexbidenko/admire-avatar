<script setup lang="ts">
import {useRouter} from 'vue-router';
import {logout} from '~/api/users';
import {
  NLayoutHeader, NButton, NSpace, NLayout,
} from 'naive-ui';

const router = useRouter();

const logoutUser = () => {
  logout().then(() => {
    router.push('/auth');
  });
};
</script>

<template>
  <n-layout class="authorizedLayout">
    <n-layout-header position="absolute" bordered>
      <n-space justify="space-between" align="center">
        <n-button type="warning" @click="logoutUser">Выйти</n-button>
      </n-space>
    </n-layout-header>
  </n-layout>
</template>

<style lang="scss">
.authorizedLayout {
  padding-top: 64px;
}
</style>
