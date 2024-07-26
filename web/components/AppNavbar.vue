<script lang="ts" setup>
import { } from "vue"
import NavLogo from "./AppNavbar/NavLogo.vue";
import NavMenu from "./AppNavbar/NavMenu.vue";
import NavAuthLink from "./AppNavbar/NavAuthLink.vue";

const links = ref([
  { name: "Home", path: "/" },
  { name: "Discover", path: "/discover" },
  { name: "About", path: "/about" },
  { name: "Blog", path: "/blog" },
]);


const el = ref<HTMLElement | null>(null)
const { x, y, isScrolling, arrivedState, directions } = useScroll(el)

const headerClass = computed(() => `
  transition-all duration-300 py-8 w-full absolute z-3 flex w-full
  ${y.value ? "bg-white shadow-md" : ""}
`);

const increment = () => {
  y.value = 30;
};
</script>

<template>
  <header :class="headerClass">
    <div class="container mx-auto flex justify-between items-center">

      <div ref="el" />

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
