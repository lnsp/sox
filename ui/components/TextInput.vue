<template>
  <input type="text"
         class="min-w-xl px-3 py-2 bg-white border border-gray-300 rounded focus:border-gray-400 focus:outline-none"
         :value="value"
         @input="$emit('input', $event.target.value)"
         :placeholder="placeholder">
</template>

<script>
export default {
  props: ["value", "placeholder"],
};
</script>