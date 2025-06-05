<script setup lang="ts">
import type { DropdownMenuItem } from '@nuxt/ui'
import type { User } from '~/types/user'

const isEditing = ref(false)
const menuOpen = ref(false)
const isSaving = ref(false)

const toast = useToast()

const authStore = useAuthStore()

const isTheSameUser = computed(() => props.user.id === authStore.user?.id)

const props = defineProps<{
  user: User
}>()

const emit = defineEmits<{
  (e: 'update'): void
}>()

const { meta, defineField, resetForm, handleSubmit } = useForm({
  initialValues: {
    role: props.user.role,
  },
})

const [role, roleAttrs] = defineField('role')

const menuItems: DropdownMenuItem[] = [
  {
    label: 'Редактировать',
    icon: 'i-heroicons-pencil',
    color: 'primary',
    onSelect: () => {
      isEditing.value = !isEditing.value
    },
  },
]

const onSubmit = handleSubmit(async (values) => {
  isSaving.value = true
  try {
    await $fetch<{ user: User }>(createUrl({ url: API_URLS.users.changeRole }), {
      method: 'PATCH',
      body: {
        id: props.user.id,
        role: values.role,
      },
      headers: getAuthHeaders(),
    })
    toast.add({ title: 'Роль пользователя успешно изменена', color: 'success' })
    isEditing.value = false
    emit('update')
  }
  catch (e) {
    toast.add({ title: `Ошибка при сохранении: ${e}`, color: 'error' })
  }
  finally {
    isSaving.value = false
  }
})
</script>

<template>
  <div class="relative">
    <div
      class="border-l-4 pl-3 py-2 bg-gray-50 cursor-pointer rounded hover:bg-gray-100 flex justify-between items-start border-blue-400"
    >
      <div>
        <h3
          class="font-medium"
        >
          {{ user.name }} {{ user.surname }}
        </h3>
        <p class="text-sm text-gray-500">
          {{ user.email }}
        </p>
        <UBadge
          v-if="!isEditing"
          :label="user.role"
          :color="user.role === 'admin' ? 'primary' : 'secondary'"
        />
        <form
          v-else
          class="flex items-center gap-2"
          @submit="onSubmit"
        >
          <USelect
            v-model="role"
            :items="['admin', 'user']"
            v-bind="roleAttrs"
            class="w-48"
          />
          <UButton
            color="primary"
            size="sm"
            type="submit"
            :disabled="!meta.valid || role === user.role"
          >
            Сохранить
          </UButton>

          <UButton
            color="neutral"
            size="sm"
            type="button"
            @click="isEditing = false; resetForm()"
          >
            Отменить
          </UButton>
        </form>
      </div>

      <UDropdownMenu
        v-if="!isTheSameUser && authStore.isAdmin"
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
