<template>
  <!-- Main Sheet: A4 Landscape -->
  <div
    class="w-[297mm] h-[210mm] p-6 flex flex-col gap-4 m-4 print:m-0 font-sans bg-emerald-50 shadow-2xl rounded-3xl overflow-hidden border-8 border-white">

    <!-- SECTION HAUTE -->
    <div class="flex justify-between gap-6 h-[135mm]">

      <!-- COLONNE GAUCHE: Registre -->
      <div class="w-[10.5cm] flex flex-col">
        <h2
          class="bg-indigo-800 text-white text-sm font-black px-4 py-2 uppercase tracking-tighter italic rounded-t-xl shadow-md flex items-center gap-2">
          <span class="text-lg">🐾</span> Registre des Empreintes
        </h2>
        <!-- Added overflow-hidden here to fix the bottom corner border issue -->
        <div class="border-2 border-indigo-900 rounded-b-xl bg-white shadow-lg overflow-hidden">
          <table class="w-full text-left text-[11px] border-collapse">
            <thead class="bg-indigo-50 border-b-2 border-indigo-200 text-indigo-900">
              <tr>
                <th class="px-3 py-2 font-black uppercase">Animal</th>
                <th class="px-1 text-center font-black">Taille</th>
                <th class="px-1 font-black text-center">Besoins</th>
                <th class="px-1 text-center font-black bg-rose-600 text-white w-10">PV</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(a, idx) in animals" :key="a.code"
                :class="['border-b border-indigo-100 last:border-0', idx % 2 === 0 ? 'bg-white' : 'bg-indigo-50/30']">
                <td class="px-3 py-1.5 font-bold text-indigo-950 flex items-center gap-2">
                  <span
                    class="w-5 h-5 bg-indigo-900 text-white flex items-center justify-center rounded-md text-[9px] font-black italic">{{
                      a.code }}</span>
                  {{ a.name }}
                </td>
                <td class="text-center font-black text-indigo-600 text-xs">{{ a.size }}</td>
                <td class="px-1 text-center">
                  <div class="flex gap-0.5 justify-center">
                    <GameSymbol v-for="(req, i) in a.reqs" :key="i" :name="req" size="w-6 h-6" class="drop-shadow-sm" />
                  </div>
                </td>
                <td class="text-center font-black text-sm text-rose-600 bg-rose-50 border-l border-rose-100">{{
                  a.vp
                  }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- COLONNE DROITE : Grille de Zonage -->
      <div class="w-[12cm] flex flex-col">
        <h2
          class="bg-emerald-700 text-white text-sm font-black px-4 py-2 uppercase tracking-tighter italic rounded-t-xl shadow-md flex items-center gap-2">
          <span class="text-lg">🗺️</span> Carte du Sanctuaire
        </h2>

        <div
          class="grid grid-cols-10 grid-rows-10 border-4 border-emerald-900 bg-white h-full aspect-square shadow-inner rounded-b-xl">
          <div v-for="i in 100" :key="i" class="border border-emerald-100/50 relative flex items-center justify-center">
          </div>
        </div>
      </div>
    </div>

    <!-- SECTION BASSE: Trackers -->
    <div class="flex-1 flex flex-col">
      <div
        class="border-2 border-slate-900 rounded-2xl bg-white shadow-lg p-3 px-6 flex flex-col justify-around flex-1">

        <!-- Track 1: Gold -->
        <div class="flex items-center gap-4">
          <div class="w-24 text-right">
            <span class="text-[11px] font-black uppercase text-amber-700 block leading-none">Trésorerie</span>
            <span class="text-[8px] font-bold text-amber-400 uppercase tracking-widest italic">Pièces d'or</span>
          </div>
          <div class="flex flex-1 border-2 border-amber-500 rounded-lg bg-white overflow-hidden h-9 shadow-inner">
            <div v-for="n in 20" :key="n"
              :class="['flex-1 border-r border-amber-100 flex items-center justify-center text-xs font-black last:border-r-0', n % 5 === 0 ? 'bg-amber-500 text-white' : 'text-amber-900']">
              {{ n }}
            </div>
          </div>
        </div>

        <!-- Track 2: Reputation -->
        <div class="flex items-center gap-4">
          <div class="w-24 text-right">
            <span class="text-[11px] font-black uppercase text-sky-700 block leading-none">Réputation</span>
            <span class="text-[8px] font-bold text-sky-400 uppercase tracking-widest italic">Attrait</span>
          </div>
          <div class="flex flex-1 flex-col border-2 border-sky-500 rounded-lg bg-white overflow-hidden shadow-inner">
            <div class="flex border-b border-sky-100">
              <div v-for="n in 25" :key="n"
                :class="['flex-1 h-5 border-r border-sky-50 flex items-center justify-center text-[10px] font-black last:border-r-0', n % 5 === 0 ? 'bg-sky-600 text-white' : 'text-sky-900']">
                {{ n }}
              </div>
            </div>
            <div class="flex">
              <div v-for="n in 25" :key="n + 25"
                :class="['flex-1 h-5 border-r border-sky-50 flex items-center justify-center text-[10px] font-black last:border-r-0', (n + 25) % 5 === 0 ? 'bg-sky-600 text-white' : 'text-sky-900']">
                {{ n + 25 }}
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
  { name: "Flamant", code: "F", size: 3, reqs: ["water"], vp: 4 },
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