<script setup lang="ts">
import {ref} from 'vue';
import {
  NButton, NButtonGroup,
  NCard,
  NGi,
  NGrid,
  NH1, NIcon,
  NImage,
  NInput,
  NInputGroup,
  NInputNumber,
  useLoadingBar,
  useMessage,
} from 'naive-ui';
import {generatePrints, getPrints} from '~/api/prints';
import {ImageType} from '~/types/image';
import {TrashAltRegular} from '@vicons/fa';
import {deleteImage} from '~/api/images';

const loader = useLoadingBar();
const message = useMessage();

const phrase = ref('');
const count = ref(1);
const images = ref<ImageType[]>([]);

loader.start();
getPrints().then(({data}) => {
  images.value = data;
}).catch(loader.error).finally(loader.finish);

const generate = () => {
  loader.start();
  generatePrints({
    count: count.value,
    phrase: phrase.value,
    tags: [],
  }).then(() => {
    message.info('Генерация принтов запущено');
  }).catch(() => {
    loader.error();
    message.error('Во время генерация принтов произшла ошибка');
  }).finally(loader.finish);
};

const deletePrint = (id: number) => {
  loader.start();
  deleteImage(id).then(() => {
    images.value = images.value.filter((el) => el.id !== id);
  }).catch(loader.error).finally(loader.finish);
};
</script>

<template>
  <n-h1>Генерация принтов</n-h1>
  <n-card style="margin-bottom: 32px">
    <n-input-group>
      <n-input placeholder="Поисковая фраза" v-model:value="phrase" />
      <n-input-number v-model:value="count" :min="1" />
    </n-input-group>
    <n-button @click="generate" type="primary" style="margin-top: 16px">Сгенерировать</n-button>
  </n-card>

  <n-grid cols="2 s:3 m:4 l:5 xl:6 2xl:8" responsive="screen" x-gap="16" y-gap="16">
    <n-gi v-for="image in images" :key="image.id">
      <n-card>
        <n-image :src="`/api/files/images/${image.source}`" class="printsPage__image" />
        <n-button-group>
          <n-button type="error" @click="deletePrint(image.id)">
            <template #icon>
              <n-icon><trash-alt-regular /></n-icon>
            </template>
          </n-button>
        </n-button-group>
      </n-card>
    </n-gi>
  </n-grid>
</template>

<style lang="scss">
.printsPage {
  &__image img {
    max-width: 100%;
  }
}
</style>