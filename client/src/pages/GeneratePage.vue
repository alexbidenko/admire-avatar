<script setup lang="ts">
import {ref} from 'vue';
import {
  NImage,
  NIcon,
  NButton,
  NInput,
  NSpace, useMessage,
} from 'naive-ui';
import {Forward as ForwardRegular, SaveRegular} from '@vicons/fa';
import {generateImage, saveImage} from '~/api/images';

const message = useMessage();

const image = ref('');
const search = ref('');
const isRequest = ref(false);

const nextImage = () => {
  isRequest.value = true;
  generateImage({
    phrase: search.value,
    tags: ['test', 'tag'],
  }).then(({data}) => {
    image.value = data.source;
  }).catch(() => {
    message.error('Во время генерации изображения произошла ошибка');
  }).finally(() => {
    isRequest.value = false;
  });
};

const likeImage = () => {
  saveImage({source: image.value}).then(() => {
    message.info('Картинка сохранена');
  });
};
</script>

<template>
  <div class="container generatePage">
    <div class="squareImageCard">
      <n-image
          class="generatePage__image"
          :src="`/api/files/temporary/${image}`"
      />
    </div>
    <n-space justify="center" align="center" style="margin-top: 16px">
      <n-button @click="likeImage" strong secondary circle>
        <n-icon>
          <save-regular />
        </n-icon>
      </n-button>
      <n-input @keydown.enter="nextImage" v-model:value="search" type="text" placeholder="Найти аватарку" />
      <n-button class="generatePage__icon" @click="nextImage" type="success" strong secondary circle :disabled="isRequest" :loading="isRequest">
        <n-icon v-if="!isRequest">
          <forward-regular />
        </n-icon>
      </n-button>
    </n-space>
  </div>
</template>

<style lang="scss">
.generatePage {
  .squareImageCard {
    padding-top: 0;
    width: 80vw;
    max-width: 640px;
    height: 80vw;
    max-height: 640px;
    margin: 0 auto;
  }

  &__image img {
    object-fit: cover;
  }

  &__icon {
    .n-button__icon {
      margin: 0;
    }
  }
}
</style>
