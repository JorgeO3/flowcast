<script lang="ts" setup>
interface Link {
  name: string;
  path: string;
}

defineProps<{ links: Link[]; isOpen: boolean; }>();
defineEmits<{ toggle: [] }>();
</script>

<template>
  <nav class="flex justify-between items-center gap-x-12">
    <!-- Mobile Menu Button -->
    <MobileMenuButton :isOpen @toggle="$emit('toggle')" />

    <!-- Navigation links -->
    <ul class="gap-x-10 hidden md:flex">
      <li v-for="link in links" :key="link.name" class="mx-2">
        <NuxtLink :to="link.path" class="text-sm font-semibold hover:text-green transition-colors ease-in-out"
          activeClass="text-green">
          {{ link.name }}
        </NuxtLink>
      </li>
    </ul>
  </nav>
</template>