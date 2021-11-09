<template>
  <div class="h-10 w-36 flex">
    <div class="flex flex-row h-10 w-full rounded-lg relative bg-transparent mt-1 overflow-hidden rounded border border-gray-300 focus-within:border-gray-500">
      <button @click.prevent="set(value - step)" :disabled="value <= min"
              class="bg-gray-100 h-full w-20 focus:outline-none border-r border-gray-300"
              :class="value <= min ? ['cursor-not-allowed', 'text-gray-400'] : ['cursor-pointer','text-gray-900', 'hover:text-gray-700', 'hover:bg-gray-300', 'focus:bg-gray-300' ]">
        <span class="m-auto text-2xl font-thin">−</span>
      </button>
      <input type="number"
             class="focus:outline-none text-center w-full bg-white font-normal hover:text-black focus:text-black flex flex-grow items-center text-gray-700"
             name="custom-input-number"
             :value="value"
             @change="set($event.target.value)">
      <button @click.prevent="set(value + step)"
              :disabled="value >= max"
              class="bg-gray-100 h-full w-20 focus:outline-none border-l border-gray-300"
              :class="value >= max ? ['cursor-not-allowed', 'text-gray-400'] : ['cursor-pointer','text-gray-900', 'hover:text-gray-700', 'hover:bg-gray-300', 'focus:bg-gray-300' ]">
        <span class="m-auto text-2xl font-thin">+</span>
      </button>
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