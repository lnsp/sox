<template>
  <nuxt-link v-slot="{ isActive, isExactActive, navigate, route }"
             custom
             :to="to">
    <div :set="active = matches(isActive, isExactActive, route)">
      <a class="rounded-sm flex items-center px-5 py-2 my-2 group border border-transparent hover:border-oxide-700 cursor-pointer"
         @click="navigate"
         :href="route.fullPath"
         :class="[active ? ['bg-oxide-900'] : []]">
        <div class="link-icon h-5"
             :class="[active ? ['text-oxide-400'] : ['text-gray-400']]">
          <slot></slot>
        </div>
        <div v-if="!collapsed"
             class="link-text ml-3"
             :class="[active ? ['text-oxide-400'] : ['text-gray-400']]">{{ name }}</div>
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