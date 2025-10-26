import { useNavigate } from "@solidjs/router";
import { appState } from "../stores";

export default function Home() {
  const navigate = useNavigate()
  if (!appState.session){
    navigate("/login")
  }
  
  return (
    <>
      <h1>Home Page</h1>
    </>
  );
}
