<script lang="ts" setup>
import NavLogo from "./AppNavbar/NavLogo.vue";
import NavMenu from "./AppNavbar/NavMenu.vue";
import NavAuthLink from "./AppNavbar/NavAuthLink.vue";

const navLinks = ref([
  { name: "Home", path: "/" },
  { name: "Discover", path: "/discover" },
  { name: "About", path: "/about" },
  { name: "Blog", path: "/blog" },
]);

const { y } = useScroll();

const HEADER_CLASSES = {
  base: "transition-all duration-300 w-full flex fixed z-10 top-0 bg-transparent",
  expanded: "py-8",
  shrunk: "py-2 border-b backdrop-blur bg-background/95 supports-[backdrop-filter]:bg-background/60"
};

const headerClasses = computed(() =>
  `${HEADER_CLASSES.base} ${y.value ? HEADER_CLASSES.shrunk : HEADER_CLASSES.expanded}`
);
</script>

<template>
  <header :class="headerClasses">
    <div class="container mx-auto flex justify-between items-center">
      <NavLogo />
      <div class="flex gap-10 items-center">
        <NavMenu :links="navLinks" />
        <NavAuthLink to="/auth/login" />
      </div>
    </div>
  </header>
</template>