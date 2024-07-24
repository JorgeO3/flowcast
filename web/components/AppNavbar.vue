<script setup lang="ts">
import NavLogo from './AppNavbar/NavLogo.vue';
import NavMenu from './AppNavbar/NavMenu.vue';
import { ArrowRightIcon } from 'lucide-vue-next';

const { y } = useWindowScroll();
const navOpen = useState('navOpen', () => false);

const links = [
  { name: 'Home', path: '/' },
  { name: 'Blog', path: '/blog' },
  { name: 'About', path: '/about' },
  { name: 'Discover', path: '/discover' },
];

const toggleNav = () => { navOpen.value = !navOpen.value };
const shrinkableClasses = computed(() => y.value ? 'border-b h-14' : 'border-b-0 h-20');
</script>

<template>
  <header
    class="flex w-full h-20 items-center border-b absolute z-10 top-0 px-4 md:px-10 justify-between transition-all ease-in-out duration-300"
    :class="shrinkableClasses">
    <!-- Title and Logo -->
    <NavLogo />

    <div class="flex gap-x-10">
      <!-- Navigation -->
      <NavMenu :links :isOpen="navOpen" @toggle="toggleNav" />

      <Button variant="default" to="/auth/login"
        class="rounded-full hidden md:block font-semibold transition-colors ease-in-out">
        <div class="flex gap-1">
          <p> Login</p>
          <ArrowRightIcon />
        </div>
      </Button>
    </div>
  </header>
</template>

<style scoped></style>