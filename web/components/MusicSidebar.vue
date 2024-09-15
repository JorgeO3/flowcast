<script setup lang="ts">
import { sectionsMusicSidebar } from '@/utils/musicSidebarLinks';
import { ScrollArea } from '@/components/ui/scroll-area';

import SideSection from '@/components/MusicSidebar/SideSection.vue';
import SideLink from '@/components/MusicSidebar/SideLink.vue';

const sections = sectionsMusicSidebar();
</script>

<template>
  <div class="flex flex-col h-full w-full justify-evenly py-4 border-r">
    <template v-for="(section, i) in sections" :key="`${section.title}-${i}`">

      <template v-if="section.title !== 'Playlists'">
        <SideSection :section="section" />
      </template>

      <template v-else>
        <SideSection :section="section" class="py-2 flex-1">

          <!-- Header section -->
          <template #header="{ title }">
            <h2 class="relative px-7 text-lg font-semibold tracking-tight">
              {{ title }}
            </h2>
          </template>

          <!-- Content section -->
          <template #content="{ items }">
            <ScrollArea class="h-52 px-1 pt-1">
              <div class="space-y-1 p-2">
                <div v-for="playlist in items" :key="playlist.path" class="w-full justify-start font-normal">
                  <SideLink :playlist />
                </div>
              </div>
            </ScrollArea>
          </template>

        </SideSection>
      </template>

    </template>
  </div>
</template>