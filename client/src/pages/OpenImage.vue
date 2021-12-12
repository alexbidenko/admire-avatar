<script setup lang="ts">
import {ref} from 'vue';
import {
  NImage,
  NCard,
  useLoadingBar,
} from 'naive-ui';
import $axios from '~/api';
import {useRoute} from 'vue-router';
import {ImageType} from '~/types/image';

const loader = useLoadingBar();
const route = useRoute();

const image = ref<ImageType>();

loader.start();
$axios.get(`images/${route.params.id}/data`).then(({data}) => {
  image.value = data;
}).catch(loader.error).finally(loader.finish);
</script>

<template>
  <div class="container openImagePage">
    <div
        class="openImagePage__movementCard"
        ref="card"
    >
      <n-card class="squareImageCard">
        <n-image
            class="openImagePage__image"
            :src="`/api/files/images/${image?.source}`"
        />
      </n-card>
    </div>
  </div>
</template>

<style lang="scss">
.openImagePage {
  padding-top: 32px;

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
}
</style>
