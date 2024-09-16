<script lang="ts" setup>
import { HeartIcon } from "lucide-vue-next";

const isLiked = useState("isLiked", () => false);
const toggleLike = () => (isLiked.value = !isLiked.value);
</script>

<template>
  <div class="flex items-center space-x-4">
    <NuxtImg
      src="https://picsum.photos/48"
      alt="Song cover"
      class="w-12 h-12 rounded-md"
    />
    <div class="flex flex-col justify-center ml-2">
      <div class="flex items-center gap-2">
        <h3 class="font-semibold">Song title</h3>
        <button @click="toggleLike">
          <HeartIcon
            class="w-4 h-4 text-pink-500"
            :fill="isLiked ? 'currentColor' : 'none'"
          />
        </button>
      </div>
      <p class="text-xs text-zinc-400">Artist name</p>
    </div>
  </div>
</template>
