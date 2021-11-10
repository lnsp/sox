<template>
  <div class="h-full flex flex-col">
    <div class="flex max-w-xl justify-between items-center">
      <NuxtLink to="/machines" class="flex justify-center items-center group">
        <span class="border border-gray-400 group-hover:border-gray-900 group-hover:text-gray-900 text-gray-400 rounded-full flex justify-center items-center w-8 h-8 mr-4">
          <svg xmlns="http://www.w3.org/2000/svg"
               class="h-4 w-4"
               fill="none"
               viewBox="0 0 24 24"
               stroke="currentColor">
            <path stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M10 19l-7-7m0 0l7-7m-7 7h18" />
          </svg>
        </span>
        <headline h1>Machines</headline>
      </NuxtLink>
    </div>
    <divider light />
    <div class="bg-white border border-gray-300 rounded-lg flex-grow flex flex-col" v-if="details">
      <div class="border-b border-gray-300 p-4 flex items-center">
        <div class="relative px-2">
          <div class="absolute h-4 w-4 animate-ping rounded-full"
               :class="[statusColor(details.status)]" />
          <div class="h-4 w-4 rounded-full"
               :class="[statusColor(details.status)]" />
        </div>
        <div class="ml-4 flex-grow">
          <div class="font-bold text-xl text-gray-900">{{ details.name }}</div>
          <div class="text-sm text-gray-500">{{ details.id }}</div>
        </div>
        <div v-if="details.networkInterfaces.length > 0"
             class="flex flex-col items-end gap-2">
          <div v-if="details.networkInterfaces[0].ipv4">
            <span class="font-mono text-gray-900 mr-2 text-sm">{{ details.networkInterfaces[0].ipv4 }}</span>
            <span class="bg-gray-700 text-gray-100 rounded text-sm px-2 py-1">IPv4</span>
          </div>
          <div v-if="details.networkInterfaces[0].ipv6">
            <span class="font-mono text-gray-900 mr-2 text-sm">{{ details.networkInterfaces[0].ipv6 }}</span>
            <span class="bg-gray-700 text-gray-100 rounded text-sm px-2 py-1">IPv6</span>
          </div>
        </div>
      </div>
      <div class="flex flex-grow">
        <div class="w-48 border-r border-gray-300">
          <div class="p-4 flex flex-col gap-2">
            <machine-nav-item :to="'/machines/' + details.id">Dashboard</machine-nav-item>
            <machine-nav-item :to="'/machines/' + details.id + '/networking'">Networking</machine-nav-item>
            <machine-nav-item :to="'/machines/' + details.id + '/storage'">Storage</machine-nav-item>
            <machine-nav-item :to="'/machines/' + details.id + '/management'">Management</machine-nav-item>
          </div>
        </div>
        <div class="flex-grow p-8">
          <NuxtChild />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    details() {
      return this.$store.state.api.machineDetails[this.$route.params.id];
    },
  },
  mounted() {
    this.$store.dispatch("api/machineDetails", this.$route.params.id);
  },
  methods: {
    statusColor(status) {
      return {
        STATUS_UNSPECIFIED: "bg-gray-500",
        CREATED: "bg-yellow-500",
        STOPPED: "bg-gray-500",
        RUNNING: "bg-green-500",
        DESTROYED: "bg-red-500",
      }[status];
    },
  },
};
</script>