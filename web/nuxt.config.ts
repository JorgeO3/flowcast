export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  css: ["~/assets/css/main.css"],
  modules: [
    "@nuxtjs/tailwindcss",
    "shadcn-nuxt",
    "@nuxtjs/color-mode",
    "@nuxt/fonts",
    "@nuxt/image",
  ],
  fonts: {
    // TODO: Optimize fonts with fontsquirrel [https://www.fontsquirrel.com/tools/webfont-generator]
    families: [{ name: 'Inter', provider: 'google', weights: [400, 700, 800], styles: ['normal'] }],
  },
  shadcn: {
    prefix: '',
    componentDir: './components/ui'
  },
  colorMode: {
    classSuffix: ''
  },
  image: {
    format: ['avif'],
    quality: 80,
  },
})