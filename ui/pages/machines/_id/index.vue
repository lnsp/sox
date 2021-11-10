<template>
  <div class="h-full flex flex-col">
    <div class="flex max-w-xl justify-between items-center">
      <headline h1>Machines</headline>
    </div>
    <divider light />
    <div class="bg-white border-2 border-gray-300 rounded-lg flex-grow flex flex-col">
      <div class="border-b-2 border-gray-300 p-4 flex">
        <div class="relative px-2">
          <div class="absolute h-3 w-3 animate-ping rounded-full"
               :class="[statusColor(details.status)]" />
          <div class="h-3 w-3 rounded-full"
               :class="[statusColor(details.status)]" />
        </div>
        <div>{{ $route.params.id }}</div>
      </div>
      <div class="flex flex-grow">
        <div class="w-48 border-r-2 border-gray-300">
          <div class="p-4">
            <machine-nav-item to="networking">Networking</machine-nav-item>
            <machine-nav-item to="storage">Storage</machine-nav-item>
            <machine-nav-item to="power">Power</machine-nav-item>
          </div>
        </div>
        <div class="flex-grow">

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
    }
  },
  mounted() {
    this.$store.dispatch('api/machineDetails', this.$route.params.id);
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
  }
}
</script>