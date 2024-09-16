<script setup lang="ts">
import {
  FacebookIcon,
  GlobeIcon,
  InstagramIcon,
  TwitterIcon,
} from "lucide-vue-next";

// Reusable list of sections and links
const sections = [
  {
    title: "Company",
    links: [
      { text: "About", to: "#" },
      { text: "Jobs", to: "#" },
      { text: "For the Record", to: "#" },
    ],
  },
  {
    title: "Communities",
    links: [
      { text: "For Artists", to: "#" },
      { text: "Developers", to: "#" },
      { text: "Advertising", to: "#" },
      { text: "Investors", to: "#" },
      { text: "Vendors", to: "#" },
    ],
  },
  {
    title: "Useful Links",
    links: [
      { text: "Support", to: "#" },
      { text: "Web Player", to: "#" },
      { text: "Free Mobile App", to: "#" },
    ],
  },
];

// Social media links
const socialLinks = [
  { icon: InstagramIcon, ariaLabel: "Instagram", to: "#" },
  { icon: TwitterIcon, ariaLabel: "Twitter", to: "#" },
  { icon: FacebookIcon, ariaLabel: "Facebook", to: "#" },
];

// Legal links
const legalLinks = [
  { text: "Legal", to: "#" },
  { text: "Privacy Center", to: "#" },
  { text: "Privacy Policy", to: "#" },
  { text: "Cookies", to: "#" },
  { text: "About Ads", to: "#" },
];
</script>

<template>
  <footer class="bg-black text-white py-12 px-6 md:px-12 lg:px-24">
    <div class="max-w-7xl mx-auto">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-8">
        <!-- Logo -->
        <div class="col-span-1 md:col-span-2 lg:col-span-1">
          <p class="font-bold text-2xl">FlowCast</p>
        </div>

        <!-- Reusable Sections -->
        <template v-for="(section, index) in sections" :key="index">
          <div>
            <h2 class="font-bold mb-4">{{ section.title }}</h2>
            <ul class="space-y-2">
              <li v-for="(link, i) in section.links" :key="i">
                <NuxtLink :to="link.to" class="text-gray-300 hover:text-white">
                  {{ link.text }}
                </NuxtLink>
              </li>
            </ul>
          </div>
        </template>

        <!-- Social Media Icons -->
        <div>
          <div class="flex space-x-4">
            <NuxtLink
              v-for="(social, index) in socialLinks"
              :key="index"
              :to="social.to"
              :aria-label="social.ariaLabel"
              class="h-11 w-11 rounded-full flex justify-center items-center hover:bg-white hover:text-black transition-all ease-in-out"
            >
              <component :is="social.icon" class="w-5 h-5" />
            </NuxtLink>
          </div>
        </div>
      </div>

      <!-- Bottom Section -->
      <div
        class="mt-12 pt-8 border-t border-gray-800 flex flex-col md:flex-row justify-between items-center"
      >
        <!-- Legal Links -->
        <div
          class="flex flex-wrap justify-center md:justify-start gap-4 mb-4 md:mb-0"
        >
          <NuxtLink
            v-for="(legal, index) in legalLinks"
            :key="index"
            :to="legal.to"
            class="text-gray-300 hover:text-white text-sm"
          >
            {{ legal.text }}
          </NuxtLink>
        </div>

        <!-- Country/Region -->
        <div class="flex items-center">
          <GlobeIcon class="w-4 h-4 mr-2" />
          <span class="text-gray-300 text-sm">USA</span>
        </div>
      </div>

      <!-- Copyright -->
      <div class="mt-4 text-center md:text-left text-gray-300 text-sm">
        © 2023 FlowCast AB
      </div>
    </div>
  </footer>
</template>
