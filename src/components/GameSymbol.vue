<template>
    <img :src="`/symbols/${name}.png`" :class="[size, 'object-contain inline-block']" :alt="name" />
</template>

<script setup lang="ts">
defineProps({
    name: {
        type: String,
        required: true
    },
    size: {
        type: String,
        default: 'w-6 h-6' // Default doubled size (24px)
    }
})
</script>