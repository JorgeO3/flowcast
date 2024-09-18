<script lang="ts" setup>
import {
  PlusCircleIcon,
  ListPlusIcon,
  RadioIcon,
  PlusIcon,
  ListMusicIcon,
} from "lucide-vue-next";

interface Playlist {
  name: string;
  id: string;
}

interface Props {
  playlists: Playlist[];
}

defineProps<Props>();
</script>

<template>
  <ContextMenuContent class="w-44">
    <ContextMenuItem>
      <PlusCircleIcon class="mr-2 h-4 w-4" />
      Add to Library
    </ContextMenuItem>
    <ContextMenuItem>
      <ListPlusIcon class="mr-2 h-4 w-4" />
      Add to Queue
    </ContextMenuItem>
    <ContextMenuItem>
      <RadioIcon class="mr-2 h-4 w-4" />
      Go to artist radio
    </ContextMenuItem>
    <ContextMenuSeparator />
    <ContextMenuSub>
      <ContextMenuSubTrigger>
        <PlusIcon class="mr-2 h-4 w-4" />
        Add to Playlist
      </ContextMenuSubTrigger>
      <ContextMenuSubContent>
        <ContextMenuItem>
          <PlusCircleIcon class="mr-2 h-4 w-4" />
          New Playlist
        </ContextMenuItem>
        <ContextMenuSeparator />
        <ContextMenuItem v-for="{ id, name } in playlists" :key="id">
          <ListMusicIcon class="mr-2 h-4 w-4" />
          {{ name }}
        </ContextMenuItem>
      </ContextMenuSubContent>
    </ContextMenuSub>
    <ContextMenuSeparator />
    <ContextMenuItem>Like</ContextMenuItem>
    <ContextMenuItem>Share</ContextMenuItem>
  </ContextMenuContent>
</template>
