<template>
  <span class="block" :class="[color]">
    <h1 v-if="h1" class="text-3xl font-medium"><slot></slot></h1>
    <h2 v-if="h2" class="text-2xl font-normal"><slot></slot></h2>
    <h3 v-if="h3" class="text-xl font-normal"><slot></slot></h3>
  </span>
</template>

<script>
export default {
  props: {
    h1: Boolean,
    h2: Boolean,
    h3: Boolean,
    light: Boolean,
    disabled: Boolean,
  },
  computed: {
    color() {
      if (this.light) {
        if (this.disabled) {
          return 'text-gray-400';
        } else {
          return 'text-gray-100';
        }
      } else {
        if (this.disabled) {
          return 'text-gray-400';
        } else {
          return 'text-gray-900';
        }
      }
    }
  }
}
</script>