<template>
  <div class="flex flex-row h-full relative">
    <Sidebar />
    <div class="flex-grow bg-groy-900 text-gray-100 p-4 sm:p-10 overflow-y-auto">
      <Nuxt />
    </div>
    <div class="fixed w-full p-2 top-0 z-10" v-if="error">
      <div class="mx-auto max-w-lg text-rod-400 py-4 px-6 border-rod-400 border bg-rod-900">
        <div class="flex items-center justify-between">
          <div class="font-bold">Error!</div>
          <button class="hover:bg-rod-400 hover:text-rod-900 rounded-sm"
                  @click="closeToast">
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-5"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <div class="mt-3">
          {{ error }}
        </div>
      </div>
    </div>
  </div>
</template>

<style>
html,
body {
  @apply bg-groy-900;
  height: 100%;
}
#__nuxt,
#__layout {
  height: 100%;
}
</style>

<script>
export default {
  computed: {
    error() {
      return this.$store.state.local.error;
    }
  },
  methods: {
    closeToast() {
      this.$store.commit('local/error', '')
    }
  }
}
</script>