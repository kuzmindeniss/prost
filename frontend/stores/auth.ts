import { defineStore } from 'pinia'

export interface User {
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

  return { user, setAuth }
}, {
  persist: {
    storage: piniaPluginPersistedstate.localStorage(),
    pick: ['token'],
  },
})
