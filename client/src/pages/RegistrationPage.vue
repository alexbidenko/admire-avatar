<script setup lang="ts">
import {signUp} from '~/api/users';
import {useRouter} from 'vue-router';
import {ref} from 'vue';
import {
  NButton, NCard, NInput, NSpace,
} from 'naive-ui';

const router = useRouter();

const name = ref('');
const password = ref('');
const repeatPassword = ref('');
const email = ref('');

if (document.cookie.includes('authorized=true')) router.push('/');

const submit = () => {
  signUp({name: name.value, email: email.value, password: password.value})
    .then(() => {
      router.push('/');
    });
};
</script>

<template>
  <div class="authPage">
    <n-card title="Регистрация">
      <form @submit.prevent="submit">
        <n-space vertical>
          Имя
          <n-input placeholder="Введите ваше имя" name="name" v-model:value="name" />
          Email
          <n-input placeholder="Введите email" name="email" v-model:value="email" />
          Пароль
          <n-input placeholder="Введите пароль" type="password" name="password" v-model:value="password" show-password-on="click" />
          Повторный пароль
          <n-input placeholder="Повторите пароль" type="password" name="repeatPassword" v-model:value="repeatPassword" show-password-on="click" />
          <n-button type="primary" attr-type="submit" class="authPage__button" :disabled="password !== repeatPassword || password.length < 8">Зарегистрироваться</n-button>
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
