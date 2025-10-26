import { createStore } from "solid-js/store";
import type { User, UserSession } from "./lib/types";

type AppState = {
  user: User;
  session?: UserSession;
};

export const [appState, setAppState] = createStore<AppState>({
  user: {} as User,
  session: JSON.parse(localStorage.getItem("session") as string) as UserSession,
});
