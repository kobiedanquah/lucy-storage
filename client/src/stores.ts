import { createStore } from "solid-js/store";
import type { User } from "./types/user";

interface AppState {
  user: User;
}

export const [appState, setAppState] = createStore<AppState>({
  user: {} as User,
});
