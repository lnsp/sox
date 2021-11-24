<template>
  <div class="border border-oxide-900 rounded p-5 flex flex-col items-start">
    <div class="uppercase font-mono text-sm text-oxide-700 bg-groy-900 -mt-8 -ml-2 px-2">{{ name }}</div>
    <div class="mt-4 w-full">
      <slot />
    </div>
  </div>
</template>

<script>
export default {
  props: ['name']
}
</script>