<script setup lang="ts">
import type { Application } from '~/types/application'

defineProps<{
  application: Application
}>()
</script>

<template>
  <div
    :class="{
      'border-l-4 pl-3 py-2 bg-gray-50 cursor-pointer rounded hover:bg-gray-100': true,
      'border-yellow-400': application.status === 'pending',
      'border-green-400': application.status === 'done',
    }"
  >
    <h3 class="font-medium">
      {{ application.text }}
    </h3>
    <p class="text-sm text-gray-600">
      {{ application.unit.name }}
    </p>
  </div>
</template>
