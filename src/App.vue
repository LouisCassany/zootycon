<template>
  <!-- Main Sheet: A4 Landscape -->
  <div
    class="w-[297mm] h-[210mm] p-6 flex flex-row gap-4 m-4 print:m-0 font-sans bg-emerald-50 shadow-2xl overflow-hidden">

    <!-- COLONNE 1: Registre des Animaux -->
    <div class="w-[100mm] flex flex-col h-full">
      <h2
        class="bg-indigo-800 text-white text-xs font-black px-3 py-2 uppercase tracking-tighter italic rounded-t-xl shadow-md flex items-center gap-2">
        <span class="text-base">🐾</span> Registre
      </h2>
      <div class="border-2 border-indigo-900 rounded-b-xl bg-white shadow-lg overflow-hidden flex-1">
        <table class="w-full text-left text-[10px] border-collapse">
          <thead class="bg-indigo-50 border-b-2 border-indigo-200 text-indigo-900">
            <tr>
              <th class="px-2 py-2 font-black uppercase">Animal</th>
              <th class="px-1 text-center font-black">Taille</th>
              <th class="px-1 font-black text-center">Besoins</th>
              <th class="px-1 text-center font-black bg-rose-600 text-white w-8">PV</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(a, idx) in animals" :key="a.code"
              :class="['border-b border-indigo-100 last:border-0', idx % 2 === 0 ? 'bg-white' : 'bg-indigo-50/30']">
              <td class="px-2 py-2 font-bold text-indigo-950 flex items-center gap-1.5">
                <span
                  class="w-5 h-5 bg-indigo-900 text-white flex items-center justify-center rounded-md text-[9px] font-black italic shrink-0">{{
                    a.code }}</span>
                <span class="truncate">{{ a.name }}</span>
              </td>
              <td class="text-center font-black text-indigo-600">{{ a.size }}</td>
              <td class="px-1 text-center">
                <div class="flex gap-0.5 justify-center flex-wrap">
                  <GameSymbol v-for="(req, i) in a.reqs" :key="i" :name="req" size="w-8 h-8" class="drop-shadow-sm" />
                </div>
              </td>
              <td class="text-center font-black text-rose-600 bg-rose-50 border-l border-rose-100">{{ a.vp }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- COLONNE 2: Carte du Sanctuaire -->
    <div class="w-[125mm] flex flex-col h-full">
      <h2
        class="bg-emerald-700 text-white text-sm font-black px-4 py-2 uppercase tracking-tighter italic rounded-t-xl shadow-md flex items-center gap-2">
        <span class="text-lg">🗺️</span> Carte du Sanctuaire
      </h2>
      <div
        class="grid grid-cols-10 grid-rows-10 border-4 border-emerald-900 bg-white aspect-square shadow-inner rounded-b-xl relative overflow-hidden">
        <div v-for="i in 100" :key="i" class="border border-black/50 relative flex items-center justify-center">
        </div>
      </div>
    </div>

    <!-- COLONNE 3: Trackers Verticaux -->
    <div class="flex-1 flex flex-col h-full font-sans">
      <!-- Header -->
      <h2
        class="bg-slate-900 text-white text-xs font-black px-4 py-2.5 uppercase tracking-widest italic rounded-t-xl shadow-md flex items-center gap-2 border-b border-slate-700">
        ⚙️ Tableau de bord
      </h2>

      <!-- Main Board Area -->
      <div class="border-x-2 border-b-4 border-slate-900 rounded-b-xl bg-slate-50 p-4 flex flex-row gap-6 flex-1">

        <!-- Vertical Gold Track -->
        <div class="flex flex-col items-center flex-1">
          <div class="flex items-center gap-1 mb-2">
            <div class="w-2 h-2 rounded-full bg-amber-500 animate-pulse"></div>
            <span class="text-[10px] font-black text-amber-800 uppercase tracking-tighter">Trésor</span>
          </div>

          <div
            class="relative w-full h-full bg-amber-100 rounded-lg border-2 border-amber-900/20 shadow-inner flex flex-col overflow-hidden">
            <div v-for="n in 20" :key="n" :class="[
              'flex-1 flex items-center justify-center relative border-t border-amber-200/50 transition-colors',
              n % 5 === 0 ? 'bg-amber-400/30' : ''
            ]">
              <!-- Regular Number -->
              <span class="text-[10px] font-bold text-amber-900/60 tabular-nums">
                {{ n }}
              </span>
            </div>
          </div>
        </div>

        <!-- Vertical Reputation Track -->
        <div class="flex flex-col items-center flex-[2]">
          <div class="flex items-center gap-1 mb-2">
            <div class="w-2 h-2 rounded-full bg-sky-500 animate-pulse"></div>
            <span class="text-[10px] font-black text-sky-800 uppercase tracking-tighter">Réputation</span>
          </div>

          <div class="flex w-full gap-2 h-full">
            <!-- Two-column track for 1-50 -->
            <div v-for="col in [0, 25]" :key="col"
              class="flex flex-col flex-1 bg-sky-100 rounded-lg border-2 border-sky-900/20 shadow-inner overflow-hidden">

              <div v-for="n in 25" :key="n + col" :class="[
                'flex-1 flex items-center justify-center relative border-t border-sky-200/50',
                (n + col) % 5 === 0 ? 'bg-sky-500/20' : ''
              ]">
                <span class="text-[9px] font-bold text-sky-900/50 tabular-nums">
                  {{ n + col }}
                </span>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import GameSymbol from './components/GameSymbol.vue';

const animals = [
  { name: "Zèbre", code: "Z", size: 4, reqs: ["grass", "grass"], vp: 5 },
  { name: "Girafe", code: "G", size: 5, reqs: ["grass", "grass", "tree"], vp: 7 },
  { name: "Lion", code: "L", size: 6, reqs: ["rock", "rock", "water"], vp: 8 },
  { name: "Éléphant", code: "E", size: 8, reqs: ["water", "water", "water", "tree", "tree"], vp: 12 },
  { name: "Flamant rose", code: "F", size: 3, reqs: ["water"], vp: 4 },
  { name: "Crocodile", code: "C", size: 5, reqs: ["water", "water", "water", "rock"], vp: 7 },
  { name: "Dauphin", code: "D", size: 7, reqs: ["water", "water", "water", "water", "water"], vp: 10 },
  { name: "Pingouin", code: "N", size: 2, reqs: ["rock", "water"], vp: 4 },
  { name: "Singe", code: "M", size: 3, reqs: ["tree", "tree"], vp: 5 },
  { name: "Panda", code: "P", size: 4, reqs: ["tree", "tree", "tree"], vp: 8 },
]
</script>

<style>
body {
  background: #064e3b;
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  margin: 0;
}

@media print {
  body {
    background: white;
  }

  .shadow-2xl,
  .shadow-lg,
  .shadow-inner,
  .shadow-md {
    box-shadow: none !important;
  }

  .m-4 {
    margin: 0 !important;
  }
}
</style>