<script setup lang="ts">
import {
  NButton, NCard, NInput, NSpace,
} from 'naive-ui';
</script>

<template>
  <div class="authPage">
    <n-card title="Регистрация">
      <form @submit.prevent="submit">
        <n-space vertical>
          Имя
          <n-input placeholder="Введите ваше имя" name="name" />
          Email
          <n-input placeholder="Введите email" name="email" />
          Пароль
          <n-input placeholder="Введите пароль" type="password" name="password" />
          Повторный пароль
          <n-input placeholder="Повторите пароль" type="password" name="repeatPassword" />
          <n-button type="info" attr-type="submit" class="authPage__button">Зарегистрироваться</n-button>
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
