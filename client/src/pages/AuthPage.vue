<script setup lang="ts">
import {signIn} from '~/api/users';
import {ref} from 'vue';
import {useRouter} from 'vue-router';
import {
  NButton, NCard, NInput, NSpace,
} from 'naive-ui';

const router= useRouter();

const email = ref('');
const password = ref('');

if (document.cookie.includes('authorized=true')) router.push('/');

const submit = () => {
  signIn(email.value, password.value).then(() => {
    router.push('/');
  });
};
</script>

<template>
  <div class="authPage">
    <n-card title="Admire Avatar">
      <form @submit.prevent="submit">
        <n-space vertical>
          Email
          <n-input placeholder="Введите email" name="email" v-model:value="email" />
          Пароль
          <n-input placeholder="Введите пароль" type="password" name="password" v-model:value="password" show-password-on="click" />
          <n-button type="primary" attr-type="submit" class="authPage__button">Войти</n-button>
        </n-space>
      </form>
    </n-card>
  </div>
</template>

<style lang="scss" scoped>
.authPage {
  height: 100%;
  display: flex;
  align-items: center;
  max-width: 400px;
  margin: 0 auto;

  &__button {
    width: 100%;
    margin-top: 16px;
  }
}
</style>
