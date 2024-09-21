<script lang="ts" setup>
defineProps<{ title: string; text: string; link: string }>();
</script>

<template>
  <div class="flex items-center justify-between">
    <div class="space-y-1">
      <h2 class="text-2xl font-semibold tracking-tight">{{ title }}</h2>
      <p class="text-sm text-muted-foreground">
        {{ text }}
      </p>
    </div>

    <NuxtLink
      :to="link"
      class="text-sm font-medium hover:text-foreground text-muted-foreground"
    >
      All
    </NuxtLink>
  </div>
</template>
