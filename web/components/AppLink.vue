<script lang="ts" setup>
const props = defineProps<{
  to: string;
  classes?: string;
}>();

const linkClass = "text-blue-500 hover:underline" + (props.classes ? props.classes : "");
</script>

<template>
  <NuxtLink :to="to" :class="linkClass">
    <slot />
  </NuxtLink>
</template>


<style></style>