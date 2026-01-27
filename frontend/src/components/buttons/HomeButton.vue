<script lang="ts">
import { getCookie } from '../../utils/cookie';
import { defineComponent, computed } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'BackButton',
  setup() {
    const router = useRouter()
    const handleBack = () => {
      router.push('/')
    }

    const buttonText = computed(() => {
      const lang = getCookie('lang')
      return lang === 'tw' ? '返回' : 'Back'
    })

    return {
      handleBack,
      buttonText
    }
  }
})
</script>

<template>
  <button class="back-button" @click="handleBack">
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
