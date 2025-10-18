import { reactive, readonly } from "vue";
import type { AuthState } from "./types/auth_state";
import type { User } from "./types/user";

const authState = reactive<AuthState>({
    user: null,
    accessToken: null,
    refreshToken: null
})

function setUser(user: User, accessToken: string, refreshToken: string) {
    authState.user = user
    authState.accessToken = accessToken
    authState.refreshToken = refreshToken
}

function clearUser() {
    authState.user = null
    authState.accessToken = null
    authState.refreshToken = null
}

export default {
    state: readonly(authState),
    setUser,
    clearUser
}