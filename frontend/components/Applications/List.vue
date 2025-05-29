<script setup lang="ts">
import type { Application } from '~/types/application'

defineProps<{
  title: string
  applications: Application[]
}>()
</script>

<template>
  <div class="bg-white rounded-lg shadow p-4">
    <h2 class="text-xl font-semibold text-gray-800 mb-4 pb-2 border-b">
      {{ title }}
    </h2>
    <div
      v-if="applications.length > 0"
      class="space-y-3"
    >
      <ApplicationsListItem
        v-for="application in applications"
        :key="application.id"
        :application="application"
      />
    </div>
    <div
      v-else
      class="text-gray-500 py-4 text-center"
    >
      Нет данных
    </div>
  </div>
</template>
