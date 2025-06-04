<script setup lang="ts">
import type { User } from '~/types/user'

const authStore = useAuthStore()

defineProps<{
  user: User
}>()
</script>

<template>
  <div class="relative">
    <div
      class="border-l-4 pl-3 py-2 bg-gray-50 cursor-pointer rounded hover:bg-gray-100 flex justify-between items-start border-blue-400"
    >
      <div>
        <h3
          v-if="user.id !== authStore.user?.id"
          class="font-medium"
        >
          {{ user.name }} {{ user.surname }}
        </h3>
        <UBadge
          v-else
          color="neutral"
        >
          {{ user.name }} {{ user.surname }}
        </UBadge>
        <p class="text-sm text-gray-500">
          {{ user.email }}
          <UBadge
            :label="user.role"
            :color="user.role === 'admin' ? 'primary' : 'secondary'"
          />
        </p>
      </div>
    </div>
  </div>
</template>
