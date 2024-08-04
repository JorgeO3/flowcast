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
    <ul class="flex items-center max-lg:space-x-4 space-x-6 text-sm font-medium">
      <li v-for="link in links" :key="link.name">
        <NuxtLink :to="link.path" class="transition-colors hover:text-foreground/80 text-foreground/60"
          active-class="transition-colors hover:text-foreground/80 text-foreground/60 font-semibold !text-foreground">
          {{ link.name }}
        </NuxtLink>
      </li>
    </ul>
  </nav>
</template>