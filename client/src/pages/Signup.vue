<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { createUser } from "../lib/api";

const router = useRouter();

const name = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");
const error = ref("");
const loading = ref(false);

const handleSubmit = async () => {
    // const store = useAppStateStore()
  
    error.value = "";

    if (password.value !== confirmPassword.value) {
        error.value = "Passwords do not match.";
        return;
    }

    const userInfo = {
        name: name.value,
        email: email.value,
        password: password.value,
    };

    loading.value = true;
    try {
        const data = await createUser(userInfo);
        console.log(data);
        
        // setAppState("user", data);
        router.push("/verification");
    } catch (e: any) {
        loading.value = false;
        error.value = e.message || "An error occurred.";
    }
};
</script>

<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
        <div class="max-w-md w-full bg-white shadow-lg rounded-2xl p-8">
            <h2 class="text-2xl font-bold text-gray-800 mb-6 text-center">Create an Account</h2>

            <form @submit.prevent="handleSubmit" class="space-y-5">
                <div>
                    <label for="name" class="form-label">Full Name</label>
                    <input id="name" type="text" v-model="name" placeholder="Jane Doe" required class="text-input" />
                </div>

                <div>
                    <label for="email" class="form-label">Email Address</label>
                    <input
                        id="email"
                        type="email"
                        v-model="email"
                        placeholder="jane@example.com"
                        required
                        class="text-input"
                    />
                </div>

                <div>
                    <label for="password" class="form-label">Password</label>
                    <input
                        id="password"
                        type="password"
                        v-model="password"
                        placeholder="••••••••"
                        required
                        minlength="6"
                        class="text-input"
                    />
                </div>

                <div>
                    <label for="confirm" class="form-label">Confirm Password</label>
                    <input
                        id="confirm"
                        type="password"
                        v-model="confirmPassword"
                        placeholder="••••••••"
                        required
                        class="text-input"
                    />
                </div>

                <p v-if="error" class="text-red-600 text-sm">{{ error }}</p>

                <button
                    type="submit"
                    :disabled="loading"
                    class="w-full bg-slate-700 text-white font-medium py-2 rounded-lg hover:bg-slate-800 transition disabled:opacity-50"
                >
                    {{ loading ? "Signing Up..." : "Sign Up" }}
                </button>
            </form>

            <p class="mt-6 text-sm text-center text-gray-600">
                Already have an account?
                <RouterLink to="/login" class="font-bold text-slate-700 hover:underline"> Log in </RouterLink>
            </p>
        </div>
    </div>
</template>
