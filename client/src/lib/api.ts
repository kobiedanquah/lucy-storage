import type { User } from "../types/user";

const baseUrl = "http://localhost:8080/api/v1";

export async function createUser(payload: {
  name: string;
  email: string;
  password: string;
}): Promise<User> {
  try {
    const response = await fetch(`${baseUrl}/auth/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      if (response.status == 409) {
        throw new Error("User already exists");
      } else {
        throw new Error(`error with status ${response.status}`);
      }
    }

    const data = await response.json();

    return data;
  } catch (err: any) {
    if (err instanceof Error) throw err;
    throw new Error("An unknown error occurred while creating the user.");
  }
}
