<script lang="ts" setup>
import type { Tab } from "@/utils/tabTypes.js";

import TabSongs from "@/components/BrowseTabs/TabSongs.vue";
import TabAlbums from "@/components/BrowseTabs/TabAlbums.vue";
import TabArtists from "@/components/BrowseTabs/TabArtists.vue";
import TabOverview from "@/components/BrowseTabs/TabOverview.vue";
import TabPodcasts from "@/components/BrowseTabs/TabPodcasts.vue";

definePageMeta({ layout: "play" });

const tabs: Tab[] = [
  { name: "All", value: "all", component: TabOverview },
  { name: "Songs", value: "music", component: TabSongs },
  { name: "Albums", value: "albums", component: TabAlbums },
  { name: "Artists", value: "artists", component: TabArtists },
  { name: "Podcasts", value: "podcasts", component: TabPodcasts },
];
</script>

<template>
  <div class="flex flex-col h-full w-full flex-1 overflow-hidden">
    <MusicTabs :tabs />
  </div>
</template>
