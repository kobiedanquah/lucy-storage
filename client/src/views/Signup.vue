<script setup lang="ts">
import { ref } from "vue";
import { createUser } from "../lib/api";

const form = ref({
  name: "",
  email: "",
  password: "",
  confirmPassword: "",
});

const error = ref("");
const loading = ref(false);

const handleSubmit = async () => {
  error.value = "";

  if (form.value.password !== form.value.confirmPassword) {
    error.value = "Passwords do not match.";
    return;
  }

  loading.value = true;

  try {
    const data = await createUser(form.value);
    console.log(data);
  } catch (e) {
    console.error(e);
  }

  // setTimeout(() => {
  //     loading.value = false
  //     alert(`Welcome, ${form.value.name}! Your account has been created.`)

  //     form.value = { name: '', email: '', password: '', confirmPassword: '' }
  // }, 1500)
};
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
    <div class="max-w-md w-full bg-white shadow-lg rounded-2xl p-8">
      <h2 class="text-2xl font-bold text-gray-800 mb-6 text-center">
        Create an Account
      </h2>

      <form @submit.prevent="handleSubmit" class="space-y-5">
        <div>
          <label for="name" class="block text-sm font-medium text-gray-700 mb-1"
            >Full Name</label
          >
          <input
            v-model="form.name"
            type="text"
            id="name"
            class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-slate-500"
            placeholder="Jane Doe"
            required
          />
        </div>

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
            placeholder="jane@example.com"
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
            minlength="8"
          />
        </div>

        <div>
          <label
            for="confirm"
            class="block text-sm font-medium text-gray-700 mb-1"
            >Confirm Password</label
          >
          <input
            v-model="form.confirmPassword"
            type="password"
            id="confirm"
            class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-slate-500"
            placeholder="••••••••"
            required
          />
        </div>

        <p v-if="error" class="text-red-600 text-sm">{{ error }}</p>

        <button
          type="submit"
          class="w-full bg-slate-600 text-white font-medium py-2 rounded-lg hover:bg-slate-700 transition disabled:opacity-50"
          :disabled="loading"
        >
          {{ loading ? "Signing Up..." : "Sign Up" }}
        </button>
      </form>

      <p class="mt-6 text-sm text-center text-gray-600">
        Already have an account?
        <a href="/login" class="text-slate-600 font-bold hover:underline"
          >Log in</a
        >
      </p>
    </div>
  </div>
</template>
