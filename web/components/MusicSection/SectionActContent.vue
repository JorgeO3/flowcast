<script lang="ts" setup>
import type { ActContent } from "@/utils/sectionTypes.js";

import SectionMenuAct from "./SectionMenuAlbum.vue";
import SectionContentCover from "./SectionContentCover.vue";
import SectionContentDetails from "./SectionContentDetails.vue";

interface Props extends ActContent {}
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
          class="flex overflow-hidden w-fit rounded-full"
        />
      </ContextMenuTrigger>

      <!-- Menu -->
      <SectionMenuAct />
    </ContextMenu>

    <!-- Details -->
    <SectionContentDetails :title :text />
  </NuxtLink>
</template>
