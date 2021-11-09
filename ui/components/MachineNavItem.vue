<template>
  <nuxt-link class="block" v-slot="{ isActive, navigate, route }" custom :to="to">
    <a class="block rounded hover:bg-gray-300 text-gray-900 py-2 px-3" :class="isActive ? ['bg-gray-300'] : ['bg-none']" @click="navigate" :href="route.fullPath">
      <slot></slot>
    </a>
  </nuxt-link>
</template>

<script>
export default {
  props: ["to"]
}
</script>