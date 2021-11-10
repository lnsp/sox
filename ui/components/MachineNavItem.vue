<template>
  <nuxt-link class="block" v-slot="{ isExactActive, navigate, route }" custom :to="to">
    <a class="block rounded hover:bg-gray-200 text-gray-900 py-2 px-3 border border-white hover:border-gray-300" :class="isExactActive ? ['bg-gray-200'] : ['bg-none']" @click="navigate" :href="route.fullPath">
      <slot></slot>
    </a>
  </nuxt-link>
</template>

<script>
export default {
  props: ["to"]
}
</script>