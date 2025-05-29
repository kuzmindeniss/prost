<script setup lang="ts">
import type { Application } from '~/types/application'

definePageMeta({
  middleware: 'auth',
})

// Sample data - replace with your actual data source
const pendingApplications = ref<Application[]>([
  { id: '1', text: 'Frontend Developer', status: 'pending', unit: { id: '1', name: 'TechCorp' }, user: { id: '1', name: 'John Doe', tgUsername: 'john_doe', unit: { id: '1', name: 'TechCorp' } } },
  { id: '2', text: 'UX Designer', status: 'pending', unit: { id: '2', name: 'DesignHub' }, user: { id: '2', name: 'Jane Smith', tgUsername: 'jane_smith', unit: { id: '2', name: 'DesignHub' } } },
  { id: '3', text: 'Product Manager', status: 'pending', unit: { id: '3', name: 'InnovateSoft' }, user: { id: '3', name: 'Alice Johnson', tgUsername: 'alice_johnson', unit: { id: '3', name: 'InnovateSoft' } } },
])

const doneApplications = ref<Application[]>([
  { id: '4', text: 'Backend Developer', status: 'done', unit: { id: '4', name: 'ServerTech' }, user: { id: '4', name: 'Bob Brown', tgUsername: 'bob_brown', unit: { id: '4', name: 'ServerTech' } } },
  { id: '5', text: 'DevOps Engineer', status: 'done', unit: { id: '5', name: 'CloudSys' }, user: { id: '5', name: 'Charlie Davis', tgUsername: 'charlie_davis', unit: { id: '5', name: 'CloudSys' } } },
])
</script>

<template>
  <div class="container mx-auto p-6">
    <h1 class="text-2xl font-bold mb-6">
      Your Applications
    </h1>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <ApplicationsList
        title="Pending Applications"
        :applications="pendingApplications"
      />
      <ApplicationsList
        title="Done Applications"
        :applications="doneApplications"
      />
    </div>
  </div>
</template>
