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
          Images
        </h1>
      </div>
    </div>
    <div class="mt-4">
      <div v-for="image in images"
           :key="image.id"
           class="flex px-3 h-16 items-center justify-between border mb-2 border-groy-500 rounded hover:border-oxide-700">
        <div class="text-gray-300">
          <system-icon :system="image.system" />
        </div>
        <div class="flex-grow px-2">
          <div class="text-gray-300">{{ image.name }}</div>
          <div class="text-xs text-gray-500">{{ image.id }}</div>
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