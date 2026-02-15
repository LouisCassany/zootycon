<template>
  <div class="text-3xl underline">
    Hello {{ world }}
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const world = ref("world!")

</script>