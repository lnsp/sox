<template>
  <div class="relative">
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
      <h1 class="ml-3 font-mono uppercase">
        / New
      </h1>
    </div>
    <div class="border border-rod-400 text-rod-400 bg-rod-900 rounded p-5 mb-4"
         v-if="error">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <div class="inline-block rounded-full text-groy-900 bg-rod-400 p-1">
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-3 w-3"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12" />
            </svg>
          </div>
          <span class="ml-3 font-bold">Error</span>
        </div>
        <button class=""
                @click="error = null">
          <svg xmlns="http://www.w3.org/2000/svg"
               class="h-5"
               fill="none"
               viewBox="0 0 24 24"
               stroke="currentColor">
            <path stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      <div class="ml-8 mt-2">
        {{ error }}
      </div>
    </div>
    <form class="flex flex-col gap-6">
      <create-section-item :index=0
                           title="Choose a name"
                           :disabled="disabled(0)">
        <text-input v-model="machine.name" />
        <div class="text-sm mt-2 text-red-500 font-mono"
             v-if="machine.name.match(/^[a-zA-Z0-9]+$/g) === null">
          Name must match
          <pre class="inline text-xs p-1">/^[a-zA-Z0-9]+$/g</pre>
        </div>
      </create-section-item>
      <create-section-item :index=1
                           title="Choose the machine size"
                           :disabled="disabled(1)">
        <div class="flex gap-6 max-w-lg">
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
               class="flex px-3 h-16 items-center justify-between border mb-2 border-groy-500 bg-groy-900 rounded hover:border-oxide-700">
            <div :class="[ machine.imageId === image.id ? ['text-oxide-400'] : 'text-gray-300']">
              <system-icon :system="image.system" />
            </div>
            <div class="flex-grow px-2">
              <div class="text-gray-300 font-mono">{{ image.name }}</div>
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
             class="flex px-3 h-16 max-w-lg items-center justify-between border mb-2 border-groy-500 rounded hover:border-oxide-700">
          <div class="flex-grow px-2">
            <div class="text-gray-300">{{ network.name }}</div>
            <div class="text-xs text-gray-500 font-mono">{{ network.ipv4 }}</div>
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
             class="flex px-3 h-16 max-w-lg items-center justify-between border mb-2 border-groy-500 rounded hover:border-oxide-700">
          <div class="flex-grow px-2">
            <div class="text-gray-300">{{ key.name }}</div>
            <div class="text-xs text-gray-500 font-mono">{{ key.fingerprint }}</div>
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
                class="mt-2 font-mono flex items-center text-oxide-400 border-b border-transparent hover:border-oxide-400">
          <span class="mr-3">
            <span v-if="!creating && !id">
              >
            </span>
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
export default {
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
      error: null,
      creating: null,
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
      if (this.id !== null) {
        this.$router.push("/machines");
        return;
      }
      this.creating = true;
      this.error = null;
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
          return this.creating || this.id !== null;
        case 1:
          return (
            this.creating ||
            this.id !== null ||
            this.machine.name.match(/^[a-zA-Z0-9]+$/g) === null
          );
        case 2:
          return (
            this.creating ||
            this.id !== null ||
            this.machine.name.match(/^[a-zA-Z0-9]+$/g) === null
          );
        case 3:
          return (
            this.creating ||
            this.id !== null ||
            this.machine.name.match(/^[a-zA-Z0-9]+$/g) === null ||
            this.imageId === ""
          );
        case 4:
          return (
            this.creating ||
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