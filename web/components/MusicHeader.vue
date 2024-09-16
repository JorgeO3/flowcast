<script lang="ts" setup>
import HeadInputSearch from "./MusicHeader/HeadInputSearch.vue";
import HeadUserMenu from "./MusicHeader/HeadUserMenu.vue";
</script>

<template>
  <header
    class="sticky top-0 flex h-14 items-center gap-4 bg-background px-4 md:px-6"
  >
    <div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
      <HeadInputSearch />
      <HeadUserMenu />
    </div>
  </header>
</template>

<!-- <script lang="ts" setup>
import HeadInputSearch from "./MusicHeader/HeadInputSearch.vue";
import HeadUserMenu from "./MusicHeader/HeadUserMenu.vue";
import { Skeleton } from "@/components/ui/skeleton";
</script>

<template>
  <header
    class="sticky top-0 flex h-14 items-center gap-4 bg-background px-4 md:px-6"
  >
    <div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
      <HeadInputSearch />
      <ClientOnly>
        <HeadUserMenu />

        <template #fallback>
          <div class="flex items-center gap-2">
            <Skeleton class="h-10 w-1`0 rounded-full" />
          </div>
        </template>
      </ClientOnly>
    </div>
  </header>
</template> -->
