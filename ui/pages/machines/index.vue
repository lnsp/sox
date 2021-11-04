<template>
  <div>
    <div class="flex max-w-xl justify-between items-center">
      <headline h1>Machines</headline>
      <div>
        <AddButton @click="machineModal = !machineModal" />
      </div>
    </div>
    <transition name="fade">
      <div class="max-w-xl bg-gray-800 overflow-hidden transition rounded-lg shadow-xl my-4 p-4"
           v-if="machineModal">
        <headline h2 light>Create a new machine</headline>
        <divider />
      </div>
    </transition>
    <divider light />
    <div class="max-w-xl">
      <div v-for="machine in machines"
           :key="machine.id"
           class="flex p-3 items-center justify-between border mb-2 border-gray-200 bg-white rounded hover:border-gray-300">
        <div class="relative px-2">
          <div class="absolute h-3 w-3 animate-ping rounded-full"
               :class="[statusColor(machine.status)]"></div>
          <div class="h-3 w-3 rounded-full"
               :class="[statusColor(machine.status)]"></div>
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
  data() {
    return {
      machineModal: false,
    };
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
    showMachineModal() {
      this.machineModal = true;
    },
    hideMachineModal() {
      this.machineModal = false;
    },
  },
  mounted() {
    this.$store.dispatch("api/machines");
  },
};
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: transform opacity 0.25s;
  transform: translateY(0);
}
.fade-enter /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
  transform: translateY(12);
}
.fade-leave-to {
  opacity: 0;
  transform: translateY(-12);
}
</style>