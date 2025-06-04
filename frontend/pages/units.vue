<script setup lang="ts">
import type { Unit } from '~/types/unit'

definePageMeta({
  middleware: 'auth',
})

const toast = useToast()

const { data, refresh } = await useFetch<{ units: Unit[] }>(
  createUrl({ url: API_URLS.units.all }),
  {
    onResponseError: (error) => {
      toast.add({
        title: `Ошибка при получении подразделений: ${error}`,
        color: 'error',
      })
    },
  },
)
</script>

<template>
  <div class="container mx-auto p-6 max-w-200">
    <h1 class="text-2xl font-bold mb-6">
      Подразделения
    </h1>
    <div class="grid grid-cols-1">
      <div class="bg-white rounded-lg shadow p-4">
        <div
          v-if="Boolean(data?.units.length)"
          class="space-y-3"
        >
          <UnitsListItem
            v-for="unit in data?.units"
            :key="unit.id"
            :unit="unit"
            @update="refresh"
          />
        </div>
        <div
          v-else
          class="text-gray-500 py-4 text-center"
        >
          Нет данных
        </div>
      </div>
    </div>
  </div>
</template>
