<template>
  <div class="flex flex-col gap-8">
    <subgroup name="Power cycle">
      <button v-if="!shutdown"
              class="font-mono text-green-400 border-b border-transparent hover:border-green-400">
        <span class="mr-3">></span>Power on
      </button>
      <button v-else
              class="font-mono border-b border-transparent hover:border-yellow-400 text-yellow-400">
        <span class="mr-3">></span>Shutdown
      </button>
    </subgroup>
    <subgroup name="Destruction">
      <div class="">
        Destruction releases all resources attached to this virtual machine, including all storage disks.
      </div>
      <div class="mt-4">
        <button class="text-red-600 font-mono border-b border-transparent hover:border-red-600 flex items-center"
                @click="destroy"
                v-if="!destroying && !destroyed">
          <span class="mr-3">></span>
          Do it anyway!
        </button>
        <svg class="text-red-600 animate-spin h-5 w-5"
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
        <div class="ml-3 text-rod-400"
             v-if="error">{{ error }}</div>
      </div>
    </subgroup>
  </div>
</template>

<script>
export default {
  data() {
    return {
      shutdown: true,
      destroying: false,
      destroyed: false,
      error: null,
    };
  },
  methods: {
    async destroy() {
      if (this.destroying) return;
      this.destroying = true;
      try {
        await this.$axios.$delete("/machines/" + this.$route.params.id);
        this.destroyed = true;
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