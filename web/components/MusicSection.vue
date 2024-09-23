<script lang="ts" setup>
import { Separator } from "@/components/ui/separator";
import type { Section, MediaContent } from "@/utils/sectionTypes";

import SectionHeader from "./MusicSection/SectionHeader.vue";
import SectionActContent from "./MusicSection/SectionActContent.vue";
import SectionSongContent from "./MusicSection/SectionSongContent.vue";
import SectionAlbumContent from "./MusicSection/SectionAlbumContent.vue";
import SectionRadioContent from "./MusicSection/SectionRadioContent.vue";
import SectionPodcastContent from "./MusicSection/SectionPodcastContent.vue";
import SectionPlaylistContent from "./MusicSection/SectionPlaylistContent.vue";

interface Props extends Section {}
defineProps<Props>();

const contentMap: Record<MediaContent["contentType"], Component> = {
  act: SectionActContent,
  song: SectionSongContent,
  album: SectionAlbumContent,
  radio: SectionRadioContent,
  podcast: SectionPodcastContent,
  playlist: SectionPlaylistContent,
};
</script>

<template>
  <div class="flex flex-col w-full h-fit mb-16">
    <!-- Header -->
    <SectionHeader :title :text :link />

    <Separator class="my-4" />

    <!-- Content -->
    <template v-for="item in content" :key="item.id">
      <component :is="contentMap[item.contentType]" :="item" />
    </template>
  </div>
</template>
