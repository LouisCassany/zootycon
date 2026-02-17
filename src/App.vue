<template>
  <div class="w-[297mm] h-[210mm] m-4 border bg-[url(/background.png)] bg-center bg-cover flex">
    <!-- Animals -->
    <div></div>

    <!-- Grid -->
    <div class="grid grid-cols-10 auto-rows-[1cm] w-[10cm] h-[10cm] m-1 border">
      <div v-for="i in 100" :key="i" class="border"></div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const animals = [
  {
    name: "Savannah",
    icon: "🦁",
    list: [
      { name: "Zebra", code: "Z", size: 4, reqs: ["🌾", "🌾"], vp: 5 },
      { name: "Giraffe", code: "G", size: 5, reqs: ["🌾", "🌾", "🌲"], vp: 7 },
      { name: "Lion", code: "L", size: 6, reqs: ["🪨", "🪨", "💧"], vp: 8 },
      { name: "Elephant", code: "E", size: 8, reqs: ["💧", "💧", "🌲", "🌲"], vp: 12 }, // Condensed 3 water to 2 for space in UI, or keep logic
    ]
  },
  {
    name: "Wetland",
    icon: "🐊",
    list: [
      { name: "Flamingo", code: "F", size: 3, reqs: ["💧"], vp: 4 },
      { name: "Crocodile", code: "C", size: 5, reqs: ["💧", "💧", "🪨"], vp: 7 },
      { name: "Dolphin", code: "D", size: 7, reqs: ["💧", "💧", "💧"], vp: 10 },
    ]
  },
  {
    name: "Alpine & Exotic",
    icon: "🐼",
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
  background-color: black;
}
</style>