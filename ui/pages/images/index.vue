<template>
  <div>
    <div class="flex max-w-xl items-center">
      <headline h1>Images</headline>
    </div>
    <divider light />
    <div class="max-w-xl">
      <div v-for="image in images"
           :key="image.id"
           class="flex p-3 items-center justify-between border mb-2 border-gray-200 bg-white rounded hover:border-gray-300">
        <div class="text-gray-700">
          <system-icon :system="image.system" />
        </div>
        <div class="flex-grow px-2">
          <div class="text-gray-900 font-medium">{{ image.name }}</div>
          <div class="text-xs text-gray-700">{{ image.id }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    images() {
      return this.$store.state.api.images;
    },
  },
  methods: {},
  mounted() {
    this.$store.dispatch("api/images");
  },
};
</script>