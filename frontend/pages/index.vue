<script setup lang="ts">
import type { Application } from '~/types/application'
import type { Unit } from '~/types/unit'

definePageMeta({
  middleware: 'auth',
})

const toast = useToast()
const selectedUnitName = ref<string | undefined>()

const { data: applications, refresh } = await useFetch<{ applications: Application[] }>(
  createUrl({ url: API_URLS.applications }),
  {
    onResponseError: (error) => {
      toast.add({
        title: `Ошибка при получении заявок: ${error}`,
        color: 'error',
      })
    },
  },
)

const { data: unitsData } = await useFetch<{ units: Unit[] }>(
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

const unitsItems = computed<string[]>(() =>
  unitsData.value?.units?.map(unit => unit.name) ?? [],
)

const filteredApplications = computed(() => {
  if (!selectedUnitName.value) {
    return applications.value?.applications ?? []
  }
  return applications.value?.applications?.filter(app => app.unit.name === selectedUnitName.value) ?? []
})

const clearFilter = () => {
  selectedUnitName.value = undefined
}

const pendingApplications = computed(() => filteredApplications.value.filter(app => app.status === 'pending'))
const doneApplications = computed(() => filteredApplications.value.filter(app => app.status === 'done'))
</script>

<template>
  <div class="container mx-auto p-6">
    <h1 class="text-2xl font-bold mb-6">
      Заявки
    </h1>
    <div class="mb-6 flex flex-wrap items-center gap-4">
      <UFormField
        label="Фильтр по подразделению"
      >
        <USelectMenu
          v-model="selectedUnitName"
          :items="unitsItems"
          placeholder="Выберите подразделение"
          class="min-w-60 mr-4"
          clearable
        />
        <UButton
          v-if="selectedUnitName"
          color="secondary"
          size="sm"
          @click="clearFilter"
        >
          Сбросить фильтр
        </UButton>
      </UFormField>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <ApplicationsList
        title="Невыполненные заявки"
        :applications="pendingApplications"
        :applications-update="refresh"
      />
      <ApplicationsList
        title="Выполненные заявки"
        :applications="doneApplications"
        :applications-update="refresh"
      />
    </div>
  </div>
</template>
