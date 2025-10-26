export interface User {
  id: string;
  name: string;
  email: string;
  profilePhoto: string;
  createdAt: Date;
  lastModified: Date;
}


export interface UserSession {
  user: User
  refreshToken: string
  refreshExpiresAt: Date
  accessToken: string
  accessExpiresAt: Date
}