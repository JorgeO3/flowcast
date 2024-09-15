<script lang="ts" setup>

interface Playlist {
  icon: any;
  text: string;
  path: string;
}

defineProps<{
  playlist: Playlist;
}>();

const linkClass = `
flex items-center px-4 py-2 text-sm font-medium 
ring-offset-background rounded-md transition-colors
hover:bg-accent hover:text-accent-foreground h-10
`;
</script>

<template>
  <NuxtLink :to="playlist.path" :class="linkClass" activeClass="bg-secondary">
    <component :is="playlist.icon" class="mr-2 h-4 w-4" />
    {{ playlist.text }}
  </NuxtLink>
</template>