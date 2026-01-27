<script lang="ts">
import { useRouter } from 'vue-router'
import { setCookie } from '../../utils/cookie'
import { defineComponent } from 'vue'
import PurchaseButton from '../buttons/PurchaseButton.vue'

export default defineComponent({
  name: 'DefaultLayout',
  components: {
    PurchaseButton
  },
  setup() {
    const router = useRouter()

    const changeLang = (lang: 'en' | 'tw') => {
      setCookie('lang', lang)
      router.push(`/${lang}`)
    }

    return {
      changeLang,
    }
  }
})
</script>

<template>
  <div class="inner-container">
    <div class="lang-switcher">
      <button @click="changeLang('en')">English</button>
      <button @click="changeLang('tw')">繁體中文</button>
    </div>
    <slot />
    <PurchaseButton/>
  </div>
</template>
