<script setup lang="ts">
import { ref } from "vue";
import { useRouter, RouterLink } from "vue-router";
import { getUserSession } from "../lib/api";

const router = useRouter();

const email = ref("");
const password = ref("");
const error = ref("");
const loading = ref(false);

async function handleLogin() {
  error.value = "";

  if (!email.value || !password.value) {
    error.value = "Please fill in all fields.";
    return;
  }

  loading.value = true;

  try {
    const data = await getUserSession({
      email: email.value,
      password: password.value,
    });
    localStorage.setItem("session", JSON.stringify(data));
    router.replace("/");
  } catch (e:any) {
    console.log(e.message);
    loading.value = false;
    error.value = e.message || "Login failed.";
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
    <div class="max-w-md w-full bg-white shadow-lg rounded-2xl p-8">
      <h2 class="text-2xl font-bold text-slate-800 mb-6 text-center">
        Welcome Back
      </h2>

      <form @submit.prevent="handleLogin" class="space-y-5">
        <div>
          <label for="email" class="block text-sm font-medium text-slate-700 mb-1">
            Email Address
          </label>
          <input
            id="email"
            type="email"
            v-model="email"
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
            v-model="password"
            placeholder="••••••••"
            required
            class="text-input"
          />
        </div>

        <p v-if="error" class="text-red-600 text-sm text-center">{{ error }}</p>

        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-slate-700 text-white font-medium py-2 rounded-lg hover:bg-slate-800 transition disabled:opacity-50"
        >
          {{ loading ? "Logging In..." : "Login" }}
        </button>
      </form>

      <div class="mt-6 text-sm text-center text-slate-600">
        <p>
          Don’t have an account?
          <RouterLink to="/signup" class="text-slate-700 font-medium hover:underline">
            Sign up
          </RouterLink>
        </p>
        <p class="mt-2">
          <RouterLink
            to="/forgot-password"
            class="text-slate-700 font-medium hover:underline"
          >
            Forgot password?
          </RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>


