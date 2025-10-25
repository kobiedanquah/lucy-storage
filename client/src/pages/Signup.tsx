import { createSignal } from "solid-js";
import { createUser } from "../lib/api";
import { useNavigate } from "@solidjs/router";

export default function Signup() {
  const navigate = useNavigate();
  const [name, setName] = createSignal("");
  const [email, setEmail] = createSignal("");
  const [password, setPassword] = createSignal("");
  const [confirmPassword, setConfirmPassword] = createSignal("");
  const [error, setError] = createSignal("");
  const [loading, setLoading] = createSignal(false);

  const handleSubmit = async (e: Event) => {
    e.preventDefault();
    setError("");

    if (password() !== confirmPassword()) {
      setError("Passwords do not match.");
      return;
    }

    const userInfo = { name: name(), email: email(), password: password() };

    setLoading(true);
    try {
      const data = await createUser(userInfo);
      console.log(data);
      navigate("/signup/verify-email");
    } catch (e) {
      console.error(e);
    }

    setLoading(false);
  };

  return (
    <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
      <div class="max-w-md w-full bg-white shadow-lg rounded-2xl p-8">
        <h2 class="text-2xl font-bold text-gray-800 mb-6 text-center">
          Create an Account
        </h2>

        <form onSubmit={handleSubmit} class="space-y-5">
          <div>
            <label for="name" class="form-label">
              Full Name
            </label>
            <input
              id="name"
              type="text"
              value={name()}
              onInput={(e) => setName(e.target.value)}
              placeholder="Jane Doe"
              required
              class="text-input"
            />
          </div>

          <div>
            <label for="email" class="form-label">
              Email Address
            </label>
            <input
              id="email"
              type="email"
              value={email()}
              onInput={(e) => setEmail(e.target.value)}
              placeholder="jane@example.com"
              required
              class="text-input"
            />
          </div>

          <div>
            <label for="password" class="form-label">
              Password
            </label>
            <input
              id="password"
              type="password"
              value={password()}
              onInput={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
              required
              minlength="6"
              class="text-input"
            />
          </div>

          <div>
            <label for="confirm" class="form-label">
              Confirm Password
            </label>
            <input
              id="confirm"
              type="password"
              value={confirmPassword()}
              onInput={(e) => setConfirmPassword(e.target.value)}
              placeholder="••••••••"
              required
              class="text-input"
            />
          </div>

          {error() && <p class="text-red-600 text-sm">{error()}</p>}

          <button
            type="submit"
            disabled={loading()}
            class="w-full bg-slate-700 text-white font-medium py-2 rounded-lg hover:bg-slate-800 transition disabled:opacity-50"
          >
            {loading() ? "Signing Up..." : "Sign Up"}
          </button>
        </form>

        <p class="mt-6 text-sm text-center text-gray-600">
          Already have an account?{" "}
          <a href="/login" class="font-bold text-slate-700 hover:underline">
            Log in
          </a>
        </p>
      </div>
    </div>
  );
}
