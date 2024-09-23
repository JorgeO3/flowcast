<script lang="ts" setup>
import type { Tab } from "@/utils/tabTypes.js";

import TabAll from "@/components/HomeTabs/TabAll.vue";
import TabMusic from "@/components/HomeTabs/TabMusic.vue";
import TabPodcasts from "@/components/HomeTabs/TabPodcasts.vue";

definePageMeta({ layout: "play" });

const tabs: Tab[] = [
  { name: "All", value: "all", component: TabAll },
  { name: "Music", value: "music", component: TabMusic },
  { name: "Podcasts", value: "podcasts", component: TabPodcasts },
];
</script>

<template>
  <div class="flex flex-col h-full w-full flex-1 overflow-hidden">
    <MusicTabs :tabs />
  </div>
</template>
