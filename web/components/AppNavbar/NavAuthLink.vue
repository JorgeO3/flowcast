<script lang="ts" setup>
import { UserIcon } from 'lucide-vue-next';
import { Button } from '@/components/ui/button';

defineProps<{ to: string }>();
</script>

<template>
  <Button variant="ghost" as-child>
    <NuxtLink :to class="flex font-semibold text-sm">
      <slot>
        <UserIcon class="w-4 h-4 mr-2" /> Login
      </slot>
    </NuxtLink>
  </Button>
</template>