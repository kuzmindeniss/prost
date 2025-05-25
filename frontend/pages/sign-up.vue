<script setup lang="ts">
import * as v from 'valibot'
import { toTypedSchema } from '@vee-validate/valibot'

const { errors, handleSubmit, meta, defineField, setFieldError } = useForm({
  validationSchema: toTypedSchema(
    v.object({
      email: v.pipe(v.string(), v.email('Невалидная почта')),
      name: v.pipe(v.string(), v.minLength(1, 'Введите имя')),
      surname: v.pipe(v.string(), v.minLength(1, 'Введите фамилию')),
      password: v.pipe(v.string(), v.minLength(6, 'Пароль должен быть не менее 6 символов')),
    }),
  ),
  initialValues: {
    name: '',
    surname: '',
    email: '',
    password: '',
  },
})

const [email, emailAttrs] = defineField('email')
const [name, nameAttrs] = defineField('name')
const [surname, surnameAttrs] = defineField('surname')
const [password, passwordAttrs] = defineField('password')

const toast = useToast()

const authStore = useAuthStore()

const onSubmit = handleSubmit(async (values) => {
  try {
    const res = await $fetch<{ user: User }>(createUrl({ url: API_URLS.signUp }), {
      method: 'POST',
      body: values,
    })

    authStore.setAuth({ newUser: res.user })

    toast.add({ title: 'Вы успешно зарегестрировались', color: 'success' })
  }
  catch (e: any) {
    setFieldError('email', e?.data?.error ?? 'Ошибка при входе')
  }
})
</script>

<template>
  <div>
    <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8 pt-20">
      <div class="sm:mx-auto sm:w-full sm:max-w-sm mb-10">
        <h1 class="text-center tracking-tight">
          Регистрация
        </h1>
      </div>

      <div class="sm:mx-auto sm:w-full sm:max-w-sm">
        <form
          class="space-y-4"
          @submit="onSubmit"
        >
          <UFormField
            label="Почта"
            name="email"
            size="xl"
            :error="errors.email"
            required
          >
            <UInput
              v-model="email"
              v-bind="emailAttrs"
              class="w-full"
              size="xl"
              placeholder="mail@mail.ru"
            />
          </UFormField>

          <UFormField
            label="Имя"
            name="name"
            size="xl"
            :error="errors.name"
            required
          >
            <UInput
              v-model="name"
              v-bind="nameAttrs"
              class="w-full"
              size="xl"
              placeholder="Иван"
            />
          </UFormField>

          <UFormField
            label="Фамилия"
            name="surname"
            size="xl"
            :error="errors.surname"
            required
          >
            <UInput
              v-model="surname"
              v-bind="surnameAttrs"
              class="w-full"
              size="xl"
              placeholder="Иванов"
            />
          </UFormField>

          <UFormField
            label="Пароль"
            name="password"
            :error="errors.password"
            required
          >
            <UInput
              v-model="password"
              v-bind="passwordAttrs"
              type="password"
              class="w-full"
              placeholder="******"
            />
          </UFormField>

          <div class="flex gap-2 flex-col">
            <UButton
              :disabled="!meta.valid"
              type="submit"
              class="justify-center"
            >
              Зарегестрироваться
            </UButton>

            <UButton
              class="justify-center"
              to="/sign-in"
            >
              На страницу входа
            </UButton>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
