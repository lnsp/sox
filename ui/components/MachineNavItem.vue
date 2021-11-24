<template>
  <nuxt-link class="block" v-slot="{ isExactActive, navigate, route }" custom :to="to">
    <a class="block rounded-sm py-2 px-3 border border-groy-900 hover:border-oxide-700" :class="isExactActive ? ['text-oxide-400', 'bg-oxide-900'] : ['text-gray-400']" @click="navigate" :href="route.fullPath">
      <slot></slot>
    </a>
  </nuxt-link>
</template>

<script>
export default {
  props: ["to"]
}
</script>