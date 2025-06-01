<script setup lang="ts">
import type { DropdownMenuItem } from '@nuxt/ui'
import type { Application } from '~/types/application'

const props = defineProps<{
  application: Application
}>()

const emit = defineEmits<{
  'mark-done': [id: string]
  'mark-undone': [id: string]
  'delete': [id: string]
}>()

const menuOpen = ref(false)

const handleMarkUndone = () => {
  emit('mark-undone', props.application.id)
  menuOpen.value = false
}

const handleMarkDone = () => {
  emit('mark-done', props.application.id)
  menuOpen.value = false
}

const handleDelete = () => {
  emit('delete', props.application.id)
  menuOpen.value = false
}

const menuItems: DropdownMenuItem[] = [
  props.application.status === 'done'
    ? {
        label: 'Отметить как невыполненное',
        icon: 'i-heroicons-x-circle',
        onSelect: handleMarkUndone,
      }
    : {
        label: 'Отметить как выполненное',
        icon: 'i-heroicons-check-circle',
        onSelect: handleMarkDone,
      },
  {
    label: 'Удалить',
    icon: 'i-heroicons-trash',
    color: 'error',
    onSelect: handleDelete,
  },
]
</script>

<template>
  <div class="relative">
    <div
      :class="{
        'border-l-4 pl-3 py-2 bg-gray-50 cursor-pointer rounded hover:bg-gray-100 flex justify-between items-start': true,
        'border-yellow-400': application.status === 'pending',
        'border-green-400': application.status === 'done',
      }"
    >
      <div>
        <h3 class="font-medium">
          {{ application.text }}
        </h3>
        <p class="text-sm text-gray-600">
          {{ application.unit?.name }}
          Отправил: {{ application.user_tg?.name }}
        </p>
      </div>

      <UDropdownMenu
        v-model:open="menuOpen"
        :items="menuItems"
        :content="{ align: 'end', side: 'bottom', sideOffset: 8 }"
      >
        <button
          class="p-1 rounded-full hover:bg-gray-200 transition-colors duration-200"
        >
          <span class="sr-only">Открыть меню</span>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 6.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 12.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 18.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5Z"
            />
          </svg>
        </button>
      </UDropdownMenu>
    </div>
  </div>
</template>
