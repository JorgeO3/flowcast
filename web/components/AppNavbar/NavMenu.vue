<script lang="ts" setup>
interface Link {
  name: string;
  path: string;
}
defineProps<{ links: Link[] }>();

const linkClass = `
  text-muted-foreground hover:text-foreground
  transition-colors ease-in-out text-sm
`;
</script>

<template>
  <nav>
    <ul
      class="flex items-center max-lg:space-x-4 space-x-6 text-sm font-medium"
    >
      <li v-for="link in links" :key="link.name">
        <NuxtLink
          nuxt-client
          :to="link.path"
          :class="linkClass"
          active-class="text-white"
        >
          {{ link.name }}
        </NuxtLink>
      </li>
    </ul>
  </nav>
</template>
