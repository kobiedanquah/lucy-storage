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
        throw new Error("a user with your email already exists");
      } else if (response.status >= 500) {
        throw new Error(`server could not process your request.`);
      } else {
        let message = "";
        response.json().then((message) => {
          message = message;
        });
        throw new Error(message)
      }
    }

    const data = await response.json();

    return data;
  } catch (e) {
    console.log(e);
    throw e;
  }
}
