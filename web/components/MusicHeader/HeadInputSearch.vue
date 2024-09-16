<script lang="ts" setup>
import { Search } from "lucide-vue-next";
</script>

<template>
  <form class="ml-auto flex-1 sm:flex-initial">
    <div class="relative">
      <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
      <Input
        type="search"
        placeholder="Search products..."
        class="pl-8 sm:w-[300px] md:w-[200px] lg:w-[300px]"
      />
    </div>
  </form>
</template>

<style></style>
