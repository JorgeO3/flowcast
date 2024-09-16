<script lang="ts" setup>
import ControlSongInfo from "@/components/MusicControlbar/ControlSongInfo.vue";
import ControlSongPlayback from "@/components/MusicControlbar/ControlSongPlayback.vue";
import ControlSongVolume from "@/components/MusicControlbar/ControlSongVolume.vue";
</script>

<template>
  <div
    class="flex w-full h-full border-t items-center justify-between px-2 py-1"
  >
    <ControlSongInfo />
    <ControlSongPlayback />
    <ControlSongVolume />
  </div>
</template>
