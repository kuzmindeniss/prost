<script setup lang="ts">
import type { User } from '~/types/user'

definePageMeta({
  middleware: 'auth',
})

const toast = useToast()

const { data, refresh } = await useFetch<{ users: User[] }>(
  createUrl({ url: API_URLS.users.all }),
  {
    headers: getAuthHeaders(),
    onResponseError: (error) => {
      toast.add({
        title: `Ошибка при получении пользователей: ${error}`,
        color: 'error',
      })
    },
  },
)
</script>

<template>
  <div class="container mx-auto p-6 max-w-200">
    <div class="flex flex-col md:flex-row items-center md:justify-between md:items-baseline mb-6">
      <h1 class="text-2xl text-center md:text-left font-bold mb-6">
        Пользователи
      </h1>
    </div>
    <div class="grid grid-cols-1">
      <div class="bg-white rounded-lg shadow p-4">
        <div
          v-if="Boolean(data?.users?.length)"
          class="space-y-3"
        >
          <UsersListItem
            v-for="user in data?.users"
            :key="user.id"
            :user="user"
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
