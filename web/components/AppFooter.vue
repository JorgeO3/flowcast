<script setup lang="ts">
import {
  FacebookIcon,
  GlobeIcon,
  InstagramIcon,
  TwitterIcon,
} from "lucide-vue-next";
</script>

<template>
  <footer class="bg-black text-white py-12 px-6 md:px-12 lg:px-24">
    <div class="max-w-7xl mx-auto">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-8">
        <div class="col-span-1 md:col-span-2 lg:col-span-1">
          <p class="font-bold text-2xl">FlowCast</p>
        </div>
        <div>
          <h2 class="font-bold mb-4">Company</h2>
          <ul class="space-y-2">
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white"
                >About</NuxtLink
              >
            </li>
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white">
                Jobs
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white">
                For the Record
              </NuxtLink>
            </li>
          </ul>
        </div>
        <div>
          <h2 class="font-bold mb-4">Communities</h2>
          <ul class="space-y-2">
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white">
                For Artists
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white"
                >Developers</NuxtLink
              >
            </li>
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white"
                >Advertising</NuxtLink
              >
            </li>
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white"
                >Investors</NuxtLink
              >
            </li>
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white"
                >Vendors</NuxtLink
              >
            </li>
          </ul>
        </div>
        <div>
          <h2 class="font-bold mb-4">Useful NuxtLinknks</h2>
          <ul class="space-y-2">
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white"
                >Support</NuxtLink
              >
            </li>
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white"
                >Web Player</NuxtLink
              >
            </li>
            <li>
              <NuxtLink to="#" class="text-gray-300 hover:text-white"
                >Free Mobile App</NuxtLink
              >
            </li>
          </ul>
        </div>
        <div>
          <div class="flex space-x-4">
            <NuxtLink
              to="#"
              class="h-11 w-11 rounded-full flex justify-center items-center hover:bg-white hover:text-black transition-all ease-in-out"
            >
              <InstagramIcon class="w-5 h-5 m-auto" />
            </NuxtLink>
            <NuxtLink
              to="#"
              class="h-11 w-11 rounded-full flex justify-center items-center hover:bg-white hover:text-black transition-all ease-in-out"
            >
              <TwitterIcon class="w-5 h-5 m-auto" />
            </NuxtLink>
            <NuxtLink
              to="#"
              class="h-11 w-11 rounded-full flex justify-center items-center hover:bg-white hover:text-black transition-all ease-in-out"
            >
              <FacebookIcon class="w-5 h-5" />
            </NuxtLink>
          </div>
        </div>
      </div>
      <div
        class="mt-12 pt-8 border-t border-gray-800 flex flex-col md:flex-row justify-between items-center"
      >
        <div
          class="flex flex-wrap justify-center md:justify-start gap-4 mb-4 md:mb-0"
        >
          <NuxtLink to="#" class="text-gray-300 hover:text-white text-sm"
            >Legal</NuxtLink
          >
          <NuxtLink to="#" class="text-gray-300 hover:text-white text-sm"
            >Privacy Center</NuxtLink
          >
          <NuxtLink to="#" class="text-gray-300 hover:text-white text-sm"
            >Privacy PoNuxtLinkcy</NuxtLink
          >
          <NuxtLink to="#" class="text-gray-300 hover:text-white text-sm"
            >Cookies</NuxtLink
          >
          <NuxtLink to="#" class="text-gray-300 hover:text-white text-sm"
            >About Ads</NuxtLink
          >
        </div>
        <div class="flex items-center">
          <GlobeIcon class="w-4 h-4 mr-2" />
          <span class="text-gray-300 text-sm">USA</span>
        </div>
      </div>
      <div class="mt-4 text-center md:text-left text-gray-300 text-sm">
        © 2023 Flowcast AB
      </div>
    </div>
  </footer>
</template>
