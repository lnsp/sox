<template>
  <li class="rounded-lg flex items-center px-5 py-3 my-2 cursor-pointer group hover:bg-gray-800"
      :class="{ 'bg-gray-700': active }">
    <div class="group-hover:text-gray-100"
         :class="{ 'text-gray-100': active, 'text-gray-400': !active }">
      <slot></slot>
    </div>

    <div v-if="!collapsed" class="ml-2 group-hover:text-gray-100" :class="{ 'text-gray-100': active, 'text-gray-400': !active }">{{ name }}</div>
  </li>
</template>

<script>
export default {
  props: ["name", "active", "collapsed"],
};
</script>