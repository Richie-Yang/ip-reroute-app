<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

export default defineComponent({
  name: 'PurchaseButton',
  setup() {
    const router = useRouter()
    const route = useRoute()

    const handlePurchase = () => {
      // console.log('Purchase clicked')
      router.push(`${route.path}/product`)
    }

    const buttonText = computed(() => {
      return route.path === '/tw' ? '購買' : 'Purchase'
    })

    return {
      handlePurchase,
      buttonText
    }
  }
})
</script>

<template>
  <button class="purchase-button" @click="handlePurchase">
    {{ buttonText }}
  </button>
</template>

<style scoped>
.purchase-button {
  padding: 1rem 2rem;
  border: 2px solid #667eea;
  background: #667eea;
  color: white;
  border-radius: 8px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 30px;
  width: 100%;
}

.purchase-button:hover {
  background: #5568d3;
  border-color: #5568d3;
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(102, 126, 234, 0.4);
}

.purchase-button:active {
  transform: translateY(0);
}
</style>
