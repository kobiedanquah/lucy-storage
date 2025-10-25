import { createSignal } from "solid-js";
import { requestVerificationCode, verifyUserEmail } from "../lib/api";
import { useNavigate } from "@solidjs/router";
import { appState } from "../stores";

export default function VerifyEmailPage() {
  const navigate = useNavigate();

  console.log(appState.user);
  if (Object.keys(appState.user).length === 0) {
    navigate("/signup", { replace: true });
  }

  const [pin, setPin] = createSignal("");
  const [loading, setLoading] = createSignal(false);
  const [resending, setResending] = createSignal(false);
  const [error, setError] = createSignal("");
  const [success, setSuccess] = createSignal("");

  const userEmail = appState.user.email;
  //   const maskedEmail = createMemo(() => {
  //     const [name, domain] = userEmail.split("@");
  //     return `${name[0]}***@${domain}`;
  //   });

  const handleVerification = async (e: Event) => {
    e.preventDefault();
    setError("");
    setSuccess("");

    if (pin().length !== 6 || !/^\d+$/.test(pin())) {
      setError("Please enter a valid 6-digit code.");
      return;
    }

    setLoading(true);
    try {
      const data = await verifyUserEmail({ email: userEmail, code: pin() });
      console.log(data);
      setSuccess("Your email has been successfully verified!");
      navigate("/", { replace: true });
    } catch (e: any) {
      setLoading(false);
      setError(e.message);
    }
  };

  const handleResend = async () => {
    setResending(true);
    setSuccess("");
    setError("");

    try {
      const data = await requestVerificationCode(userEmail);
      console.log(data);
      setResending(false);
      setSuccess("A new verification code has been sent to your email.");
      navigate("/verification");
    } catch (e: any) {
      setLoading(false);
      setError(e.message);
    }
  };

  return (
    <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
      <div class="max-w-md w-full bg-white shadow-lg rounded-2xl p-8">
        <h2 class="text-2xl font-bold text-slate-800 mb-2 text-center">
          Verify Your Email
        </h2>
        <p class="text-slate-600  mb-6">
          We've sent a 6-digit verification code to{" "}
          <span class="font-medium">{userEmail}</span>. Please enter it below.
        </p>

        <form onSubmit={handleVerification} class="space-y-5">
          <div>
            <label
              for="pin"
              class="block text-sm font-medium text-slate-700 mb-1"
            >
              Verification Code
            </label>
            <input
              id="pin"
              type="text"
              value={pin()}
              maxlength="6"
              onInput={(e) => setPin(e.target.value)}
              placeholder="123456"
              required
              class="text-input py-2 text-center tracking-widest text-lg "
            />
          </div>

          {error() && <p class="text-red-600 text-sm text-center">{error()}</p>}

          {success() && (
            <p class="text-green-600 text-sm text-center">{success()}</p>
          )}

          <button
            type="submit"
            disabled={loading()}
            class="w-full bg-slate-700 text-white font-medium py-2 rounded-lg hover:bg-slate-800 transition disabled:opacity-50"
          >
            {loading() ? "Verifying..." : "Verify Email"}
          </button>
        </form>

        <div class="mt-6 text-sm text-center text-slate-600">
          Didn't receive the code?{" "}
          <button
            onClick={handleResend}
            disabled={resending()}
            class="text-slate-700 font-bold hover:underline"
          >
            {resending() ? "Resending..." : "Resend Code"}
          </button>
        </div>
      </div>
    </div>
  );
}
