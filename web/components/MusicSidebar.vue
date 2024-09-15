<script setup lang="ts">
import { CirclePlayIcon, LayoutGridIcon, RadioIcon, ListMusic, Music2Icon, UserIcon, MicVocalIcon, DiscIcon, LibraryBigIcon } from 'lucide-vue-next';
import { ScrollArea } from '@/components/ui/scroll-area';
import SideSection from './MusicSidebar/SideSection.vue';
import SideLink from './MusicSidebar/SideLink.vue';

interface SidebarLink {
  icon: any;
  text: string;
  path: string;
}

interface SidebarSection {
  title: string;
  items: SidebarLink[];
}

const sectionsMusicSidebar: SidebarSection[] = [
  {
    title: 'Discover',
    items: [
      { icon: CirclePlayIcon, text: 'Listen Now', path: '/play' },
      { icon: LayoutGridIcon, text: 'Browse', path: '/play/browse' },
      { icon: RadioIcon, text: 'Radio', path: '/play/radio' },
    ],
  },
  {
    title: 'Library',
    items: [
      { icon: ListMusic, text: 'Playlists', path: '/play/playlists' },
      { icon: Music2Icon, text: 'Songs', path: '/play/songs' },
      { icon: UserIcon, text: 'Made for You', path: '/play/made-for-you' },
      { icon: MicVocalIcon, text: 'Artists', path: '/play/artists' },
      { icon: LibraryBigIcon, text: 'Albums', path: '/play/albums' },
    ],
  },
  {
    title: 'Playlists',
    items: [
      { icon: DiscIcon, text: 'Chill Mix', path: '/play/playlist/1' },
      { icon: DiscIcon, text: 'Workout Mix', path: '/play/playlist/2' },
      { icon: DiscIcon, text: 'Focus Mix', path: '/play/playlist/3' },
      { icon: DiscIcon, text: 'Party Mix', path: '/play/playlist/4' },
      { icon: DiscIcon, text: 'Study Mix', path: '/play/playlist/5' },
      { icon: DiscIcon, text: 'Sleep Mix', path: '/play/playlist/6' },
      { icon: DiscIcon, text: 'Relax Mix', path: '/play/playlist/7' },
      { icon: DiscIcon, text: 'Travel Mix', path: '/play/playlist/8' },
      { icon: DiscIcon, text: 'Road Trip Mix', path: '/play/playlist/9' },
      { icon: DiscIcon, text: 'Summer Mix', path: '/play/playlist/10' },
      { icon: DiscIcon, text: 'Winter Mix', path: '/play/playlist/11' },
    ],
  }
];

</script>

<template>
  <div class="flex flex-col h-full w-full justify-evenly py-4 border-r">
    <template v-for="(section, i) in sectionsMusicSidebar" :key="`${section}-${i}`">

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
                <template v-for="(playlist, j) in items" :key="`${playlist}-${j}`"
                  class="w-full justify-start font-normal">
                  <SideLink :playlist />
                </template>
              </div>
            </ScrollArea>
          </template>

        </SideSection>
      </template>

    </template>
  </div>
</template>