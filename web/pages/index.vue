<script lang="ts" setup>
useHead({ title: "Flowcast - Home" });
</script>

<template>
  <main class="min-h-screen overflow-y-scroll">

    <section class="flex w-full justify-center items-stretch min-h-screen">

      <div class="flex justify-center items-center w-1/2 p-8">
        <div class="w-full max-w-xl space-y-8">
          <h1 class="font-bold text-4xl md:text-5xl lg:text-6xl leading-tight">
            Discover a World of Music with Flowcast
          </h1>
          <p class="text-sm md:text-base">
            Welcome to Flowcast, where music comes alive. Experience a universe of
            endless tunes, handpicked playlists, and personalized recommendations
            just for you.
          </p>
          <NuxtLink to="/"
            class="inline-flex items-center font-semibold text-sm border py-2 px-4 rounded-full transition hover:bg-primary hover:text-primary-foreground cursor-pointer">
            DOWNLOAD THE APP
            <Icon name="ph:arrow-circle-up-right-fill" class="ml-1" color="black" :size="34" />
          </NuxtLink>
          <div class="flex items-center space-x-4">
            <NuxtImg src="/users.png" alt="Users" />
            <div>
              <p class="font-bold">2 million+</p>
              <p class="text-sm">Flowcast users worldwide</p>
            </div>
          </div>
        </div>
      </div>

      <div class="visual-container w-1/2">
        <div class="logo-wrapper">
          <NuxtImg class="logo" src="/logo.svg" alt="Logo" width="450" />
        </div>
        <div class="phone-wrapper">
          <NuxtImg class="phone" src="/phone.png" alt="Phone" width="550" />
        </div>
      </div>
    </section>

    <section class="flex min-h-screen">
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
  content: '';
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
  transform: translateY(18%);
}
</style>