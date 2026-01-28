<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getCookie } from '../utils/cookie'

const router = useRouter()

onMounted(() => {
  const lang = getCookie('lang')
  switch (lang) {
    case 'tw':
      router.push('/tw')
      break
    case 'en':
    default:
      router.push('/en')
      break
  }
})
</script>

<template>
  <div>Home</div>
</template>
