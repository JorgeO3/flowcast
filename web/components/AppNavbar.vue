<script setup lang="ts">
import { MenuIcon, XIcon } from 'lucide-vue-next';

const navOpen = useState('navOpen', () => false);
const links = [
  { name: 'Premiun', path: '/premiun' },
  { name: 'Support', path: '/support' },
  { name: 'Download', path: '/download' },
];

const toggleNav = () => {
  navOpen.value = !navOpen.value;
};

</script>

<template>
  <header class="flex w-full h-20 items-center px-10 justify-between">
    <!-- Title and Logo -->
    <NuxtLink class="flex items-center w-fit" to="/">
      <NuxtImg class="mr-2" src="/logo.svg" alt="FlowCast Logo" width="30" />
      <p class="font-bold text-2xl">FlowCast</p>
    </NuxtLink>

    <!-- Navigation -->
    <nav class="flex justify-between">

      <!-- Menu button (Burger) -->
      <button class="md:hidden dark:bg-secondary" @click="toggleNav" type="button">
        <XIcon v-if="navOpen" />
        <MenuIcon v-else />
      </button>

      <!-- Navigation links -->
      <ul v-if="navOpen" class="flex">
        <li v-for="link in links" :key="link.name" class="mx-2">
          <!-- <NuxtLink :to="link.path" class="text-lg font-semibold">{{ link.name }}</NuxtLink> -->
        </li>
      </ul>
    </nav>
  </header>
</template>