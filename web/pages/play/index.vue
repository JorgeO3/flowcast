<script lang="ts" setup>
definePageMeta({ layout: "play" });
</script>

<template>
  <div class="h-full w-full">
    <!-- Header -->
    <MusicHeader />

    <!-- Main Content -->
    <div class="h-full w-full px-4 py-6">
      <MusicTabs />
    </div>
  </div>
</template>
