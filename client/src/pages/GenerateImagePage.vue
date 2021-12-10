<script setup lang="ts">
import {NImage, NIcon, NButton} from 'naive-ui';
import {SaveRound, ArrowForwardFilled} from '@vicons/material';
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
      <n-button class="buttonAction" strong secondary circle>
        <n-icon>
          <SaveRound />
        </n-icon>
      </n-button>
      <n-button class="buttonAction" strong secondary circle>
        <n-icon>
          <ArrowForwardFilled />
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
