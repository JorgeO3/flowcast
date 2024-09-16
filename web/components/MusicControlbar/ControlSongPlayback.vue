<script lang="ts" setup>
import {
  PlayIcon,
  PauseIcon,
  SkipBackIcon,
  SkipForward,
  ShuffleIcon,
  RepeatIcon,
} from "lucide-vue-next";
import { Button } from "@/components/ui/button";

const isPlaying = useState("isPlaying", () => false);
const isShuffle = useState("isShuffle", () => false);
const isRepeat = useState("isRepeat", () => false);

const togglePlay = () => (isPlaying.value = !isPlaying.value);
const toggleShuffle = () => (isShuffle.value = !isShuffle.value);
const toggleRepeat = () => (isRepeat.value = !isRepeat.value);
const nextSong = () => console.log("Next song");
const previousSong = () => console.log("Previous song");

const baseButtonClasses = `
  p-3 rounded-full text-secondary hover:text-foreground 
  active:text-foreground transition-colors
`;

const activeShuffleClasses = computed(
  () => `${baseButtonClasses} ${isShuffle.value ? "text-white" : ""}`
);

const activeRepeatClasses = computed(
  () => `${baseButtonClasses} ${isRepeat.value ? "text-white" : ""}`
);
</script>

<template>
  <div class="flex items-center space-x-4">
    <button :class="activeShuffleClasses" @click="toggleShuffle">
      <ShuffleIcon class="w-5 h-5" />
    </button>
    <Button class="p-3 rounded-full" variant="ghost" @click="previousSong">
      <SkipBackIcon class="w-5 h-5" />
    </Button>
    <Button class="p-3 rounded-full" variant="ghost">
      <PauseIcon v-if="isPlaying" class="w-5 h-5" @click="togglePlay" />
      <PlayIcon v-else class="w-5 h-5" @click="togglePlay" />
    </Button>
    <Button class="p-3 rounded-full" variant="ghost" @click="nextSong">
      <SkipForward class="w-5 h-5" />
    </Button>
    <button :class="activeRepeatClasses" @click="toggleRepeat">
      <RepeatIcon class="w-5 h-5" />
    </button>
  </div>
</template>
