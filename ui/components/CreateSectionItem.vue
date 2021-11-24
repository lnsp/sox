<template>
  <div class="" :class="[disabled ? 'opacity-25' : 'opacity-100']">
    <div class="flex items-center">
      <div class="w-8 text-center mr-4">
        <circled :disabled="disabled">
          {{ index }}
        </circled>
      </div>
      <div>
        <h2 :disabled="disabled">
          {{ title }}
        </h2>
      </div>
    </div>
    <div class="ml-12 mt-4">
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