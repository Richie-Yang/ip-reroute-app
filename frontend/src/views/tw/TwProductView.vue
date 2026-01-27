<script setup lang="ts">
import PurchaseLayout from '../../components/layouts/PurchaseLayout.vue'
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getGeoIP } from '../../utils/geoip';
import { getCookie, setCookie } from '../../utils/cookie';

const router = useRouter()
onMounted(async () => {
  let country = getCookie('country')
  if (!country) {
    const geoIP = await getGeoIP();
    setCookie('country', geoIP.country)
  }
  country = getCookie('country')
  switch (country) {
    case 'TW':
      router.push('/tw/product')
      break
    default:
      router.push('/en/product')
      break
  }
})
</script>

<template>
  <PurchaseLayout>
    <div class="content">繁體中文產品</div>
  </PurchaseLayout>
</template>
