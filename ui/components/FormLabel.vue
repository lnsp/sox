<template>
  <label :for="forInput" class="uppercase text-sm font-mono block mb-2 text-gray-500"><slot /></label>
</template>

<script>
export default {
  props: ['forInput']
}
</script>