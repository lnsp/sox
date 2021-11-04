<template>
  <div class="w-full border my-4" :class="[light ? ['border-gray-200'] : ['border-gray-700']]">
  </div>
</template>

<script>
export default {
  props: {
    light: Boolean
  }
}
</script>