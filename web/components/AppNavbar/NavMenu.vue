<script lang="ts" setup>
interface Link {
  name: string;
  path: string;
}

defineProps<{ links: Link[] }>();

const linkClass = `
  text-muted-foreground hover:text-foreground 
  transition-colors ease-in-out text-sm font-semibold
`;
</script>

<template>
  <nav>
    <ul class="flex gap-x-4">
      <li v-for="link in links" :key="link.name">
        <NuxtLink :to="link.path" :class="linkClass">
          {{ link.name }}
        </NuxtLink>
      </li>
    </ul>
  </nav>
</template>