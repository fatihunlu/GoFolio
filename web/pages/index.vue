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
          <stock-table/>
        </div>
      </div>
      <div class="justify-center">
        <div class="m-20">
          <stock-table/>
        </div>
      </div>
      <div class="justify-center">
        <div class="m-20">
          <stock-table/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import StockTable from '../components/StockTable.vue'

export default {
  name: 'IndexPage',
    components: [
        StockTable  
    ]
}
</script>
