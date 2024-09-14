import { defineNuxtModule, useNitro, addPlugin, createResolver } from '@nuxt/kit'

export default defineNuxtModule({
    setup(options, nuxt) {
      const resolver = createResolver(import.meta.url)
  
      nuxt.hook('ready', () => {
        const nitro = useNitro()
        if (nitro.options.static && nuxt.options.experimental.payloadExtraction === undefined) {
          console.warn('Using experimental payload extraction for full-static output. You can opt-out by setting `experimental.payloadExtraction` to `false`.')
          nuxt.options.experimental.payloadExtraction = true
        }
        nitro.options.replace['process.env.NUXT_PAYLOAD_EXTRACTION'] = String(!!nuxt.options.experimental.payloadExtraction)
        nitro.options._config.replace!['process.env.NUXT_PAYLOAD_EXTRACTION'] = String(!!nuxt.options.experimental.payloadExtraction)
  
        if (!nuxt.options.dev && nuxt.options.experimental.payloadExtraction) {
          addPlugin(resolver.resolve(nuxt.options.appDir, 'plugins/payload.client'))
        }
      })
    }
  })
  