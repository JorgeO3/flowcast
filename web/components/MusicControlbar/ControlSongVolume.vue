<script lang="ts" setup>
import {
  CastIcon,
  Volume2Icon,
  VolumeXIcon,
  ListEndIcon,
} from "lucide-vue-next";

const isMuted = useState("isMuted", () => false);

const toggleMute = () => (isMuted.value = !isMuted.value);
</script>

<template>
  <div class="flex items-center space-x-4">
    <button
      class="p-3 rounded-full transition-all text-foreground"
      @click="toggleMute"
    >
      <VolumeXIcon v-if="isMuted" class="w-5 h-5" />
      <Volume2Icon v-else class="w-5 h-5" />
    </button>

    <button
      class="p-3 rounded-full hover:text-foreground transition-all active:text-foreground text-foreground/30"
    >
      <CastIcon class="w-5 h-5" />
    </button>

    <button
      class="p-3 rounded-full hover:text-foreground transition-all active:text-foreground text-foreground/30"
    >
      <ListEndIcon class="w-5 h-5" />
    </button>
  </div>
</template>
