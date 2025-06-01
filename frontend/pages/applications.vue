<script setup lang="ts">
import type { Application } from '~/types/application'

definePageMeta({
  middleware: 'auth',
})

const toast = useToast()

const { data: applications, refresh } = await useFetch<{ applications: Application[] }>(
  createUrl({ url: '/applications' }),
  {
    onResponseError: (error) => {
      toast.add({
        title: `Ошибка при получении заявок: ${error}`,
        color: 'error',
      })
    },
  },
)

const pendingApplications = computed(() => applications.value?.applications.filter(app => app.status === 'pending'))
const doneApplications = computed(() => applications.value?.applications.filter(app => app.status === 'done'))
</script>

<template>
  <div class="container mx-auto p-6">
    <h1 class="text-2xl font-bold mb-6">
      Заявки
    </h1>
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
