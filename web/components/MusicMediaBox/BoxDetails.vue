<script lang="ts" setup>
interface Props {
  title: string;
  text: string;
}

const { text, title } = defineProps<Props>();

const slicedTitle = title.slice(0, 19);
const formattedTitle = slicedTitle + (title.length > 20 ? "..." : "");

const slicedText = text.slice(0, 45);
const formattedText = slicedText + (text.length > 50 ? "..." : "");
</script>

<template>
  <div class="space-y-1 text-sm h-14">
    <h3 class="font-medium leading-none">
      {{ formattedTitle }}
    </h3>
    <p class="text-xs text-muted-foreground">{{ formattedText }}</p>
  </div>
</template>
