import type { Unit } from './unit'
import type { UserTg } from './user'

export type ApplicationStatus = 'pending' | 'done'

export interface Application {
  id: string
  text: string
  status: ApplicationStatus
  unit: Unit
  user_tg: UserTg
}
