<template>
  <button class="flex h-9 pl-2 pr-4 items-center font-mono rounded-sm border border-oxide-700 hover:bg-oxide-900 text-oxide-400" @click="$emit('click')">
    <svg xmlns="http://www.w3.org/2000/svg"
         class="h-4"
         fill="none"
         viewBox="0 0 24 24"
         stroke="currentColor">
      <path stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
    </svg>
    <span class="ml-2">add</span>
  </button>
</template>