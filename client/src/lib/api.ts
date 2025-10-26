import type { User } from "./types";

const baseUrl = "http://localhost:8080/api/v1";

export async function createUser(payload: { name: string; email: string; password: string }): Promise<User> {
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
      } else {
        const resBody = await response.json();
        throw new Error(resBody.message);
      }
    }

    const data = await response.json();

    return data;
  } catch (e) {
    console.log(e);
    throw e;
  }
}

export async function verifyUserEmail(payload: { code: string; email: string }): Promise<User> {
  try {
    const response = await fetch(`${baseUrl}/auth/verify`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      const resBody = await response.json();
      throw new Error(resBody.message);
    }
    const data = await response.json();

    return data;
  } catch (e) {
    console.log(e);
    throw e;
  }
}

export async function requestVerificationCode(email: string) {
  try {
    const response = await fetch(`${baseUrl}/auth/verify/request`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email: email }),
    });

    if (!response.ok) {
      const resBody = await response.json();
      throw new Error(resBody.message);
    }
    const data = await response.json();

    return data;
  } catch (e) {
    console.log(e);
    throw e;
  }
}

export async function getUserSession(payload: { email: string; password: string }) {
  try {
    const response = await fetch(`${baseUrl}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      const resBody = await response.json();

      throw new Error(resBody.message);
    }

    const data = await response.json();

    return data;
  } catch (e) {
    console.log(e);
    throw e;
  }
}
