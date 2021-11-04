<template>
  <button class="flex py-2 pl-2 pr-4 items-center rounded-lg bg-gray-800 text-gray-100" @click="$emit('click')">
    <svg xmlns="http://www.w3.org/2000/svg"
         class="h-5 w-5"
         fill="none"
         viewBox="0 0 24 24"
         stroke="currentColor">
      <path stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
    </svg>
    <span class="ml-2">Add</span>
  </button>
</template>