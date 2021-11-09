<template>
  <span :class="[color]" class="text-lg inline-block w-8 h-8 text-gray-100 flex items-center justify-center font-bold rounded-full">
    <slot />
  </span>
</template>

<script>
export default {
  props: {
    disabled: Boolean,
  },
  computed: {
    color() {
      if (this.disabled) return 'bg-gray-500';
      else return 'bg-gray-700';
    }
  }
}
</script>