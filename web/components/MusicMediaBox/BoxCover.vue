<script lang="ts" setup>
import { cn } from "@/utils/cn.js";
import type { Cover } from "@/utils/tabAllSections.js";

interface Props extends Cover {}
const { type } = defineProps<Props>();

const coverClasses =
  "flex overflow-hidden w-fit " +
  (type === "Square" ? "rounded-md" : "rounded-full");
</script>

<template>
  <div :class="coverClasses">
    <NuxtImg
      :src
      sizes="160px"
      fit="fill"
      :class="cn('h-auto w-auto object-cover transition-all hover:scale-105')"
    />
  </div>
</template>
