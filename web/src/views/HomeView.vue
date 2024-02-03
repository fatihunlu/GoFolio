<template>
  <div class="h-screen">
    <div class="h-full grid grid-cols-2 items-center">
      <div class="justify-center">
        <div class="m-20">
          <stock-table/>
        </div>
      </div>
      <div class="justify-center">
        <div class="m-20">
          <div>section 2</div>
        </div>
      </div>
      <div class="justify-center">
        <div class="m-20">
          <div>section 3</div>
        </div>
      </div>
      <div class="justify-center">
        <div class="m-20">
          <div>section 4</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>

import { ref, computed } from 'vue'
import StockTable from '@/components/StockTable.vue'
const showAddStock = ref(false)



</script>
