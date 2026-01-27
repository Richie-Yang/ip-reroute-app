<script lang="ts">

import { defineComponent } from 'vue'
import BackButton from '../buttons/HomeButton.vue';

export default defineComponent({
  name: 'PurchaseLayout',
  components: {
    BackButton
  },
  setup() {

  }
})
</script>

<template>
  <div class="inner-container">
    <BackButton />
    <slot />
  </div>
</template>
