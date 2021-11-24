<template>
  <div class="max-w-2xl">
    <div class="h-10 flex justify-between items-center">
      <div class="flex items-center text-gray-400">
        <svg viewBox="0 0 10 11"
             class="h-3 mr-4 transform rotate-90"
             style="fill: currentColor">
          <g transform="matrix(-0.000438321,-0.925173,1,-0.000473772,-155.809,482.889)">
            <path d="M516.305,156.035L521.86,165.665L510.749,165.665L516.305,156.035Z" />
          </g>
        </svg>
        <h1 class="font-mono uppercase">
          Machines
        </h1>
      </div>
      <div>
        <transition name="fade">
          <AddButton @click="$router.push('machines/create')"
                     v-if="machines.length" />
        </transition>
      </div>
    </div>
    <div class="mt-4">
      <div v-for="machine in machines"
           :key="machine.id"
           class="flex p-3 items-center justify-between border mb-2 border-groy-700 rounded cursor-pointer hover:border-groy-400"
           @click="$router.push('/machines/' + machine.id)">
        <div class="relative px-2">
          <div class="absolute h-3 w-3 animate-ping rounded-full"
               :class="[statusColor(machine.status)]" />
          <div class="h-3 w-3 rounded-full"
               :class="[statusColor(machine.status)]" />
        </div>
        <div class="flex-grow px-2 text-gray-300">
          <div class="font-medium">
            {{ machine.name }}
          </div>
          <div class="text-xs font-mono">
            {{ machine.id }}
          </div>
        </div>
      </div>
      <div v-if="!machines.length"
           class="h-full flex flex-col justify-center items-center">
        <div class="text-xl text-center my-8">You have no machines defined.</div>
        <button class="max-w-max flex h-9 px-4 justify-center items-center font-mono rounded-sm border border-oxide-700 hover:bg-oxide-900 text-oxide-400"
                @click="$router.push('machines/create')">
          Create one!
        </button>
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
        STATUS_UNSPECIFIED: "bg-gray-300",
        CREATED: "bg-yellow-500",
        STOPPED: ["border", "border-gray-500"],
        RUNNING: "bg-oxide-400",
        CRASHED: "bg-rod-400",
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
