export const API_URLS = {
  signIn: 'sign-in',
  signUp: 'sign-up',
  main: 'main',
  authByToken: 'auth',
  changeStatus: 'applications/change-status',
}

export const createUrl = ({ baseURL, url }: { baseURL?: string, url: string }) => {
  const config = useRuntimeConfig()

  return (new URL(url, baseURL ?? config.public.backendUrl)).toString()
}
