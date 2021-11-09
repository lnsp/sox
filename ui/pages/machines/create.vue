<template>
  <div>
    <headline h1>
      Create a new machine
    </headline>
    <divider light />
    <form>
      <create-section-item :index=0
                           title="Choose a name"
                           :disabled="disabled(0)">
        <text-input v-model="name" />
        <div class="text-sm mt-2 text-red-900"
             v-if="name.match(/^[a-zA-Z0-9]+$/g) === null">
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
            <number-input v-model="cpus"
                          :step="1"
                          :min="1" />
          </form-group>
          <form-group>
            <form-label>memory (mb)</form-label>
            <number-input v-model="memory"
                          :step="512"
                          :min="512" />
          </form-group>
          <form-group>
            <form-label>disk (gb)</form-label>
            <number-input v-model="disk"
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
               @click="imageId = image.id"
               class="flex p-3 items-center justify-between border mb-2 border-gray-200 bg-white rounded hover:border-gray-300">
            <div class="text-gray-700">
              <system-icon :system="image.system" />
            </div>
            <div class="flex-grow px-2">
              <div class="text-gray-900 font-medium">{{ image.name }}</div>
            </div>
            <div v-if="image.id === imageId">
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
            <div v-if="networkIds.includes(network.id)">
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
            <div v-if="sshKeyIds.includes(key.id)">
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
      <create-section-item :index=5 title="Verify settings and confirm" :disabled="disabled(5)">
        <button @click.prevent="createMachine" class="mt-2 bg-gray-700 font-medium text-gray-100 px-8 py-3 rounded-lg shadow-lg hover:bg-gray-900">Looks good!</button>
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
      cpus: 2,
      memory: 2048,
      disk: 10,
      imageId: '',
      networkIds: [],
      sshKeyIds: [],
      name: '', 
    };
  },
  mounted() {
    this.$store.dispatch("api/sshKeys");
    this.$store.dispatch("api/images");
    this.$store.dispatch("api/networks");
  },
  methods: {
    async createMachine() {
      let response = await this.$axios.$post('/machines', {
        specs: {
          cpus: this.cpus,
          memory: this.memory,
          disk: this.disk,
        },
        image: this.imageId,
        networks: this.networkIds,
        sshKeys: this.sshKeys,
        name: this.name,
      })
      this.$route.push('/machines')
    },
    disabled(stage) {
      switch (stage) {
        case 0:
          return false;
        case 1:
          return this.name.match(/^[a-zA-Z0-9]+$/g) === null;
        case 2:
          return this.name.match(/^[a-zA-Z0-9]+$/g) === null;
        case 3:
          return (
            this.name.match(/^[a-zA-Z0-9]+$/g) === null || this.imageId === ""
          );
        case 4:
          return (
            this.name.match(/^[a-zA-Z0-9]+$/g) === null ||
            this.imageId === "" ||
            this.networkIds.length == 0
          );
        case 5:
          return (
            this.name.match(/^[a-zA-Z0-9]+$/g) === null ||
            this.imageId === "" ||
            this.networkIds.length === 0 ||
            this.sshKeyIds.length === 0
          );
      }
    },
    toggleSSHKey(keyId) {
      if (this.sshKeyIds.includes(keyId)) {
        this.sshKeyIds = this.sshKeyIds.filter(e => e != keyId);
      } else {
        this.sshKeyIds.push(keyId);
      }
    },
    toggleNetwork(networkId) {
      if (this.networkIds.includes(networkId)) {
        this.networkIds = this.networkIds.filter(e => e != networkId);
      } else {
        this.networkIds.push(networkId);
      }
    }
  },
};
</script>
