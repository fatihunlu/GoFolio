<template>
  <div>
    <div class="h-screen w-screen fixed top-0 left-0 opacity-[.86]">
    </div>
    <div class="fixed top-1/3 left-1/3 z-20">
      <div class="flex w-[630px] h-[370px] mt-10 ml-5 flex-col modal bg-white">
        <div class="flex flex-row justify-between mt-4">
          <span></span>
          <span class="text-[22px] text-[#9b45b2]">Add Stock</span>
          <svg width="20px" height="20px" class="mt-1 mr-[30px] cursor-pointer hover:bg-opacity-75" @click="close"
            viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M20 20L4 4.00003M20 4L4.00002 20" stroke="#9b45b2" stroke-width="2" stroke-linecap="round" />
          </svg>
        </div>

        <div class="flex flex-col h-full justify-between">
          <div>
            <div class="flex flex-row justify-center items-center mt-4">
              <div class="w-5/6">
                <stock-dropdown @selected="selectedData" :ignoredList="selectedItemIds"
                  :clearInputWhenClicked="false">
                </stock-dropdown>
              </div>
            </div>

            <div class="flex justify-center w-full mt-4">
              <div class="flex items-center justify-center w-5/6">
                <div class="w-full mr-4">
                  <input v-model="price" type="number"
                    class="border border-gray-400 text-gray-900 text-sm rounded-lg focus:border-indigo-600 block w-full p-2.5"
                    placeholder="price">
                </div>
                <div class="w-full ml-4">
                  <input v-model="count" type="number" 
                    class="border border-gray-400 text-gray-900 text-sm rounded-lg focus:border-indigo-600 block w-full p-2.5"
                    placeholder="count">
                </div>
              </div>
            </div>
          </div>
          <div class="mb-6">
            <div class="flex justify-center w-full">
              <div class="flex items-center justify-center w-5/6">
                <div align="center" class="w-4/6">
                  <button class="flex justify-center items-center rounded-[12px] bg-[#9b45b2] w-5/6 h-[36px]">
                    <span class="text-[#fff] text-base">Save</span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>

import StockDropdown from './core/StockDropdown.vue';

export default {
  name: 'AddStock',
  components: {
    StockDropdown
  },
  data() {
    return {
      count: 1,
      price: 0,
      selectedItemIds: [],
      selectedItem: null
    }
  },
  methods: {
    close() {
      this.$emit('close');
    },
    selectedData(value) {
      this.selectedItem = value
    },
  },
};
</script>
<style>
.custom-text {
  position: relative;
  display: block;
  color: #002b6d;
  width: 80%;
  border: none;
  border-bottom: 1px solid #d8e7fb;
}

.modal {
  box-shadow: 0 3px 20px 0 rgba(59, 64, 87, 0.22);
  border-radius: 6px;
}
</style>