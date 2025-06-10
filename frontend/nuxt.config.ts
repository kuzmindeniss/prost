// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    '@nuxt/ui',
    '@vee-validate/nuxt',
    '@nuxt/eslint',
    '@pinia/nuxt',
    '@vueuse/nuxt',
  ],
  ssr: false,
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'],
  colorMode: {
    preference: 'light',
  },
  runtimeConfig: {
    public: {
      backendUrl: process.env.NUXT_PUBLIC_BACKEND_URL || '',
    },
  },
  compatibilityDate: '2025-05-15',
  nitro: {
    prerender: {
      crawlLinks: false,
      ignore: [],
      routes: [],
    },
  },
  eslint: {
    config: {
      stylistic: true,
    },
  },
})
