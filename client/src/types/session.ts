import type { User } from "./user";

export interface UserSession {
  user: User
  refreshToken: string
  refreshExpiresAt: Date
  accessToken: string
  accessExpiresAt: Date
}