export const API_URLS = {
  applications: 'applications',
  signIn: 'sign-in',
  signUp: 'sign-up',
  main: 'main',
  authByToken: 'auth',
  changeStatus: 'applications/change-status',
  delete: 'applications/delete',
  units: {
    all: 'units',
    changeName: 'units/change-name',
    delete: 'units/delete',
    create: 'units/create',
  },
}

export const createUrl = ({ baseURL, url }: { baseURL?: string, url: string }) => {
  const config = useRuntimeConfig()

  return (new URL(url, baseURL ?? config.public.backendUrl)).toString()
}
