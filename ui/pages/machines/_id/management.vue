<template>
  <div class="flex flex-col gap-8">
    <div class="border border-gray-300 rounded-lg p-5 flex flex-col items-start">
      <div class="uppercase font-semibold text-gray-900 bg-white -mt-8 -ml-2 px-2">Power management</div>
    </div>
    <div class="border border-gray-300 rounded-lg p-5 flex flex-col items-start">
      <div class="uppercase font-semibold text-gray-900 bg-white -mt-8 -ml-2 px-2">Destruction</div>
      <div class="mt-2">
        Destruction releases all resources attached to this virtual machine, including all storage disks.
      </div>
      <div class="mt-2 flex items-center">
        <button class="bg-red-600 hover:bg-red-700 px-6 py-2 rounded text-white flex items-center"
                @click="destroy">
          <span class="-ml-3 mr-3">
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-5 w-5 animation-pulse"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor"
                 v-if="!destroying">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
            </svg>
            <svg class="animate-spin h-5 w-5"
                 xmlns="http://www.w3.org/2000/svg"
                 fill="none"
                 viewBox="0 0 24 24"
                 v-else>
              <circle class="opacity-25"
                      cx="12"
                      cy="12"
                      r="10"
                      stroke="currentColor"
                      stroke-width="4"></circle>
              <path class="opacity-75"
                    fill="currentColor"
                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
          </span>
          Do it anyway!
        </button>
        <div class="ml-3 text-red-600" v-if="error">{{ error }}</div>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      destroying: false,
      error: null,
    };
  },
  methods: {
    async destroy() {
      if (this.destroying) return;
      this.destroying = true;
      try {
        await this.$axios.$delete("/machines/" + this.$route.params.id);
        this.$router.push("/machines");
      } catch (err) {
        if (err.response && err.response.data) this.error = err.response.data;
        else this.error = err;
      } finally {
        this.destroying = false;
      }
    },
  },
};
</script>