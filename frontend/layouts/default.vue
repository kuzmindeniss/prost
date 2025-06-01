<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'

const authStore = useAuthStore()

const authedNavigationItems: NavigationMenuItem[] = [
  {
    label: 'Заявки',
    to: '/applications',
    icon: 'i-lucide-book-open',
  },
  {
    label: 'Подразделения',
    to: '/units',
    icon: 'i-lucide-building',
  },
  {
    label: 'Выйти',
    to: '/sign-in',
    onSelect: (e) => {
      e.preventDefault()
      authStore.setAuth({ newUser: null })
      useCookie('Authorization').value = ''
    },
    icon: 'i-lucide-log-out',
  },
]

const unauthedNavigationItems: NavigationMenuItem[] = [
  {
    label: 'Логин',
    to: '/sign-in',
    icon: 'i-lucide-log-in',
  },
  {
    label: 'Регистрация',
    to: '/sign-up',
    icon: 'i-lucide-user-plus',
  },
]

const navigationItems = computed(() => authStore.user?.email ? authedNavigationItems : unauthedNavigationItems)
</script>

<template>
  <div class="absolute top-0 z-[-2] h-screen w-screen bg-white bg-[radial-gradient(100%_50%_at_50%_0%,rgba(0,163,255,0.13)_0,rgba(0,163,255,0)_50%,rgba(0,163,255,0)_100%)]">
    <div class="flex flex-col h-screen">
      <UNavigationMenu
        :items="navigationItems"
        orientation="horizontal"
        class="flex-shrink-0 p-4 justify-center"
      />
      <slot />
    </div>
  </div>
</template>
