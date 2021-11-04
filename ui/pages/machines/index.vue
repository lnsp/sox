<template>
  <div>
    <h1 class="text-3xl font-medium text-gray-900">Machines</h1>
    <divider light />
    <div class="max-w-xl">
      <div v-for="machine in machines"
           :key="machine.id"
           class="flex p-3 items-center justify-between border mb-2 border-gray-200 rounded hover:border-gray-900">
        <div class="relative px-2">
          <div class="absolute h-3 w-3 animate-ping rounded-full" :class="[statusColor(machine.status)]"></div>
          <div class="h-3 w-3 rounded-full" :class="[statusColor(machine.status)]"></div>
        </div>
        <div class="flex-grow px-2">
          <div class="text-gray-900 font-medium">{{ machine.name }}</div>
          <div class="text-xs text-gray-700">{{ machine.id }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    machines() {
      return this.$store.state.api.machines;
    },
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
  mounted() {
    this.$store.dispatch("api/machines");
  },
};
</script>


</style>