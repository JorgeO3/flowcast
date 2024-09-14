<script setup lang="ts">
import { MoveUpRightIcon } from "lucide-vue-next";
</script>

<template>
  <main class="flex flex-col w-full min-h-screen antialiased">
    <section class="flex w-full justify-center items-stretch min-h-screen">
      <div class="flex justify-center items-center w-1/2 p-8">
        <div class="w-full max-w-xl space-y-8">
          <h1 class="font-bold text-4xl md:text-5xl lg:text-6xl leading-tight">
            Discover a World of Music with Flowcast
          </h1>
          <p class="text-sm md:text-base">
            Welcome to Flowcast, where music comes alive. Experience a universe
            of endless tunes, handpicked playlists, and personalized
            recommendations just for you.
          </p>
          <NuxtLink to="/"
            class="inline-flex items-center font-semibold text-sm border border-primary py-2 pl-4 pr-3 rounded-full transition hover:bg-primary hover:text-primary-foreground cursor-pointer">
            DOWNLOAD THE APP
            <div class="bg-primary rounded-full ml-1 w-8 h-8 flex items-center justify-center">
              <MoveUpRightIcon class="text-black" />
            </div>
          </NuxtLink>
          <div class="flex items-center space-x-4">
            <NuxtImg src="/users.png" alt="Users" quality="70" />
            <div>
              <p class="font-bold">2 million+</p>
              <p class="text-sm">Flowcast users worldwide</p>
            </div>
          </div>
        </div>
      </div>

      <div class="visual-container w-1/2">
        <div class="logo-wrapper">
          <NuxtImg class="logo" src="/logo.png" alt="Logo" width="450" quality="10" />
        </div>
        <div class="phone-wrapper">
          <NuxtImg class="phone" src="/phone.png" alt="Phone" width="550" format="webp" preload quality="40" />
        </div>
      </div>
    </section>

    <section class="flex flex-col items-center py-28 gap-16">
      <h1 class="text-5xl font-extrabold">Scan To Download Flowcast App</h1>
      <NuxtImg src="/qr-phone.png" alt="phone with qr" width="400" format="avif" />
    </section>

    <section class="flex flex-col items-center container py-28 text-center gap-10">
      <h1 class="text-5xl font-extrabold">
        Transform your music experience with
        <br />
        Flowcast
      </h1>

      <p>
        Access a vast catalog of music without paying a penny. Flowcast offers
        you continuous playback, personalized
        <br />
        playlists and much more, all for free.
      </p>

      <div class="flex container flex-wrap mt-14 justify-evenly gap-10">
        <div class="flex flex-col items-center gap-14">
          <div class="h-16 w-16 flex justify-center items-center">
            <NuxtImg src="/adds-free.png" alt="adds free" format="webp" />
          </div>
          <div>
            <h1 class="font-semibold text-lg pb-1">
              Add-free music listenning
            </h1>
            <small class="font-medium">Enjoy uninterrupted music</small>
          </div>
        </div>

        <div class="flex flex-col items-center gap-14">
          <div class="h-16 w-16 flex justify-center items-center">
            <NuxtImg src="/offline.png" alt="adds free" format="webp" quality="60" />
          </div>
          <div>
            <h1 class="font-semibold text-lg pb-1">
              Add-free music listenning
            </h1>
            <small class="font-medium">Enjoy uninterrupted music</small>
          </div>
        </div>

        <div class="flex flex-col items-center gap-14">
          <div class="h-16 w-16 flex justify-center items-center">
            <NuxtImg src="/play.png" alt="adds free" format="webp" />
          </div>
          <div>
            <h1 class="font-semibold text-lg pb-1">
              Add-free music listenning
            </h1>
            <small class="font-medium">Enjoy uninterrupted music</small>
          </div>
        </div>

        <div class="flex flex-col items-center gap-14">
          <div class="h-16 w-16 flex justify-center items-center">
            <NuxtImg src="/pay.png" alt="adds free" format="webp" sizes="" />
          </div>
          <div>
            <h1 class="font-semibold text-lg pb-1">
              Add-free music listenning
            </h1>
            <small class="font-medium">Enjoy uninterrupted music</small>
          </div>
        </div>
      </div>
    </section>
  </main>
</template>

<style scoped>
.visual-container {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  border-radius: 42px;
  overflow: hidden;
}

.visual-container::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(to bottom right, #102c1b, #111813);
  opacity: 0.7;
  z-index: 0;
}

.logo-wrapper {
  position: relative;
  z-index: -1;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.phone-wrapper {
  position: absolute;
  bottom: 0;
  right: 0;
  max-width: 100%;
  z-index: 2;
}

.logo {
  max-width: 100%;
  height: auto;
  opacity: 0.9;
}

.phone {
  max-width: 100%;
  height: auto;
}
</style>
