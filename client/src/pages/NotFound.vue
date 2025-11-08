<script setup>
import { RouterLink } from "vue-router";

const currentYear = new Date().getFullYear();
</script>

<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-gray-50 text-slate-700 px-6">
    <div class="text-center">
      <h1 class="text-8xl font-extrabold text-slate-800 mb-4">404</h1>
      <h2 class="text-2xl font-semibold mb-2">Page Not Found</h2>
      <p class="text-slate-600 mb-6">
        The page you're looking for doesn't exist or has been moved.
      </p>

      <RouterLink
        to="/"
        class="inline-block bg-slate-700 text-white px-6 py-2 rounded-lg font-medium hover:bg-slate-800 transition"
      >
        Go Home
      </RouterLink>
    </div>

    <footer class="mt-10 text-sm text-slate-500">
      © {{ currentYear }} Lucy Storage. All rights reserved.
    </footer>
  </div>
</template>
