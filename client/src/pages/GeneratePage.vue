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
  <div class="container">
    <n-space justify="center">
      <div class="squareImageCard">
        <n-image
            v-if="image"
            class="generatePage__image"
            :src="`/api/files/temporary/${image}`"
        />
      </div>
    </n-space>
    <n-space justify="center" align="center">
      <n-button @click="likeImage" strong secondary circle>
        <n-icon>
          <save-regular />
        </n-icon>
      </n-button>
      <n-input @keydown.enter="nextImage" v-model:value="search" type="text" placeholder="Найти аватарку" />
      <n-button @click="nextImage" type="success" strong secondary circle :loader="isRequest">
        <n-icon>
          <forward-regular />
        </n-icon>
      </n-button>
    </n-space>
  </div>
</template>

<style lang="scss">
.generatePage {
  &__image img {
    width: 100%;
    max-width: 100%;
    object-fit: cover;

    @media (min-width: 500px) {
      width: 400px;
    }

    @media (min-width: 900px) {
      width: 450px;
    }

    @media (min-width: 1300px) {
      width: 720px;
    }
  }
}
</style>
