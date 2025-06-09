<script setup lang="ts">
import * as v from 'valibot'
import { toTypedSchema } from '@vee-validate/valibot'
import type { Unit } from '~/types/unit'

definePageMeta({
  middleware: 'auth',
})

const authStore = useAuthStore()

const isCreatingModalOpen = ref(false)

const toast = useToast()

const { meta, defineField, resetForm, handleSubmit } = useForm({
  validationSchema: toTypedSchema(
    v.object({
      name: v.pipe(v.string(), v.minLength(1, 'Введите имя')),
    }),
  ),
  initialValues: {
    name: '',
  },
})

const [name, nameAttrs] = defineField('name')

const { data, refresh } = await useFetch<{ units: Unit[] }>(
  createUrl({ url: API_URLS.units.all }),
  {
    onResponseError: (error) => {
      toast.add({
        title: `Ошибка при получении подразделений: ${error}`,
        color: 'error',
      })
    },
    headers: getAuthHeaders(),
  },
)

const createUnit = handleSubmit(async (values) => {
  await $fetch(createUrl({ url: API_URLS.units.create }), {
    method: 'POST',
    body: {
      name: values.name,
    },
    headers: getAuthHeaders(),
  })
  refresh()
  isCreatingModalOpen.value = false
  resetForm()
})
</script>

<template>
  <div class="container mx-auto p-6 max-w-200">
    <div class="flex flex-col md:flex-row items-center justify-between md:items-baseline mb-6">
      <h1 class="text-2xl font-bold mb-6">
        Подразделения
      </h1>
      <UModal
        v-if="authStore.isAdmin"
        v-model:open="isCreatingModalOpen"
        title="Создание подразделения"
      >
        <UButton color="primary">
          Создать подразделение
        </UButton>
        <template #body>
          <form
            class="flex flex-col gap-4"
            @submit="createUnit"
          >
            <UFormField
              label="Название"
              required
            >
              <UInput
                v-model="name"
                v-bind="nameAttrs"
                class="w-full"
              />
            </UFormField>
            <div class="flex justify-center gap-4">
              <UButton
                color="primary"
                type="submit"
                :disabled="!meta.valid"
                class="w-full justify-center"
              >
                Создать
              </UButton>
            </div>
          </form>
        </template>
      </UModal>
    </div>
    <div class="grid grid-cols-1">
      <div class="bg-white rounded-lg shadow p-4">
        <div
          v-if="Boolean(data?.units?.length)"
          class="space-y-3"
        >
          <UnitsListItem
            v-for="unit in data?.units"
            :key="unit.id"
            :unit="unit"
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
