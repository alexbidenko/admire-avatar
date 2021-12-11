<script setup lang="ts">
import {
  NButton, NCard, NInput, NSpace, useLoadingBar,
} from 'naive-ui';
import {computed, ref} from 'vue';
import {useRouter} from 'vue-router';
import $axios from '~/api';
import {UserType} from '~/types/user';

const loader = useLoadingBar();
const router = useRouter();

const password = ref('');
const repeatPassword = ref('');

const isValid = computed(() => password.value === repeatPassword.value && password.value.length >= 8);

const submit = () => {
  if (!isValid.value) return;
  loader.start();
  $axios.post<{ user: UserType; token: string }>('user/password', {password: password.value}).then(() => {
    router.push('/');
  }).catch(() => {
    loader.error();
  }).finally(() => {
    loader.finish();
  });
};
</script>

<template>
  <div class="passwordPage">
    <n-card title="Изменение пароля">
      <form @submit.prevent="submit">
        <n-space vertical>
          Пароль
          <n-input placeholder="Введите пароль" type="password" name="password" v-model:value="password" show-password-on="click" />
          Повторный пароль
          <n-input placeholder="Повторите пароль" type="password" name="password" v-model:value="repeatPassword" show-password-on="click" />
          <n-button type="primary" attr-type="submit" class="passwordPage__button" :disabled="!isValid">Сохранить</n-button>
        </n-space>
      </form>
    </n-card>
  </div>
</template>

<style lang="scss" scoped>
.passwordPage {
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
