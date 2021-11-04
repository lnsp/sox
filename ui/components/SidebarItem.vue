<template>
  <nuxt-link v-slot="{ isActive, isExactActive, navigate, route }"
             custom
             :to="to">
    <div :set="active = matches(isActive, isExactActive, route)">
      <a class="rounded-lg flex items-center px-5 py-3 my-2 group hover:bg-gray-800 cursor-pointer"
         @click="navigate"
         :href="route.fullPath"
         :class="[active ? ['bg-gray-700'] : ['bg-gray-900']]">
        <div class="group-hover:text-gray-100 link-icon"
             :class="[active ? ['text-gray-100'] : ['text-gray-400']]">
          <slot></slot>
        </div>
        <div v-if="!collapsed"
             class="link-text ml-3 group-hover:text-gray-100"
             :class="[active ? ['text-gray-100'] : ['text-gray-400']]">{{ name }}</div>
      </a>
    </div>
  </nuxt-link>
</template>

<script>
export default {
  props: ["name", "collapsed", "to"],
  methods: {
    matches(active, exactActive, route) {
      return route.fullPath === "/" ? exactActive : active;
    },
  },
};
</script>