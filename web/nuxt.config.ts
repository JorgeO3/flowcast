// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  devtools: {
    enabled: true,

    timeline: {
      enabled: true,
    },
  },
  css: ["~/assets/css/main.css"],
  modules: ["@nuxt/image", "@unocss/nuxt", "@nuxt/icon", "@nuxt/fonts"],
  experimental: {
    componentIslands: true,
  },
  unocss: {
    preflight: true,
  },
  routeRules: {
    "/": { prerender: true },
  },
});