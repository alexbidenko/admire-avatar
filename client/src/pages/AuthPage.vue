<script setup lang="ts">
import {
  NButton, NCard, NInput, NSpace,
} from 'naive-ui';
</script>

<template>
  <div class="authPage">
    <n-card title="Admire Avatar">
      <form @submit.prevent="submit">
        <n-space vertical>
          Email
          <n-input placeholder="Введите email" name="email" />
          Пароль
          <n-input placeholder="Введите пароль" type="password" name="password" />
          <n-button type="info" attr-type="submit" class="authPage__button">Войти</n-button>
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
