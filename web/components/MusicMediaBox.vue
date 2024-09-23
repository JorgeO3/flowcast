<script lang="ts" setup>
import type { MediaContent } from "@/utils/tabAllSections.js";
import BoxCover from "./MusicMediaBox/BoxCover.vue";
import BoxMenuAct from "./MusicMediaBox/BoxMenuAct.vue";
import BoxDetails from "./MusicMediaBox/BoxDetails.vue";
import BoxMenuAlbum from "./MusicMediaBox/BoxMenuAlbum.vue";

type Props = MediaContent;
const { type, link } = defineProps<Props>();

const boxClasses = `
flex flex-col space-y-3 w-48 justify-between ring-offset-background
rounded-md transition-colors hover:bg-accent hover:text-accent-foreground p-4
`;
</script>

<template>
  <NuxtLink :to="link" :class="boxClasses">
    <ContextMenu>
      <!-- Cover -->
      <ContextMenuTrigger>
        <BoxCover :="cover" />
      </ContextMenuTrigger>

      <!-- Menu -->
      <BoxMenuAlbum v-if="playlists && type === 'Album'" :playlists />

      <BoxMenuAct v-else />
    </ContextMenu>

    <!-- Details -->
    <BoxDetails :title :text />
  </NuxtLink>
</template>
