<script lang="ts" setup>
import { sections } from "@/utils/tabAllSections.js";
</script>

<template>
  <div class="flex flex-col h-full">
    <div class="flex-1 overflow-y-auto px-4">
      <div class="flex flex-col w-full">
        <template v-for="section in sections" :key="section.link">
          <MusicSection v-bind="section" />
        </template>
      </div>
    </div>
  </div>
</template>
