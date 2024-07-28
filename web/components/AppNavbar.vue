<script lang="ts" setup>
import NavLogo from "./AppNavbar/NavLogo.vue";
import NavMenu from "./AppNavbar/NavMenu.vue";
import NavAuthLink from "./AppNavbar/NavAuthLink.vue";

const links = ref([
  { name: "Home", path: "/" },
  { name: "Discover", path: "/discover" },
  { name: "About", path: "/about" },
  { name: "Blog", path: "/blog" },
]);

const { y } = useScroll();

const headerClass = computed(() => `
  transition-all duration-300 w-full flex w-full absolute z-3 
  sticky top-0 ${y.value ? "bg-background py-4 border-b" : "py-8"}
`);
</script>

<template>
  <header :class="headerClass">
    <div class="container mx-auto flex justify-between items-center">
      <!-- Logo -->
      <NavLogo />

      <div class="flex gap-10">
        <!-- Navigation Menu -->
        <NavMenu :links="links" />

        <!-- Auth Buttons -->
        <NavAuthLink to="/auth/login" />
      </div>
    </div>
  </header>
</template>