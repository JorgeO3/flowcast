<script lang="ts" setup>
import type { AlbumContent } from "@/utils/sectionTypes.js";

import SectionMenuAlbum from "./SectionMenuAlbum.vue";
import SectionContentCover from "./SectionContentCover.vue";
import SectionContentDetails from "./SectionContentDetails.vue";

interface Props extends AlbumContent {}
defineProps<Props>();

const contentClases = `
  flex flex-col space-y-3 w-48 justify-between ring-offset-background
  rounded-md transition-colors hover:bg-accent hover:text-accent-foreground p-4
`;
</script>

<template>
  <NuxtLink :to="link" :class="contentClases">
    <ContextMenu>
      <!-- Cover -->
      <ContextMenuTrigger>
        <SectionContentCover
          :="cover"
          class="flex overflow-hidden w-fit rounded-md"
        />
      </ContextMenuTrigger>

      <!-- Menu -->
      <SectionMenuAlbum />
    </ContextMenu>

    <!-- Details -->
    <SectionContentDetails :title :text />
  </NuxtLink>
</template>
