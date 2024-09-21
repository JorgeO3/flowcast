<script lang="ts" setup>
import { HomeIcon, BellIcon } from "lucide-vue-next";

import { Button } from "@/components/ui/button";
import HeadUserMenu from "./MusicHeader/HeadUserMenu.vue";
</script>

<template>
  <header class="flex-shrink-0 h-14 flex items-center gap-4 px-4 md:px-6">
    <div
      class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4 justify-end"
    >
      <!-- Home button -->
      <Button as-child variant="secondary" size="icon" class="rounded-full">
        <NuxtLink to="/play" class="text-foreground">
          <HomeIcon class="w-5 h-5" />
        </NuxtLink>
      </Button>

      <!-- Notification button -->
      <Button variant="secondary" size="icon" class="rounded-full">
        <BellIcon class="w-5 h-5" />
      </Button>

      <!-- User menu -->
      <HeadUserMenu />
    </div>
  </header>
</template>
