<script setup lang="ts">
import type { Application, ApplicationStatus } from '~/types/application'

const props = defineProps<{
  title: string
  applications: Application[]
}>()

const emit = defineEmits<{
  'update-applications': [applications: Application[]]
}>()

// Handle marking an application as done
const handleMarkDone = (id: string) => {
  const updatedApplications = props.applications.map((app) => {
    if (app.id === id) {
      return { ...app, status: 'done' as ApplicationStatus }
    }
    return app
  })
  emit('update-applications', updatedApplications)
}

// Handle deleting an application
const handleDelete = (id: string) => {
  const updatedApplications = props.applications.filter(app => app.id !== id)
  emit('update-applications', updatedApplications)
}
</script>

<template>
  <div class="bg-white rounded-lg shadow p-4">
    <h2 class="text-xl font-semibold text-gray-800 mb-4 pb-2 border-b">
      {{ title }}
    </h2>
    <div
      v-if="applications?.length > 0"
      class="space-y-3"
    >
      <ApplicationsListItem
        v-for="application in applications"
        :key="application.id"
        :application="application"
        @mark-done="handleMarkDone"
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
