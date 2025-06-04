import type { Unit } from './unit'

export interface User {
  id: string
  name: string
  surname: string
  email: string
  role: 'admin' | 'user'
}

export interface UserTg {
  id: string
  name: string
  tgUsername: string
  unit: Unit
}
