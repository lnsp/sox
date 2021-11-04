<template>
  <span :class="[light ? 'text-gray-100' : 'text-gray-900']">
    <h1 v-if="h1" class="text-3xl font-medium"><slot></slot></h1>
    <h2 v-if="h2" class="text-1xl font-medium"><slot></slot></h2>
  </span>
</template>

<script>
export default {
  props: {
    h1: Boolean,
    h2: Boolean,
    light: Boolean,
  }
}
</script>