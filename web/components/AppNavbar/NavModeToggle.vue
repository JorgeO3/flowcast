<script lang="ts" setup>
import { Button } from '@/components/ui/button';
import { MoonIcon, SunIcon } from 'lucide-vue-next';

const colorMode = useColorMode();

const handleColorMode = () => {
  const current = colorMode.preference;
  colorMode.preference = current === 'light' ? 'dark' : 'light';
}
</script>

<template>
  <Button @click="handleColorMode" variant="ghost" size="icon" class="transition-all ease-in-out duration-200">
    <template v-if="colorMode.preference == 'dark'">
      <MoonIcon class="w-5 h-5" />
    </template>

    <template v-else>
      <SunIcon class="w-5 h-5" />
    </template>
  </Button>
</template>