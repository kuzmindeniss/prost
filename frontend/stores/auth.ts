import { defineStore } from 'pinia'

export interface User {
  id: string
  name: string
  surname: string
  email: string
  role: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)

  const setAuth = ({
    newUser,
  }: {
    newUser: User | null
  }) => {
    user.value = newUser
  }

  const isAdmin = computed(() => user.value?.role === 'admin')

  return { user, setAuth, isAdmin }
}, {
  persist: {
    storage: piniaPluginPersistedstate.localStorage(),
    pick: ['token'],
  },
})
