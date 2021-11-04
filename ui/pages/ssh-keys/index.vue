<template>
  <div>
    <div class="flex max-w-xl justify-between items-center">
      <headline h1>SSH Keys</headline>
      <div>
        <AddButton @click="$router.push('/ssh-keys/new')" />
      </div>
    </div>
    <divider light />
    <div class="max-w-xl">
      <div v-for="key in sshKeys"
           :key="key.id"
           class="flex p-3 items-center justify-between border mb-2 border-gray-200 bg-white rounded hover:border-gray-300">
        <div class="flex-grow px-2">
          <div class="text-gray-900 font-medium">{{ key.name }}</div>
          <div class="text-xs text-gray-700">{{ key.id }}</div>
        </div>
        <div class="font-mono text-sm">
          {{ key.fingerprint }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    sshKeys() {
      return this.$store.state.api.sshKeys;
    },
  },
  methods: {
  },
  mounted() {
    this.$store.dispatch("api/sshKeys");
  },
};
</script>


</style>