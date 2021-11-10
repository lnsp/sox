<template>
  <div>
    <div class="flex max-w-xl justify-between items-center">
      <headline h1>
        Machines
      </headline>
      <div>
        <transition name="fade">
          <AddButton @click="$router.push('machines/create')" />
        </transition>
      </div>
    </div>
    <divider light />
    <div class="max-w-xl">
      <div v-for="machine in machines"
           :key="machine.id"
           class="flex p-3 items-center justify-between border mb-2 border-gray-200 bg-white rounded cursor-pointer hover:border-gray-300"
           @click="$router.push('/machines/' + machine.id)">
        <div class="relative px-2">
          <div class="absolute h-3 w-3 animate-ping rounded-full"
               :class="[statusColor(machine.status)]" />
          <div class="h-3 w-3 rounded-full"
               :class="[statusColor(machine.status)]" />
        </div>
        <div class="flex-grow px-2">
          <div class="text-gray-900 font-medium">
            {{ machine.name }}
          </div>
          <div class="text-xs text-gray-700">
            {{ machine.id }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      machineModal: false,
    };
  },
  computed: {
    machines() {
      return this.$store.state.api.machines;
    },
  },
  mounted() {
    this.$store.dispatch("api/machines");
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
};
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: transform opacity 0.5s;
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
