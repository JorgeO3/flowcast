<script lang="ts" setup>
import {
  PlayIcon,
  PauseIcon,
  SkipBackIcon,
  SkipForwardIcon,
  ShuffleIcon,
  RepeatIcon,
} from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Slider } from "@/components/ui/slider";

const isPlaying = useState("isPlaying", () => false);
const isShuffle = useState("isShuffle", () => false);
const isRepeat = useState("isRepeat", () => false);
const volume = useState("volume", () => [50]);

const togglePlay = () => (isPlaying.value = !isPlaying.value);
const toggleShuffle = () => (isShuffle.value = !isShuffle.value);
const toggleRepeat = () => (isRepeat.value = !isRepeat.value);
const nextSong = () => console.log("Next song");
const previousSong = () => console.log("Previous song");

const baseButtonClasses = `
  p-0 h-fit rounded-full text-secondary hover:text-foreground 
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
  <div class="flex flex-col items-center gap-y-3 flex-1 h-fit">
    <div class="flex items-center space-x-4">
      <button :class="activeShuffleClasses" @click="toggleShuffle">
        <ShuffleIcon class="w-4 h-4" />
      </button>
      <Button
        class="p-0 h-0 rounded-full"
        variant="ghost"
        @click="previousSong"
      >
        <SkipBackIcon class="w-5 h-5 hover:fill-white" />
      </Button>
      <Button class="p-0 h-0 rounded-full" variant="ghost">
        <PauseIcon
          v-if="isPlaying"
          class="w-5 h-5 hover:fill-white transition-all"
          @click="togglePlay"
        />
        <PlayIcon
          v-else
          class="w-5 h-5 hover:fill-white transition-all"
          @click="togglePlay"
        />
      </Button>
      <Button class="p-0 h-0 rounded-full" variant="ghost" @click="nextSong">
        <SkipForwardIcon class="w-5 h-5 hover:fill-white transition-all" />
      </Button>
      <button :class="activeRepeatClasses" @click="toggleRepeat">
        <RepeatIcon class="w-4 h-4" />
      </button>
    </div>

    <Slider v-model="volume" :max="100" :step="1" class="w-3/5 mt-1" />
  </div>
</template>
