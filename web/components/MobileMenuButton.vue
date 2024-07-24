<script lang="ts" setup>
import { XIcon, MenuIcon } from 'lucide-vue-next';

defineProps<{ isOpen: boolean }>();
defineEmits<{ toggle: [] }>();
</script>

<template>
  <div class="md:hidden">
    <div class="transition-all ease-in-out duration-200">
      <Button variant="outline" size="icon" @click="$emit('toggle')" type="button">
        <XIcon v-if="isOpen" :style="{ opacity: isOpen ? 100 : 0 }" />
        <MenuIcon v-else />
      </Button>
    </div>
  </div>
</template>