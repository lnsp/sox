<template>
  <div class="my-4">
    <div class="flex items-center">
      <div class="w-8 text-center mr-4">
        <circled :disabled="disabled">
          {{ index }}
        </circled>
      </div>
      <div>
        <headline h2 :disabled="disabled">
          {{ title }}
        </headline>
      </div>
    </div>
    <div class="ml-12 mt-4" v-if="!disabled">
      <slot />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    index: Number,
    title: String,
    disabled: Boolean,
  }
}
</script>