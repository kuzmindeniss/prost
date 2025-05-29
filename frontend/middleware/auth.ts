export default defineNuxtRouteMiddleware(async () => {
  const token = useCookie('Authorization')
  if (!token.value) {
    return navigateTo('/sign-in')
  }

  console.log(token.value)

  const authStore = useAuthStore()

  try {
    const res = await $fetch<{ user: User, token: string }>(createUrl({ url: API_URLS.authByToken }), {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    })
    authStore.setAuth({ newUser: res.user })
  }
  catch {
    return navigateTo('/sign-in')
  }
})
