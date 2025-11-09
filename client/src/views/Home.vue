<script setup lang="ts">
import { onMounted } from "vue";
import { useRouter } from "vue-router";
import { useAppStateStore } from "../stores/appState";

const router = useRouter();
const appState = useAppStateStore();

onMounted(() => {
    if (!appState.session) {
        router.push("/login");
    }
});
</script>

<template>
    <div class="p-6">
        <h1 class="text-2xl font-bold text-slate-800">Home Page</h1>
    </div>
</template>
