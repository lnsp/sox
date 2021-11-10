<template>
  <div class="relative">
    <headline h1>
      Create a new machine
    </headline>
    <divider light />
    <div class="border border-red-500 text-red-500 rounded-lg flex items-center p-3 bg-white"
         v-if="error">
      <div>
        <svg xmlns="http://www.w3.org/2000/svg"
             class="h-6 w-6"
             fill="none"
             viewBox="0 0 24 24"
             stroke="currentColor">
          <path stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      </div>
      <div class="ml-3">
        {{ error }}
      </div>
    </div>
    <form>
      <create-section-item :index=0
                           title="Choose a name"
                           :disabled="disabled(0)">
        <text-input v-model="machine.name" />
        <div class="text-sm mt-2 text-red-900"
             v-if="machine.name.match(/^[a-zA-Z0-9]+$/g) === null">
          Name must match
          <pre class="inline text-xs bg-gray-200 p-1">/^[a-zA-Z0-9]+$/g</pre>.
        </div>
      </create-section-item>
      <create-section-item :index=1
                           title="Choose the machine size"
                           :disabled="disabled(1)">
        <div class="flex gap-6">
          <form-group>
            <form-label>cpus (cores)</form-label>
            <number-input v-model="machine.specs.cpus"
                          :step="1"
                          :min="1" />
          </form-group>
          <form-group>
            <form-label>memory (mb)</form-label>
            <number-input v-model="machine.specs.memory"
                          :step="512"
                          :min="512" />
          </form-group>
          <form-group>
            <form-label>disk (gb)</form-label>
            <number-input v-model="machine.specs.disk"
                          :min="10" />
          </form-group>
        </div>
      </create-section-item>
      <create-section-item :index=2
                           title="Pick an image"
                           :disabled="disabled(2)">
        <div class="max-w-lg gap-4">
          <div v-for="image in images"
               :key="image.id"
               @click="machine.imageId = image.id"
               class="flex p-3 items-center justify-between border mb-2 border-gray-200 bg-white rounded hover:border-gray-300">
            <div class="text-gray-700">
              <system-icon :system="image.system" />
            </div>
            <div class="flex-grow px-2">
              <div class="text-gray-900 font-medium">{{ image.name }}</div>
            </div>
            <div v-if="machine.imageId === image.id">
              <svg xmlns="http://www.w3.org/2000/svg"
                   class="h-6 w-6"
                   fill="none"
                   viewBox="0 0 24 24"
                   stroke="currentColor">
                <path stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M5 13l4 4L19 7" />
              </svg>
            </div>
          </div>
        </div>
      </create-section-item>
      <create-section-item :index=3
                           title="Attach to networks"
                           :disabled="disabled(3)">
        <div v-for="network in networks"
             :key="network.id"
             @click="toggleNetwork(network.id)"
             class="flex p-3 max-w-lg items-center justify-between border mb-2 border-gray-200 bg-white rounded hover:border-gray-300">
          <div class="flex-grow px-2">
            <div class="text-gray-900 font-medium">{{ network.name }}</div>
            <div class="text-xs text-gray-700 font-mono">{{ network.ipv4 }}</div>
          </div>
          <div v-if="machine.networkIds.includes(network.id)">
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-6 w-6"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 13l4 4L19 7" />
            </svg>
          </div>
        </div>
      </create-section-item>
      <create-section-item :index=4
                           title="Pick an SSH key"
                           :disabled="disabled(4)">
        <div v-for="key in sshKeys"
             :key="key.id"
             @click="toggleSSHKey(key.id)"
             class="flex p-3 max-w-lg items-center justify-between border mb-2 border-gray-200 bg-white rounded hover:border-gray-300">
          <div class="flex-grow px-2">
            <div class="text-gray-900 font-medium">{{ key.name }}</div>
            <div class="text-xs text-gray-700 font-mono">{{ key.fingerprint }}</div>
          </div>
          <div v-if="machine.sshKeyIds.includes(key.id)">
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-6 w-6"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 13l4 4L19 7" />
            </svg>
          </div>
        </div>
      </create-section-item>
      <create-section-item :index=5
                           title="Verify settings and confirm"
                           :disabled="disabled(5)">
        <button @click.prevent="createMachine"
                class="mt-2 bg-gray-700 font-medium flex items-center text-gray-100 px-8 py-3 rounded-lg shadow-lg hover:bg-gray-900">
          <span class="-ml-2 mr-3 text-gray-100">
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-5 w-5"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor"
                 v-if="!creating && !id">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M14 5l7 7m0 0l-7 7m7-7H3" />
            </svg>
            <svg class="animate-spin h-5 w-5"
                 xmlns="http://www.w3.org/2000/svg"
                 fill="none"
                 viewBox="0 0 24 24"
                 v-if="creating">
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
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-5 w-5"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor"
                 v-if="!creating && id">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 13l4 4L19 7" />
            </svg>

          </span>

          <transition name="fade">
            <span v-if="!creating && !id">Looks good!</span>
            <span v-if="creating">Creating machine</span>
            <span v-if="!creating && id">Done!</span>
          </transition>
        </button>
      </create-section-item>
    </form>
  </div>
