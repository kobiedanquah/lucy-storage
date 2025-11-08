import { defineStore } from "pinia";
import type { User, UserSession } from "../lib/types";
import { ref } from "vue";

export const useAppStateStore = defineStore("appState", {
  state: () => {
    const user = ref<User | undefined>(undefined);
    const session = ref<UserSession | undefined>(undefined);

    return { user, session };
  },
});
