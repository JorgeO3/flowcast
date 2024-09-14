// import { defineNuxtModule, useNitro, createResolver } from '@nuxt/kit'
// import { promises as fs } from 'fs'
// import { join } from 'path'

// export default defineNuxtModule({
//   setup(options, nuxt) {
//     const resolver = createResolver(import.meta.url)

//     nuxt.hook('nitro:build:before', async (nitro) => {
//       const publicDir = join(nuxt.options.srcDir, 'public')
//       const faviconPath = join(publicDir, 'favicon.png')

//       try {
//         // Check if favicon exists
//         await fs.access(faviconPath)
        
//         // Read the favicon file
//         const faviconContents = await fs.readFile(faviconPath)


//         // Add the favicon to the Nitro prerendered routes
//         nitro.options.prerender.routes.push(faviconPath)
//       } catch (err) {
//         console.error('Favicon not found or could not be read:', err)
//       }
//     })
//   }
// })