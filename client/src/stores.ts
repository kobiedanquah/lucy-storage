import { createStore } from "solid-js/store";
import type { User } from "./types/user";
import type { UserSession } from "./types/session";

type AppState = {
  user: User;
  session?: UserSession;
};

export const [appState, setAppState] = createStore<AppState>({
  user: {} as User,
 
});
