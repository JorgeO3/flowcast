<script lang="ts" setup>
defineProps<{ to: string }>();

const linkClass = `
  text-muted-foreground hover:text-foreground 
  transition-colors ease-in-out text-sm font-semibold
  flex items-center gap-1
  
`;
</script>

<template>
  <NuxtLink :to :class="linkClass">
    <slot>
      Login
      <Icon name="lucide:arrow-right" size="22" />
    </slot>
  </NuxtLink>
</template>
