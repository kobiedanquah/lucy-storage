<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { requestVerificationCode, verifyUserEmail } from "../lib/api";
import { useAppStateStore } from "../stores/appState";
const router = useRouter();

const pin = ref("");
const loading = ref(false);
const resending = ref(false);
const error = ref("");
const success = ref("");

const appState = useAppStateStore();

const userEmail = appState.user?.email || "";

// Redirect to signup if user is missing
onMounted(() => {
    if (!userEmail) {
        router.replace("/signup");
    }
});

async function handleVerification() {
    error.value = "";
    success.value = "";

    if (pin.value.length !== 6 || !/^\d+$/.test(pin.value)) {
        error.value = "Please enter a valid 6-digit code.";
        return;
    }

    loading.value = true;
    try {
        const data = await verifyUserEmail({ email: userEmail, code: pin.value });
        console.log(data);
        success.value = "Your email has been successfully verified!";
        router.replace("/");
    } catch (e: any) {
        loading.value = false;
        error.value = e.message || "Verification failed.";
    }
}

async function handleResend() {
    resending.value = true;
    success.value = "";
    error.value = "";

    try {
        const data = await requestVerificationCode(userEmail);
        console.log(data);
        resending.value = false;
        success.value = "A new verification code has been sent to your email.";
        router.push("/verification");
    } catch (e:any) {
        resending.value = false;
        error.value = e.message || "Failed to resend code.";
    }
}
</script>

<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
        <div class="max-w-md w-full bg-white shadow-lg rounded-2xl p-8">
            <h2 class="text-2xl font-bold text-slate-800 mb-2 text-center">Verify Your Email</h2>
            <p class="text-slate-600 mb-6 text-center">
                We've sent a 6-digit verification code to
                <span class="font-medium">{{ userEmail }}</span
                >. Please enter it below.
            </p>

            <form @submit.prevent="handleVerification" class="space-y-5">
                <div>
                    <label for="pin" class="block text-sm font-medium text-slate-700 mb-1"> Verification Code </label>
                    <input
                        id="pin"
                        type="text"
                        v-model="pin"
                        maxlength="6"
                        placeholder="123456"
                        required
                        class="text-input py-2 text-center tracking-widest text-lg"
                    />
                </div>

                <p v-if="error" class="text-red-600 text-sm text-center">{{ error }}</p>
                <p v-if="success" class="text-green-600 text-sm text-center">{{ success }}</p>

                <button
                    type="submit"
                    :disabled="loading"
                    class="w-full bg-slate-700 text-white font-medium py-2 rounded-lg hover:bg-slate-800 transition disabled:opacity-50"
                >
                    {{ loading ? "Verifying..." : "Verify Email" }}
                </button>
            </form>

            <div class="mt-6 text-sm text-center text-slate-600">
                Didn't receive the code?
                <button @click="handleResend" :disabled="resending" class="text-slate-700 font-bold hover:underline">
                    {{ resending ? "Resending..." : "Resend Code" }}
                </button>
            </div>
        </div>
    </div>
</template>
