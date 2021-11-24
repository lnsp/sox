<template>
  <div class="flex flex-col gap-8">
        <spinner class="text-red-600 h-5 w-5" />
    <subgroup name="Power cycle">
      <button v-if="details.status != 'RUNNING'"
              class="flex items-center font-mono text-green-400 border-b border-transparent hover:border-green-400" @click="trigger('POWERON')">
        <span class="w-5 mr-2" v-if="!triggering">></span>
        <spinner class="text-green-400 w-5 mr-2" v-else />
        Power on
      </button>
      <button v-else
              class="flex items-center font-mono border-b border-transparent hover:border-yellow-400 text-yellow-400" @click="trigger('POWEROFF')">
        <span class="w-5 mr-2" v-if="!triggering">></span>
        <spinner class="text-yellow-400 w-5 mr-2" v-else />
        Shutdown
      </button>
    </subgroup>
    <subgroup name="Destruction">
      <div class="">
        Destruction releases all resources attached to this virtual machine, including all storage disks.
      </div>
      <div class="mt-4">
        <button class="text-red-600 font-mono border-b border-transparent hover:border-red-600 flex items-center"
                @click="destroy" v-if="!destroyed">
          <span class="w-5 mr-2" v-if="!destroying">></span>
          <spinner class="text-red-600 w-5 mr-2" v-else />
          Do it anyway!
        </button>
      </div>
    </subgroup>
  </div>
</template>

<script>
export default {
  computed: {
    details() {
      return this.$store.state.api.machineDetails[this.$route.params.id];
    },
  },
  data() {
    return {
      destroying: false,
      destroyed: false,
      triggering: null,
    };
  },
  methods: {
    async trigger(event) {
      if (this.triggering) return;
      this.triggering = true;
      try {
        await this.$axios.$post("/machines/" + this.$route.params.id + "/trigger?event=" + event)
        this.$store.dispatch("api/machineDetails", this.$route.params.id);
      } catch (err) {
        this.$store.commit('local/error', err)
      } finally {
        this.triggering = false;
      }
    },
    async destroy() {
      if (this.destroying) return;
      this.destroying = true;
      try {
        await this.$axios.$delete("/machines/" + this.$route.params.id);
        this.destroyed = true;
        this.$router.push("/machines");
      } catch (err) {
        this.$store.commit('local/error', err)
      } finally {
        this.destroying = false;
      }
    },
  },
};
</script>