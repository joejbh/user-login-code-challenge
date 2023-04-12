import { LoginData } from "../components/LoginForm";

export async function login(data: LoginData) {
  const response = await fetch(
    `${process.env.REACT_APP_SERVER_URL}/api/login`,
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ data }),
    }
  );

  const content = await response.json();

  return content;
}
