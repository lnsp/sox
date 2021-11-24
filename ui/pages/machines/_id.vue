<template>
  <div class="h-full flex flex-col">
    <div class="h-10 flex max-w-xl items-center text-gray-400 mb-4">
      <NuxtLink to="/machines"
                class="flex justify-center items-center group hover:text-oxide-700 border-b border-transparent hover:border-oxide-700">
        <svg viewBox="0 0 10 11"
             class="h-3 mr-4"
             style="fill: currentColor">
          <g transform="matrix(-0.000438321,-0.925173,1,-0.000473772,-155.809,482.889)">
            <path d="M516.305,156.035L521.86,165.665L510.749,165.665L516.305,156.035Z" />
          </g>
        </svg>
        <h1 class="font-mono uppercase">
          Machines
        </h1>
      </NuxtLink>
      <h1 class="ml-3 font-mono uppercase" v-if="details">
        / {{ details.id }}
      </h1>
    </div>
    <div class="border border-groy-500 rounded flex-grow flex flex-col"
         v-if="details">
      <div class="border-b border-groy-500 p-4 flex items-center">
        <div class="relative px-2">
          <div class="absolute h-4 w-4 animate-ping rounded-full"
               :class="[statusColor(details.status)]" />
          <div class="h-4 w-4 rounded-full"
               :class="[statusColor(details.status)]" />
        </div>
        <div class="ml-4 flex-grow">
          <div class="font-normal text-xl">{{ details.name }}</div>
        </div>
        <div v-if="details.networkInterfaces.length > 0"
             class="flex flex-col items-end gap-2 font-mono">
          <div v-if="details.networkInterfaces[0].ipv4">
            <span class="text-gray-300 mr-2 text-xs">{{ details.networkInterfaces[0].ipv4 }}</span>
            <span class="border border-oxide-900 text-oxide-700 rounded text-xs px-2 py-1">IPv4</span>
          </div>
          <div v-if="details.networkInterfaces[0].ipv6">
            <span class="text-gray-300 mr-2 text-xs">{{ details.networkInterfaces[0].ipv6 }}</span>
            <span class="border border-oxide-900 text-oxide-700 rounded text-xs px-2 py-1">IPv6</span>
          </div>
        </div>
      </div>
      <div class="flex flex-grow">
        <div class="w-48 border-r border-groy-500 flex-shrink-0">
          <div class="p-4 flex flex-col gap-2">
            <machine-nav-item :to="'/machines/' + details.id">Dashboard</machine-nav-item>
            <machine-nav-item :to="'/machines/' + details.id + '/networking'">Networking</machine-nav-item>
            <machine-nav-item :to="'/machines/' + details.id + '/storage'">Storage</machine-nav-item>
            <machine-nav-item :to="'/machines/' + details.id + '/management'">Management</machine-nav-item>
          </div>
        </div>
        <div class="flex-grow p-8 overflow-y-scroll">
          <NuxtChild class="h-0" />
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
        STATUS_UNSPECIFIED: "bg-groy-700",
        CREATED: "bg-yellow-500",
        STOPPED: ["border", "border-gray-500"],
        RUNNING: "bg-green-500",
        CRASHED: "bg-red-500",
      }[status];
    },
  },
};
</script>