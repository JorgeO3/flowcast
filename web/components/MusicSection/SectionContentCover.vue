<script lang="ts" setup>
import type { Cover } from "@/utils/sectionTypes";

interface Props extends Cover {
  imgClass?: string;
}

defineProps<Props>();
</script>

<template>
  <div :class="$attrs.class">
    <NuxtImg
      :src
      sizes="160px"
      fit="fill"
      :class="[
        cn('h-auto w-auto object-cover transition-all hover:scale-105'),
        imgClass || '',
      ]"
    />
  </div>
</template>