</template>

<script>
import CreateSectionItem from "../../components/CreateSectionItem.vue";
export default {
  components: { CreateSectionItem },
  computed: {
    images() {
      return this.$store.state.api.images;
    },
    sshKeys() {
      return this.$store.state.api.sshKeys;
    },
    networks() {
      return this.$store.state.api.networks;
    },
  },
  data() {
    return {
      machine: {
        name: "",
        specs: {
          cpus: 2,
          memory: 2048,
          disk: 10,
        },
        imageId: "",
        networkIds: [],
        sshKeyIds: [],
      },
      error: "",
      creating: false,
      id: null,
    };
  },
  mounted() {
    this.$store.dispatch("api/sshKeys");
    this.$store.dispatch("api/images");
    this.$store.dispatch("api/networks");
  },
  methods: {
    async createMachine() {
      if (this.id !== null) return;
      this.creating = true;
      try {
        let response = await this.$axios.$post("/machines", this.machine);
        this.id = response.id;
      } catch (err) {
        this.error = err.response.data;
      } finally {
        this.creating = false;
      }
    },
    disabled(stage) {
      switch (stage) {
        case 0:
          return this.id !== null;
        case 1:
          return this.id !== null || this.machine.name.match(/^[a-zA-Z0-9]+$/g) === null;
        case 2:
          return this.id !== null || this.machine.name.match(/^[a-zA-Z0-9]+$/g) === null;
        case 3:
          return (
            this.id !== null ||
            this.machine.name.match(/^[a-zA-Z0-9]+$/g) === null ||
            this.imageId === ""
          );
        case 4:
          return (
            this.id !== null ||
            this.machine.name.match(/^[a-zA-Z0-9]+$/g) === null ||
            this.machine.imageId === "" ||
            this.machine.networkIds.length == 0
          );
        case 5:
          return (
            this.machine.name.match(/^[a-zA-Z0-9]+$/g) === null ||
            this.machine.imageId === "" ||
            this.machine.networkIds.length === 0 ||
            this.machine.sshKeyIds.length === 0
          );
      }
    },
    toggleSSHKey(keyId) {
      if (this.machine.sshKeyIds.includes(keyId)) {
        this.machine.sshKeyIds = this.machine.sshKeyIds.filter(
          (e) => e != keyId
        );
      } else {
        this.machine.sshKeyIds.push(keyId);
      }
    },
    toggleNetwork(networkId) {
      if (this.machine.networkIds.includes(networkId)) {
        this.machine.networkIds = this.machine.networkIds.filter(
          (e) => e != networkId
        );
      } else {
        this.machine.networkIds.push(networkId);
      }
    },
  },
};
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>