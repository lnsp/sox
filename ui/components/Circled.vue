<template>
  <span :class="[color]" class="inline-block w-8 h-8 text-gray-100 flex items-center justify-center font-mono border">
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
      if (this.disabled) return ['border-oxide-900', 'text-oxide-900'];
      else return ['border-oxide-700', 'text-oxide-400'];
    }
  }
}
</script>