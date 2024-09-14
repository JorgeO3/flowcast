export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  css: ["~/assets/css/main.css"],
  modules: [
    "@nuxtjs/tailwindcss",
    "shadcn-nuxt",
    "@nuxtjs/color-mode",
    "@nuxt/fonts",
    "@nuxt/image",
    "nuxt-delay-hydration",
  ],
  shadcn: {
    prefix: '',
    componentDir: './components/ui'
  },
  colorMode: {
    classSuffix: ''
  },
  fonts: {
    families: [{ name: 'Inter', provider: 'bunny', weights: [400, 500, 700], subsets: ['latin'] }],
  },
  devtools: {
    timeline: {
      enabled: true
    }
  },
  routeRules: {
    '/': { static: true, prerender: true },
    '/auth/login': { static: true, prerender: true },
    '/auth/register': { static: true, prerender: true },
  },
  delayHydration: {
    debug: process.env.NODE_ENV === 'development',
    mode: 'mount',
  },
  nitro: {
    compressPublicAssets: true,
    prerender: {
      crawlLinks: true,
    }
    
  },
  experimental: {
    componentIslands: true,
  },
})