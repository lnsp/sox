<template>
  <div class="w-full border my-4" :class="[light ? ['border-groy-600'] : ['border-groy-700']]">
  </div>
</template>

<script>
export default {
  props: {
    light: Boolean
  }
}
</script>