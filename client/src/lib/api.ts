const baseUrl = "http://localhost:8080/api/v1";

export async function createUser(payload: {
  name: string;
  email: string;
  password: string;
}): Promise<any> {
  const response = await fetch(`${baseUrl}/auth/register`, {
    method: "POST",
    body: JSON.stringify(payload),
  });

  if (!response.ok) {
    if (response.status >= 500) {
      console.log("");
    }else{
      
    }
  }

  return response.json();
}

