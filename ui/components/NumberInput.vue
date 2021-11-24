<template>
  <div class="h-10 w-full flex text-gray-900">
    <div class="flex flex-row h-12 font-mono w-full rounded-sm border border-groy-500 focus-within:border-oxide-700 relative mt-1 overflow-hidden rounded">
      <input type="number"
             class="w-full bg-groy-900 text-gray-300 px-3 focus:outline-none overflow-hidden"
             name="custom-input-number"
             :value="value"
             @change="set($event.target.value)">
      <div class="flex flex-col w-8">
        <button @click.prevent="set(value + step)"
                :disabled="value >= max"
                class="h-1/2 w-full flex justify-center items-center focus:outline-none border-l border-b border-transparent"
                :class="value >= max ? ['cursor-not-allowed', 'text-gray-700'] : ['cursor-pointer','text-gray-400', 'hover:text-oxide-400', 'focus:border-oxide-700' ]">
          <svg xmlns="http://www.w3.org/2000/svg"
               class="w-4"
               fill="none"
               viewBox="0 0 24 24"
               stroke="currentColor">
            <path stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 15l7-7 7 7" />
          </svg>
        </button>
        <button @click.prevent="set(value - step)"
                :disabled="value <= min"
                class="h-1/2 w-full flex justify-center items-center focus:outline-none border border-transparent"
                :class="value <= min ? ['cursor-not-allowed', 'text-gray-700'] : ['cursor-pointer','text-gray-400', 'hover:text-oxide-400', 'foxus:border-oxide-700' ]">
          <svg xmlns="http://www.w3.org/2000/svg"
               class="w-4 transform rotate-180"
               fill="none"
               viewBox="0 0 24 24"
               stroke="currentColor">
            <path stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 15l7-7 7 7" />
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<style>
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>

<script>
export default {
  props: {
    value: Number,
    min: {
      type: Number,
      default: Number.MIN_VALUE,
    },
    max: {
      type: Number,
      default: Number.MAX_VALUE,
    },
    step: {
      type: Number,
      default: 1,
    },
  },
  methods: {
    set(value) {
      this.$emit("input", Math.min(Math.max(value, this.min), this.max));
    },
  },
};
</script>