<script setup lang="ts">
import {ref} from 'vue';
import {NImage, NIcon, NButton} from 'naive-ui';
import {Forward as ForwardRegular, SaveRegular} from '@vicons/fa';
import {generateImage, saveImage} from '~/api/images';

const image = ref('');

const nextImage = () => {
  generateImage({
    phrase: 'test',
    tags: ['test', 'tag'],
  }).then(({data}) => {
    image.value = data.source;
  });
};

const likeImage = () => {
  saveImage({source: image.value}).then(() => {
    alert('Картинка сохранена');
  });
};
</script>

<template>
  <div class="container">
    <div class="containerImages">
      <n-image
          width="100"
          src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
      />
    </div>
    <div class="buttonsContainer">
      <n-button @click="likeImage" class="buttonAction" strong secondary circle>
        <n-icon>
          <save-regular />
        </n-icon>
      </n-button>
      <n-button @click="nextImage" class="buttonAction" strong secondary circle>
        <n-icon>
          <forward-regular />
        </n-icon>
      </n-button>
    </div>
  </div>
</template>

<style lang="scss">
.containerImages {
  display: flex;
  justify-content: center;
}

.buttonsContainer {
  display: flex;
  justify-content: space-between;
}

.buttonAction {
  font-size: 70px;
  width: 90px;
  height: 90px;
}
</style>
