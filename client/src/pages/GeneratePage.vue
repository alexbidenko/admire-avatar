<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref} from 'vue';
import {
  NImage,
  NIcon,
  NButton,
  NInput,
  NSpace, useMessage, NCard, NSkeleton,
} from 'naive-ui';
import {Forward as ForwardRegular, SaveRegular} from '@vicons/fa';
import {generateImage, saveImage} from '~/api/images';

const message = useMessage();

const card = ref<HTMLDivElement>();
const xState = ref(-1);
const xRotate = ref(0);
const state = ref(0);

const image = ref('');
const search = ref('');
const isRequest = ref(false);

const nextImage = () => {
  if (isRequest.value) return;
  isRequest.value = true;
  generateImage({
    phrase: search.value,
    tags: [],
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

const onTouchStart = (e: MouseEvent | TouchEvent) => {
  xState.value = ('touches' in e ? e.touches[0] : e).clientX;
};

onMounted(() => {
  const onTouchEnd = () => {
    if (state.value === 0) {
      xState.value = -1;
      xRotate.value = 0;
    }
  };
  const onTouchMove = (e: MouseEvent | TouchEvent) => {
    if (isRequest.value) return;
    if (xState.value !== -1) xRotate.value = (('touches' in e ? e.touches[0] : e).clientX - xState.value) / 28;
    if (Math.abs(xRotate.value) > 8) {
      window.removeEventListener('mousemove', onTouchMove);
      window.removeEventListener('mouseup', onTouchEnd);
      xState.value = -1;
      xRotate.value *= 2;

      if (xRotate.value > 0) {
        likeImage();
      }
      nextImage();

      state.value = 1;
      setTimeout(() => {
        state.value = 2;

        setTimeout(() => {
          state.value = 3;
          xRotate.value = 0;

          setTimeout(() => {
            state.value = 0;

            window.addEventListener('mousemove', onTouchMove);
            window.addEventListener('mouseup', onTouchEnd);
          }, 300);
        }, 100);
      }, 400);
    }
  };

  window.addEventListener('mousemove', onTouchMove);
  window.addEventListener('mouseup', onTouchEnd);
  window.addEventListener('touchmove', onTouchMove);
  window.addEventListener('touchend', onTouchEnd);

  card.value?.addEventListener('touchstart', onTouchStart);

  nextImage();

  onUnmounted(() => {
    window.removeEventListener('mousemove', onTouchMove);
    window.removeEventListener('mouseup', onTouchEnd);
    window.removeEventListener('touchmove', onTouchMove);
    window.removeEventListener('touchend', onTouchEnd);
  });
});

const opacity = computed(() => Math.max(0, (200 - Math.abs(xRotate.value) * 2) / 200));
</script>

<template>
  <div class="container generatePage">
    <div
      class="generatePage__movementCard"
      :class="{
        'generatePage__movementCard_moved': xState !== -1,
        'generatePage__movementCard_finish': state === 1,
        'generatePage__movementCard_reset': state === 2,
        'generatePage__movementCard_restart': state === 3,
        'generatePage__movementCard_toLeft': state === 1 && xRotate < 0,
        'generatePage__movementCard_toRight': state === 1 && xRotate > 0,
      }"
      :style="{ transform: `rotate(${xRotate}deg)`, opacity: opacity }"
      @mousedown="onTouchStart"
      ref="card"
    >
      <n-card class="squareImageCard">
        <n-skeleton height="100%" width="100%" v-if="state !== 1 && isRequest" />
        <n-image
          v-else
          class="generatePage__image"
          :src="`/api/files/temporary/${image}`"
        />
      </n-card>
    </div>
    <n-space justify="center" align="center" style="margin-top: 16px">
      <n-button class="generatePage__icon" @click="nextImage" type="success" strong secondary circle :disabled="isRequest" :loading="isRequest">
        <n-icon v-if="!isRequest">
          <forward-regular />
        </n-icon>
      </n-button>
      <n-input @keydown.enter="nextImage" v-model:value="search" type="text" placeholder="Введите фразу" style="width: 220px" />
      <n-button @click="likeImage" strong secondary circle :disabled="isRequest">
        <n-icon>
          <save-regular />
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

  &__image {
    pointer-events: all;

    img {
      pointer-events: none;
      object-fit: cover;
    }
  }

  &__icon {
    .n-button__icon {
      margin: 0 !important;
    }
  }

  .n-card__content {
    padding: 0 !important;
  }

  &__movementCard {
    transform-origin: 50% 240%;
    transition: transform 0.5s ease, opacity 0.3s ease;

    &_moved {
      transition: none;
    }

    &_toRight {
      transition: transform 0.3s ease, opacity 0.3s ease;
      transform: rotate(26deg) !important;
      opacity: 0 !important;
    }

    &_toLeft {
      transition: transform 0.5s ease, opacity 0.5s ease;
      transform: rotate(-26deg) !important;
      opacity: 0 !important;
    }

    &_finish {
      transition: transform 0.5s ease, opacity 0.3s ease;
    }

    &_reset {
      transform-origin: 50% 60%;
      transform: scale(0.5) !important;
      opacity: 0 !important;
      transition: none;
    }

    &_restart {
      transform-origin: 50% 60%;
      opacity: 1 !important;
      transform: scale(1) !important;
      transition: transform 0.3s ease, opacity 0.3s ease;
    }
  }
}
</style>
