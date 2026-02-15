<template>
  <div
    class="m-4 w-[297mm] h-[210mm] bg-slate-50 text-slate-900 p-6 mx-auto shadow-2xl print:shadow-none print:m-0 font-sans border flex flex-col gap-4 overflow-hidden relative">

    <div class="flex-grow flex gap-6">
      <div class="w-1/2 flex flex-col gap-4">
        <!-- TRACKS -->
        <div class="grid grid-cols-1 gap-3">
          <!-- Coins -->
          <div class="bg-amber-50 border-2 border-amber-500 p-2 rounded-lg">
            <div class="flex items-center gap-2 mb-2">
              <span class="bg-amber-500 text-white p-1 rounded-full text-xs">💰</span>
              <span class="text-xs font-black uppercase text-amber-800 tracking-tight">Coin Warehouse (Limit 20)</span>
            </div>
            <div class="flex flex-wrap gap-1.5">
              <div v-for="n in 20" :key="'c' + n"
                class="w-6 h-6 border-2 border-amber-400 bg-white rounded-full flex items-center justify-center text-[10px] font-bold text-amber-700">
                {{ n }}
              </div>
            </div>
          </div>

          <!-- Appeal -->
          <div class="bg-purple-50 border-2 border-purple-500 p-2 rounded-lg">
            <div class="flex items-center gap-2 mb-2">
              <span class="bg-purple-500 text-white p-1 rounded-full text-xs">🏆</span>
              <span class="text-xs font-black uppercase text-purple-800 tracking-tight">Appeal Track (Victory
                Points)</span>
            </div>
            <div class="flex flex-wrap gap-1">
              <div v-for="n in 50" :key="'a' + n"
                class="w-5 h-5 border border-purple-300 bg-white flex items-center justify-center text-[9px] font-bold text-purple-700">
                {{ n }}
              </div>
            </div>
          </div>
        </div>

        <!-- REGISTRY -->
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const zones = ref([
  {
    name: "Savannah Zone",
    animals: [
      { code: "Z", name: "Zebra", size: 4, reqs: "2🌾", vp: 5 },
      { code: "G", name: "Giraffe", size: 5, reqs: "2🌾 1🌲", vp: 7 },
      { code: "L", name: "Lion", size: 6, reqs: "2🪨 1💧", vp: 8 },
      { code: "E", name: "Elephant", size: 8, reqs: "3💧 2🌲", vp: 12 },
    ]
  },
  {
    name: "Wetland Zone",
    animals: [
      { code: "F", name: "Flamingo", size: 3, reqs: "1💧", vp: 4 },
      { code: "C", name: "Crocodile", size: 5, reqs: "3💧 1🪨", vp: 7 },
      { code: "D", name: "Dolphin", size: 7, reqs: "5💧", vp: 10 },
    ]
  },
  {
    name: "Alpine & Exotic",
    animals: [
      { code: "Penguin", name: "Penguin", size: 2, reqs: "1🪨 1💧", vp: 4 }, // Changed code to Penguin to match icon filename if N was just a placeholder
      { code: "M", name: "Monkey", size: 3, reqs: "2🌲", vp: 5 },
      { code: "P", name: "Panda", size: 4, reqs: "3🌲", vp: 8 },
    ]
  }
])
</script>

<style scoped>
@media print {
  @page {
    size: A4 landscape;
    margin: 0;
  }

  body {
    background: white;
  }
}

/* Ensure the grid keeps a consistent square ratio */
.grid-cols-10 {
  grid-template-columns: repeat(10, minmax(0, 1fr));
}

.grid-rows-10 {
  grid-template-rows: repeat(10, minmax(0, 1fr));
}
</style>