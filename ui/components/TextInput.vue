<template>
  <input type="text"
         class="w-full max-w-lg h-12 text-gray-200 placeholder-gray-600 px-3 bg-groy-900 border border-groy-500 rounded-sm font-mono focus:border-oxide-700 focus:outline-none"
         :value="value"
         @input="$emit('input', $event.target.value)"
         :placeholder="placeholder">
</template>

<script>
export default {
  props: ["value", "placeholder"],
};
</script>