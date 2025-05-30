<script setup lang="ts">
import type { Application } from '~/types/application'

definePageMeta({
  middleware: 'auth',
})

const toast = useToast()

const { data: applications } = await useFetch<{ applications: Application[] }>(
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

console.log(applications.value?.applications)

const pendingApplications = computed(() => applications.value?.applications.filter(app => app.status === 'pending'))
const doneApplications = computed(() => applications.value?.applications.filter(app => app.status === 'done'))

// Update pending applications
const updatePendingApplications = (applications: Application[]) => {
  // Applications with 'done' status should be moved to the doneApplications list
  const newPending = applications.filter(app => app.status === 'pending')
  const newDone = applications.filter(app => app.status === 'done')
  pendingApplications.value = newPending
  if (newDone.length > 0) {
    doneApplications.value = [...doneApplications.value, ...newDone]
  }
}

// Update done applications
const updateDoneApplications = (applications: Application[]) => {
  doneApplications.value = applications
}
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
        @update-applications="updatePendingApplications"
      />
      <ApplicationsList
        title="Выполненные заявки"
        :applications="doneApplications"
        @update-applications="updateDoneApplications"
      />
    </div>
  </div>
</template>
