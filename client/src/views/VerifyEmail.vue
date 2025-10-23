<script setup lang="ts">
import { ref } from 'vue'

const pin = ref('')
const loading = ref(false)
const resending = ref(false)
const error = ref('')
const success = ref('')

const userEmail = 'jane.doe@example.com' //TODO: use email from signup response

const handleVerify = () => {
    error.value = ''
    success.value = ''

    if (pin.value.length !== 6 || !/^\d+$/.test(pin.value)) {
        error.value = 'Please enter a valid 6-digit code.'
        return
    }

    loading.value = true

    setTimeout(() => {
        loading.value = false

        if (pin.value === '123456') {
            success.value = 'Your email has been successfully verified!'

        } else {
            error.value = 'Invalid verification code. Please try again.'
        }
    }, 1500)
}

const handleResend = () => {
    resending.value = true
    success.value = ''
    error.value = ''

    setTimeout(() => {
        resending.value = false
        success.value = 'A new verification code has been sent to your email.'
    }, 1200)
}
</script>


<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
        <div class="max-w-md w-full bg-white shadow-lg rounded-2xl p-8">
            <h2 class="text-2xl font-bold text-slate-800 mb-2 text-center">Verify Your Email</h2>
            <p class="text-slate-600 text-center mb-6">
                We've sent a 6-digit verification code to <span class="font-medium">{{ userEmail }}</span>.
                Please enter it below to verify your account.
            </p>

            <form @submit.prevent="handleVerify" class="space-y-5">
                <div>
                    <label for="pin" class="block text-sm font-medium text-slate-700 mb-1">Verification Code</label>
                    <input v-model="pin" type="text" id="pin" maxlength="6"
                        class="w-full border border-slate-300 rounded-lg px-3 py-2 text-center tracking-widest text-lg focus:outline-none focus:ring-2 focus:ring-slate-500"
                        placeholder="123456" required />
                </div>

                <p v-if="error" class="text-red-600 text-sm text-center">{{ error }}</p>

                <p v-if="success" class="text-green-600 text-sm text-center">{{ success }}</p>

                <button type="submit"
                    class="w-full bg-slate-700 text-white font-medium py-2 rounded-lg hover:bg-slate-800 transition disabled:opacity-50"
                    :disabled="loading">
                    {{ loading ? 'Verifying...' : 'Verify Email' }}
                </button>
            </form>

            <div class="mt-6 text-sm text-center text-slate-600">
                Didn't receive the code?
                <button @click="handleResend" class="text-slate-700 font-medium hover:underline" :disabled="resending">
                    {{ resending ? 'Resending...' : 'Resend Code' }}
                </button>
            </div>
        </div>
    </div>
</template>
