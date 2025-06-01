<script setup lang="ts">
import type { Application, ApplicationStatus } from '~/types/application'

const props = defineProps<{
  title: string
  applications?: Application[]
  applicationsUpdate: () => void
}>()

const toast = useToast()

const changeStatus = async (id: string, status: ApplicationStatus) => {
  try {
    await $fetch(createUrl({ url: API_URLS.changeStatus }), {
      method: 'PATCH',
      body: {
        id,
        status,
      },
    })
    toast.add({ title: 'Статус заявки изменен', color: 'success' })
    props.applicationsUpdate()
  }
  catch {
    toast.add({ title: 'Ошибка при изменении статуса заявки', color: 'error' })
  }
}

const handleDelete = async (id: string) => {
  try {
    await $fetch(createUrl({ url: API_URLS.delete }), {
      method: 'DELETE',
      body: { id },
    })
    toast.add({ title: 'Заявка удалена', color: 'success' })
    props.applicationsUpdate()
  }
  catch {
    toast.add({ title: 'Ошибка при удалении заявки', color: 'error' })
  }
}

const handleMarkDone = async (id: string) => {
  await changeStatus(id, 'done')
}

const handleMarkUndone = async (id: string) => {
  await changeStatus(id, 'pending')
}
</script>

<template>
  <div class="bg-white rounded-lg shadow p-4">
    <h2 class="text-xl font-semibold text-gray-800 mb-4 pb-2 border-b">
      {{ title }}
    </h2>
    <div
      v-if="Boolean(applications?.length)"
      class="space-y-3"
    >
      <ApplicationsListItem
        v-for="application in applications"
        :key="application.id"
        :application="application"
        @mark-done="handleMarkDone"
        @mark-undone="handleMarkUndone"
        @delete="handleDelete"
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
