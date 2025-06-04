<script setup lang="ts">
import * as v from 'valibot'
import { toTypedSchema } from '@vee-validate/valibot'

const { errors, handleSubmit, meta, defineField, setFieldError } = useForm({
  validationSchema: toTypedSchema(
    v.object({
      email: v.pipe(v.string(), v.email('Невалидная почта')),
      password: v.pipe(v.string(), v.minLength(6, 'Пароль должен быть не менее 6 символов')),
    }),
  ),
  initialValues: {
    email: '',
    password: '',
  },
})

const [email, emailAttrs] = defineField('email')
const [password, passwordAttrs] = defineField('password')

const toast = useToast()

const authStore = useAuthStore()

const onSubmit = handleSubmit(async (values) => {
  try {
    const res = await $fetch<{ user: User, token: string }>(createUrl({ url: API_URLS.signIn }), {
      method: 'POST',
      body: values,
    })
    authStore.setAuth({ newUser: res.user })
    useCookie('Authorization').value = res.token
    toast.add({ title: 'Вы успешно вошли', color: 'success' })
    navigateTo('/')
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
          Вход в аккаунт
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
              Войти
            </UButton>

            <UButton
              class="justify-center"
              to="/sign-up"
            >
              На страницу регистрации
            </UButton>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
