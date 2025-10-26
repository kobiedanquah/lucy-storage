import { A, useNavigate } from "@solidjs/router";
import { createSignal } from "solid-js";
import { getUserSession } from "../lib/api";

export default function Login() {
  const navigate = useNavigate();
  const [email, setEmail] = createSignal("");
  const [password, setPassword] = createSignal("");
  const [error, setError] = createSignal("");
  const [loading, setLoading] = createSignal(false);

  async function handleLogin(e: Event) {
    e.preventDefault();
    setError("");

    if (!email() || !password()) {
      setError("Please fill in all fields.");
      return;
    }

    setLoading(true);

    try {
      const data = await getUserSession({ email: email(), password: password() });
      localStorage.setItem("session", JSON.stringify(data));
      navigate("/", { replace: true });
    } catch (e: any) {
      console.log(e.message);
      setLoading(false);
      setError(e.message);
    }
  }

  return (
    <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
      <div class="max-w-md w-full bg-white shadow-lg rounded-2xl p-8">
        <h2 class="text-2xl font-bold text-slate-800 mb-6 text-center">Welcome Back</h2>

        <form onSubmit={handleLogin} class="space-y-5">
          <div>
            <label for="email" class="block text-sm font-medium text-slate-700 mb-1">
              Email Address
            </label>
            <input
              id="email"
              type="email"
              value={email()}
              onInput={(e) => setEmail(e.target.value)}
              placeholder="you@example.com"
              required
              class="text-input"
            />
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-slate-700 mb-1">
              Password
            </label>
            <input
              id="password"
              type="password"
              value={password()}
              onInput={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
              required
              class="text-input"
            />
          </div>

          {error() && <p class="text-red-600 text-sm text-center">{error()}</p>}

          <button
            type="submit"
            disabled={loading()}
            class="w-full bg-slate-700 text-white font-medium py-2 rounded-lg hover:bg-slate-800 transition disabled:opacity-50"
          >
            {loading() ? "Logging In..." : "Login"}
          </button>
        </form>

        <div class="mt-6 text-sm text-center text-slate-600">
          <p>
            Don’t have an account?{" "}
            <A href="/signup" class="text-slate-700 font-medium hover:underline">
              Sign up
            </A>
          </p>
          <p class="mt-2">
            <A href="/forgot-password" class="text-slate-700 font-medium hover:underline">
              Forgot password?
            </A>
          </p>
        </div>
      </div>
    </div>
  );
}
