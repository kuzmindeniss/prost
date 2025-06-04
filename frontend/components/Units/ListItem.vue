<script setup lang="ts">
import type { DropdownMenuItem } from '@nuxt/ui'
import * as v from 'valibot'
import { toTypedSchema } from '@vee-validate/valibot'
import type { Unit } from '~/types/unit'

const props = defineProps<{
  unit: Unit
}>()

const emit = defineEmits<{
  (e: 'update'): void
}>()

const isSaving = ref(false)
const isEditing = ref(false)

const onDelete = async () => {
  try {
    await $fetch(createUrl({ url: API_URLS.units.delete }), {
      method: 'DELETE',
      body: {
        id: props.unit.id,
      },
    })
    emit('update')
  }
  catch (e: any) {
    toast.add({ title: e.data.error ?? 'Ошибка при удалении', color: 'error' })
  }
}

const { meta, defineField, resetForm, handleSubmit } = useForm({
  validationSchema: toTypedSchema(
    v.object({
      name: v.pipe(v.string(), v.minLength(1, 'Введите имя')),
    }),
  ),
  initialValues: {
    name: props.unit.name,
  },
})

const [name, nameAttrs] = defineField('name')

const menuOpen = ref(false)

const menuItems: DropdownMenuItem[] = [
  {
    label: 'Редактировать',
    icon: 'i-heroicons-pencil',
    color: 'primary',
    onSelect: () => {
      isEditing.value = !isEditing.value
    },
  },
  {
    label: 'Удалить',
    icon: 'i-heroicons-trash',
    color: 'error',
    onSelect: onDelete,
  },
]

const toast = useToast()

const onSubmit = handleSubmit(async (values) => {
  isSaving.value = true
  try {
    await $fetch<{ unit: Unit }>(createUrl({ url: API_URLS.units.changeName }), {
      method: 'PATCH',
      body: {
        id: props.unit.id,
        name: values.name,
      },
    })
    toast.add({ title: 'Имя подразделения успешно изменено', color: 'success' })
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
          v-if="!isEditing"
          class="font-medium"
        >
          {{ unit.name }}
        </h3>
        <form
          v-else
          class="flex items-center gap-2"
          @submit="onSubmit"
        >
          <UInput
            v-model="name"
            v-bind="nameAttrs"
            autofocus
            @keydown.esc="isEditing = false"
          />
          <UButton
            color="primary"
            size="sm"
            type="submit"
            :loading="isSaving"
            :disabled="!meta.valid || name === unit.name"
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
