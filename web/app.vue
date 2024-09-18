<script lang="ts" setup>
import { ConfigProvider } from "radix-vue";
import { useId } from "vue";

const useIdFunction = () => useId() || "";
</script>

<template>
  <!-- @vue-ignore -->
  <ConfigProvider :use-id="useIdFunction">
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </ConfigProvider>
</template>
