<template>
  <li>
    <div class="uppercase text-xs font-bold text-gray-400 pl-5 py-2"><span v-if="collapsed">{{ shortname }}</span><span v-else>{{ name }}</span></div>
    <ul>
      <slot></slot>
    </ul>
  </li>
</template>

<script>
export default {
  props: ['name', 'shortname', 'collapsed']
}
</script>