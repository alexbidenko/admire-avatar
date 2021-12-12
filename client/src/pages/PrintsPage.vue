<script setup lang="ts">
import {onUnmounted, ref} from 'vue';
import {
  NButton, NButtonGroup,
  NCard,
  NGi,
  NGrid,
  NH1, NIcon,
  NImage,
  NInput,
  NInputGroup,
  NInputNumber, NSelect,
  useLoadingBar,
  useMessage,
} from 'naive-ui';
import {clear, generatePrints, getPrints, saveAsAvatar} from '~/api/prints';
import {ImageType} from '~/types/image';
import {Download as DownloadRegular, SaveRegular, TrashAltRegular} from '@vicons/fa';
import {deleteImage} from '~/api/images';
import $axios from '~/api';

const loader = useLoadingBar();
const message = useMessage();

const phrase = ref('');
const count = ref(1);
const tag = ref<string>();
const tags = ref<{ value: string; label: string }[]>([]);
const images = ref<ImageType[]>([]);
const generateCount = ref(0);

const ws = new WebSocket(`${location.protocol === 'https:' ? 'wss' : 'ws'}://${location.host}/api/prints/channel`);

ws.onmessage = (e: MessageEvent<string>) => {
  const image = JSON.parse(e.data);
  generateCount.value -= 1;
  if (image) {
    if (!images.value.find((el) => el.id === image.id)) images.value.unshift(image);
  } else message.error('Во время генерации одного из изобращений произошла ошибка.');
};

onUnmounted(() => {
  ws.close();
});

loader.start();
Promise.all([
  $axios.get('images/tags').then(({data}) => {
    tags.value = data;
  }),
  getPrints().then(({data}) => {
    images.value = data.images;
    generateCount.value = data.generate;
  }),
]).catch(loader.error).finally(loader.finish);

const wordForms = (count: number, words: string[]) => {
  if (count % 10 === 1) return words[0];
  if (count % 10 > 1 && count % 10 < 5) return words[1];
  return words[2];
};

const generate = () => {
  loader.start();
  generatePrints({
    count: count.value,
    phrase: phrase.value,
    tags: tag.value ? [tag.value] : [],
  }).then(() => {
    generateCount.value += count.value;
    message.info('Генерация принтов запущена');
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

const toAvatar = (id: number) => {
  loader.start();
  saveAsAvatar(id).then(() => {
    images.value = images.value.filter((el) => el.id !== id);
  }).catch(loader.error).finally(loader.finish);
};

const deleteAll = () => {
  loader.start();
  clear().then(() => {
    images.value = [];
  }).catch(loader.error).finally(loader.finish);
};

const timer = setInterval(() => {
  getPrints().then(({data}) => {
    images.value = data.images;
    generateCount.value = data.generate;
  });
}, 8000);
onUnmounted(() => {
  clearInterval(timer);
});
</script>

<template>
  <div class="printsPage">
    <n-h1>Генерация контента</n-h1>
    <n-card style="margin-bottom: 32px">
      <n-input-group>
        <n-input placeholder="Поисковая фраза" v-model:value="phrase" />
        <n-select placeholder="Тема" v-model:value="tag" :options="tags" class="printsPage__select" />
        <n-input-number v-model:value="count" :min="1" style="width: 140px" />
      </n-input-group>
      <n-button-group style="margin-top: 16px">
        <n-button @click="generate" type="primary">Сгенерировать</n-button>
        <n-button @click="deleteAll" type="error">Очистить</n-button>
        <a href="/api/prints/archive" download="archive.zip">
          <n-button type="info">
            <n-icon>
              <download-regular />
            </n-icon>
          </n-button>
        </a>
      </n-button-group>
    </n-card>

    <n-card v-if="!generateCount && !images.length">
      Введите поисковую фразу и выберите стиль и количество изображений. Мы создадим для вас уникальный контент!
      Генерация происходит в фоновом режиме и вы можете не ограничивать себя в изучении нашего приложения.
    </n-card>

    <n-card v-else-if="generateCount" style="margin-bottom: 24px">
      В процессе генерации {{ generateCount }} {{ wordForms(generateCount, ['изображение', 'изображения', 'изображений']) }}.
      Вы увидите их на этой странице через некоторое время, а пока можете изучить публичный контент других пользователей.
    </n-card>

    <n-grid cols="2 s:3 m:4 l:5 xl:6 2xl:8" responsive="screen" x-gap="16" y-gap="16">
      <n-gi v-for="image in images" :key="image.id">
        <n-card content-style="padding: 0">
          <div class="squareImageCard">
            <n-image
                :src="`/api/files/images/${image.source}`"
                class="printsPage__image"
            />
          </div>
          <n-button-group style="padding-bottom: 10px; width: 100%; justify-content: center">
            <n-button type="success" @click="toAvatar(image.id)">
              <template #icon>
                <n-icon><save-regular /></n-icon>
              </template>
            </n-button>
            <n-button type="error" @click="deletePrint(image.id)">
              <template #icon>
                <n-icon><trash-alt-regular /></n-icon>
              </template>
            </n-button>
          </n-button-group>
        </n-card>
      </n-gi>
    </n-grid>
  </div>
</template>

<style lang="scss">
.printsPage {
  &__image img {
    max-width: 100%;
  }

  &__select {
    width: fit-content;
    min-width: 160px;
  }
}
</style>
