// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  css: ["~/assets/css/main.css"],
  devtools: { enabled: true },
  modules: [
    "@nuxt/image",
    "@nuxt/fonts",
    "@vueuse/nuxt",
    "@unocss/nuxt",
    "@nuxt/icon",
  ],
  unocss: {
    preflight: true,
  },
});
