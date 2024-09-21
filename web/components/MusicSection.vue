<script lang="ts" setup>
import { Separator } from "@/components/ui/separator";
import type { Section, MediaContent } from "~/utils/tabAllSections.js";
import SectionHeader from "./MusicSection/SectionHeader.vue";

interface Props extends Section {}

defineProps<Props>();

const sliceBoxes = (boxes: MediaContent[]) => {
  return boxes.slice(0, 5);
};
</script>

<template>
  <div class="flex flex-col w-full h-fit mb-16">
    <!-- Header -->
    <SectionHeader :title :text :link />

    <Separator class="my-4" />

    <div class="flex justify-around h-72">
      <MusicMediaBox
        v-for="box in sliceBoxes(boxes)"
        :key="box.title"
        :="box"
      />
    </div>
  </div>
</template>
