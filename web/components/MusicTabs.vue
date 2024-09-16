<script lang="ts" setup>
import { Tabs, TabsContent, TabsTrigger, TabsList } from "@/components/ui/tabs";
import TabAll from "./MusicTabs/TabAll.vue";
import TabMusic from "./MusicTabs/TabMusic.vue";
import TabPodcasts from "./MusicTabs/TabPodcasts.vue";
import type { Component } from "vue";

interface Tab {
  name: string;
  value: string;
  component: Component;
}

const tabs: Tab[] = [
  { name: "All", value: "all", component: TabAll },
  { name: "Music", value: "music", component: TabMusic },
  { name: "Podcasts", value: "podcasts", component: TabPodcasts },
];
</script>

<template>
  <Tabs default-value="all" class="h-full space-y-6">
    <TabsList>
      <TabsTrigger v-for="tab in tabs" :key="tab.value" :value="tab.value">
        {{ tab.name }}
      </TabsTrigger>
    </TabsList>

    <TabsContent v-for="tab in tabs" :key="tab.value" :value="tab.value">
      <component :is="tab.component" />
    </TabsContent>
  </Tabs>
</template>

<!-- <template>
  <Tabs default-value="all" class="h-full space-y-6">
    <TabsList>
      <TabsTrigger value="all">All</TabsTrigger>
      <TabsTrigger value="music">Music</TabsTrigger>
      <TabsTrigger value="podcasts">Podcasts</TabsTrigger>
    </TabsList>

    <TabsContent value="all">
      <h1>All</h1>
    </TabsContent>

    <TabsContent value="music">
      <h2>Music</h2>
    </TabsContent>

    <TabsContent value="podcasts">
      <h2>Podcasts</h2>
    </TabsContent>
  </Tabs>
</template> -->
