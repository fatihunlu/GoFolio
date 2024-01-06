<template>
    <div class="w-full">
        <div class="mt-1 flex rounded-md shadow-sm">
            <input v-model="search" type="text"
                class="w-full px-3 py-2 border border-gray-400 rounded-lg outline-none focus:shadow-outline"
                :class="selectedItem ? 'pl-10' : ''" placeholder="search stock code" />
            <svg v-if="selectedItem" width="17px" height="17px"
                class="absolute ml-[8px] mt-[12px] cursor-pointer hover:bg-opacity-75" @click="clear" viewBox="0 0 24 24"
                fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M20 20L4 4.00003M20 4L4.00002 20" stroke="#9b45b2" stroke-width="2" stroke-linecap="round" />
            </svg>
        </div>
        <aside class="absolute z-20 flex flex-col items-start bg-white border rounded-md shadow-md mt-1 w-5/6" role="menu"
            v-if="searchItemList.length > 0 && showSearchItems == true">
            <ul class="flex flex-col w-full">
                <li class="px-2 py-3 space-x-2 hover:bg-blue-600 hover:text-white focus:bg-blue-600 focus:text-white focus:outline-none "
                    v-for="(item, index) in searchItemList" :key="index"
                    @click="selectSearchItem(item); showSearchItems = false;">{{ item.symbol }}</li>
            </ul>
        </aside>
    </div>
</template>
<script>

export default {
    name: "StockDropdown",
    data() {
        return {
            search: "",
            selectedItem: null,
            showSearchItems: false,
            isMouseOverList: false,
            searchItemList: [],
            ignoredList: []
        };
    },
    watch: {
        search: function (newVal, oldVal) {
            if (newVal == null || newVal == '' || this.selectedItem != null) {
                return;
            }

            // TODO
            this.$axios.get(`http://localhost:1323/search?q=${this.search}`)
                .then(response => {
                    this.searchItemList = response.data.quotes;
                    this.showSearchItems = true;
                })
                .catch(err => {
                    console.error(err)
                    return [];
                })

        },
    },
    methods: {
        selectSearchItem(item) {
            this.search = item.symbol;
            this.selectedItem = item;
            this.showSearchItems = false;
            this.$emit('selected', item)
        },

        clear() {
            this.search = '';
            this.selectedItem = null;
        }

    },

    created() {
    },

};
</script>
<style></style>