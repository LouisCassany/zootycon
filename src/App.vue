<template>
  <!-- Main Print Sheet (A4 Landscape) -->
  <div class="w-[297mm] h-[210mm] bg-slate-50 p-6 flex flex-col gap-4 relative m-4 print:m-0">

    <!-- Blueprint Header -->
    <div class=" flex justify-between items-end border-b-2 border-slate-300 pb-2">
      <div>
        <h1 class="text-4xl font-black tracking-tighter text-slate-900">CONSERVATORY BLUEPRINT</h1>
        <p class="text-xs font-bold uppercase tracking-widest text-slate-500">Zoo Director Master Plan // Version 2.0
        </p>
      </div>
      <div class="flex gap-4">
        <div class="border-2 border-slate-800 px-4 py-1">
          <span class="text-[10px] block font-bold leading-none">DIRECTOR NAME</span>
          <div class="h-6 w-32"></div>
        </div>
        <div class="border-2 border-slate-800 px-4 py-1 bg-slate-800 text-white">
          <span class="text-[10px] block font-bold leading-none">PROFILE</span>
          <div class="h-6 w-32"></div>
        </div>
      </div>
    </div>

    <div class="flex flex-1 gap-6">
      <!-- LEFT COLUMN: Animal Registry -->
      <div class="w-1/3 flex flex-col gap-3">
        <h2 class="bg-slate-800 text-white text-xs font-bold px-2 py-1">PAW-PRINT REGISTRY (POPULATE)</h2>

        <div v-for="category in animals" :key="category.name" class="border border-slate-300 rounded overflow-hidden">
          <div class="bg-slate-200 px-2 py-0.5 flex justify-between items-center">
            <span class="text-xs font-bold uppercase">{{ category.icon }} {{ category.name }}</span>
          </div>

          <table class="w-full text-left text-[10px]">
            <thead class="bg-slate-50 border-b border-slate-200 text-slate-500">
              <tr>
                <th class="px-2 py-1">NAME</th>
                <th class="px-1 text-center">SQ</th>
                <th class="px-1">NEEDS</th>
                <th class="px-1 text-center font-bold text-slate-800">VP</th>
                <th class="px-1 text-center">DONE</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="a in category.list" :key="a.code"
                class="border-b border-slate-100 last:border-0 hover:bg-white">
                <td class="px-2 py-1 font-bold">({{ a.code }}) {{ a.name }}</td>
                <td class="text-center font-mono">{{ a.size }}</td>
                <td class="text-[12px] leading-none">{{ a.reqs.join('') }}</td>
                <td class="text-center font-black">{{ a.vp }}</td>
                <td class="text-center">
                  <div class="w-4 h-4 border border-slate-400 rounded-full mx-auto bg-white/50"></div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Director Profile Reference (Small) -->
        <div class="mt-auto border-t-2 border-dashed border-slate-300 pt-2">
          <h3 class="text-[9px] font-bold text-slate-400 mb-1 uppercase">Director Quick-Ref</h3>
          <div class="grid grid-cols-2 gap-1 text-[8px] leading-tight text-slate-500">
            <div><strong>Architect:</strong> +2 Fences (L)</div>
            <div><strong>Botanist:</strong> +1 Terrain (L)</div>
            <div><strong>Vet:</strong> Ignore 1 Symbol</div>
            <div><strong>Marketer:</strong> Kiosk = 1 Appeal</div>
          </div>
        </div>
      </div>

      <!-- CENTER COLUMN: Zoning Grid -->
      <div class="flex-1 flex flex-col items-center">
        <div class="mb-2 flex justify-between w-full items-center">
          <span class="text-[10px] font-bold text-slate-400 uppercase">Zoning Grid (10x10)</span>
          <div class="flex gap-2">
            <span class="text-[10px] px-2 py-0.5 bg-slate-100 border border-slate-300">🌲 Tree</span>
            <span class="text-[10px] px-2 py-0.5 bg-slate-100 border border-slate-300">🪨 Rock</span>
            <span class="text-[10px] px-2 py-0.5 bg-slate-100 border border-slate-300">💧 Water</span>
            <span class="text-[10px] px-2 py-0.5 bg-slate-100 border border-slate-300">🌾 Grass</span>
          </div>
        </div>

        <!-- The Dot Grid -->
        <div class="grid grid-cols-10 grid-rows-10 border">
          <div v-for="i in 100" :key="i"
            class="w-10 h-10 border border-black relative group flex items-center justify-center">
          </div>
        </div>

        <!-- Scoring Legend -->
        <div class="mt-4 w-full grid grid-cols-3 gap-2 text-[9px] border border-slate-200 p-2 bg-slate-50 rounded">
          <div><strong>(+) GAIN:</strong> Animal VP + 2 per Kiosk + 1 per 3 Coins</div>
          <div class="border-x border-slate-200 px-2 text-red-700"><strong>(-) LOSS:</strong> -1 per Terrain outside
            cage</div>
          <div class="text-red-700"><strong>(-) LOSS:</strong> -1 per Orphan (open) Fence</div>
        </div>
      </div>

      <!-- RIGHT COLUMN: Warehouse & Tracks -->
      <div class="w-16 flex flex-col gap-4 items-center">
        <!-- Coin Track -->
        <div class="flex-1 flex flex-col items-center">
          <span
            class="text-[9px] font-black uppercase text-amber-600 rotate-90 h-12 w-12 flex items-center justify-center">Coins</span>
          <div class="flex flex-col-reverse border-2 border-amber-500 rounded bg-white overflow-hidden">
            <div v-for="n in 20" :key="n"
              class="w-10 h-6 border-b border-amber-100 flex items-center justify-center text-[10px] font-bold">
              {{ n }}
            </div>
          </div>
        </div>
      </div>

      <div class="w-12 flex flex-col gap-4 items-center">
        <!-- Appeal Track -->
        <div class="flex-1 flex flex-col items-center">
          <span
            class="text-[9px] font-black uppercase text-blue-600 rotate-90 h-12 w-12 flex items-center justify-center">Appeal</span>
          <div class="flex flex-col-reverse border-2 border-blue-500 rounded bg-white overflow-hidden">
            <div v-for="n in 50" :key="n"
              :class="['w-8 h-[9.3px] border-b border-blue-100 flex items-center justify-center text-[7px] font-bold', n % 5 === 0 ? 'bg-blue-50' : '']">
              {{ n }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer Action Reference -->
    <div class="grid grid-cols-4 gap-4 mt-2">
      <div v-for="act in actions" :key="act.title" class="border border-slate-300 p-2 bg-white/50 rounded">
        <h4 class="text-[10px] font-black uppercase border-b border-slate-300 mb-1">{{ act.title }}</h4>
        <p class="text-[8px] leading-tight text-slate-600"><strong>Leader:</strong> {{ act.lead }}</p>
        <p class="text-[8px] leading-tight text-slate-600"><strong>Follow:</strong> {{ act.follow }}</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const actions = [
  { title: "1. Build", lead: "5 Fences + 1 Kiosk", follow: "3 Fences" },
  { title: "2. Landscape", lead: "3 Terrain Symbols", follow: "1 Terrain Symbol" },
  { title: "3. Populate", lead: "Place + Gain 3 Coins", follow: "Place (Cost 2 Coins)" },
  { title: "4. Sponsor", lead: "5 Coins + 1 Appeal", follow: "2 Coins" },
]

const animals = [
  {
    name: "Savannah",
    icon: "🌾",
    list: [
      { name: "Zebra", code: "Z", size: 4, reqs: ["🌾", "🌾"], vp: 5 },
      { name: "Giraffe", code: "G", size: 5, reqs: ["🌾", "🌾", "🌲"], vp: 7 },
      { name: "Lion", code: "L", size: 6, reqs: ["🪨", "🪨", "💧"], vp: 8 },
      { name: "Elephant", code: "E", size: 8, reqs: ["💧", "💧", "🌲", "🌲"], vp: 12 },
    ]
  },
  {
    name: "Wetland",
    icon: "💧",
    list: [
      { name: "Flamingo", code: "F", size: 3, reqs: ["💧"], vp: 4 },
      { name: "Crocodile", code: "C", size: 5, reqs: ["💧", "💧", "🪨"], vp: 7 },
      { name: "Dolphin", code: "D", size: 7, reqs: ["💧", "💧", "💧"], vp: 10 },
    ]
  },
  {
    name: "Alpine / Exotic",
    icon: "🏔️",
    list: [
      { name: "Penguin", code: "N", size: 2, reqs: ["🪨", "💧"], vp: 4 },
      { name: "Monkey", code: "M", size: 3, reqs: ["🌲", "🌲"], vp: 5 },
      { name: "Panda", code: "P", size: 4, reqs: ["🌲", "🌲", "🌲"], vp: 8 },
    ]
  }
]
</script>

<style>
body {
  background: black
}
</style>