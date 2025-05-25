import { defineStore } from 'pinia'

export interface User {
  name: string
  surname: string
  email: string
  role: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)

  const setAuth = ({
    newUser,
    newToken,
  }: {
    newUser: User | null
    newToken: string | null
  }) => {
    user.value = newUser
    token.value = newToken
  }

  return { user, token, setAuth }
}, {
  persist: {
    storage: piniaPluginPersistedstate.localStorage(),
    pick: ['token'],
  },
})
