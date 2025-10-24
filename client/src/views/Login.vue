<script setup lang="ts">
import { ref } from "vue";

const form = ref({
  email: "",
  password: "",
});

const error = ref("");
const loading = ref(false);

const handleLogin = () => {
  error.value = "";

  if (!form.value.email || !form.value.password) {
    error.value = "Please fill out all fields.";
    return;
  }

  loading.value = true;

  setTimeout(() => {
    loading.value = false;

    if (
      form.value.email === "user@example.com" &&
      form.value.password === "password123"
    ) {
      alert("Login successful!");
    } else {
      error.value = "Invalid email or password.";
    }
  }, 1200);
};
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
    <div class="max-w-md w-full bg-white shadow-lg rounded-2xl p-8">
      <h2 class="text-2xl font-bold text-gray-800 mb-6 text-center">
        Welcome Back
      </h2>

      <form @submit.prevent="handleLogin" class="space-y-5">
        <div>
          <label
            for="email"
            class="block text-sm font-medium text-gray-700 mb-1"
            >Email Address</label
          >
          <input
            v-model="form.email"
            type="email"
            id="email"
            class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-slate-500"
            placeholder="your@email.com"
            required
          />
        </div>

        <div>
          <label
            for="password"
            class="block text-sm font-medium text-gray-700 mb-1"
            >Password</label
          >
          <input
            v-model="form.password"
            type="password"
            id="password"
            class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-slate-500"
            placeholder="••••••••"
            required
          />
        </div>

        <p v-if="error" class="text-red-600 text-sm">{{ error }}</p>

        <button
          type="submit"
          class="w-full bg-slate-500 text-white font-medium py-2 rounded-lg hover:bg-slate-700 transition disabled:opacity-50"
          :disabled="loading"
        >
          {{ loading ? "Logging In..." : "Login" }}
        </button>
      </form>

      <div class="mt-6 text-sm text-center text-gray-600">
        <p>
          Don't have an account?
          <a href="/signup" class="text-slate-600 hover:underline">Sign up</a>
        </p>
        <p class="mt-2">
          <a
            href="/forgot-password"
            class="text-slate-500 font-bold hover:underline"
            >Forgot password?</a
          >
        </p>
      </div>
    </div>
  </div>
</template>
