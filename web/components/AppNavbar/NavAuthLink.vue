<script lang="ts" setup>
import { UserIcon } from "lucide-vue-next";
defineProps<{ to: string }>();

const linkClass = `
  text-muted-foreground hover:text-foreground
  transition-colors ease-in-out text-sm font-semibold flex items-center
`;
</script>

<template>
  <NuxtLink :to class="flex font-semibold text-sm">
    <slot>
      <div :class="linkClass"><UserIcon class="w-4 h-4 mr-2" /> Login</div>
    </slot>
  </NuxtLink>
</template>
