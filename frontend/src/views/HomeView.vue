<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getCookie } from '../utils/cookie'

const router = useRouter()

onMounted(() => {
  const lang = getCookie('lang')
  if (lang === 'en')   router.push('/en')
   else if (lang === 'tw')   router.push('/tw')
})
</script>

<template>
  <div>Home</div>
</template>
