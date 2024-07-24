<script lang="ts" setup>
import { MoveUpRight } from 'lucide-vue-next';

useHead({ title: 'FlowCast - Home' });
</script>

<template>
  <main>
    <section>
      <div class="content-container">
        <div class="text-wrapper">
          <h1>
            Discover a World
            of Music with
            Flowcast
          </h1>
          <h6>
            Welcome to Flowcast, where music comes alive. Experience a universe of
            <br />
            endless tunes, handpicked playlists, and personalized recommendations
            <br />
            just for you.
          </h6>
          <NuxtLink class="download-link">
            DOWNLOAD THE APP
            <div class="icon-wrapper">
              <MoveUpRight color="black" :size="28" />
            </div>
          </NuxtLink>
        </div>
      </div>
      <div class="visual-container">
        <div class="logo-wrapper">
          <NuxtImg class="logo" src="/logo.svg" alt="Logo" width="450" />
        </div>
        <div class="phone-wrapper">
          <NuxtImg class="phone" src="/phone.png" alt="Phone" width="550" />
        </div>
      </div>
    </section>
  </main>
</template>

<style scoped>
main {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
}

section {
  display: flex;
  height: 100%;
}

.content-container {
  display: flex;
  flex: 1;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 24px;
}

.text-wrapper {
  width: 83.333333%;
  height: 66.666667%;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
}

h1 {
  font-weight: bold;
  font-size: 4rem;
}

h6 {
  margin-bottom: 16px;
}

.download-link {
  font-weight: 600;
  font-size: 0.875rem;
  border: 1px solid white;
  padding: 0.5rem 1.25rem;
  border-radius: 9999px;
  display: flex;
  align-items: center;
  transition: all 0.3s ease-in;
  width: fit-content;
}

.download-link:hover {
  background-color: var(--primary-color);
  color: var(--primary-foreground-color);
  cursor: pointer;
}

.icon-wrapper {
  width: 2.25rem;
  height: 2.25rem;
  margin-left: 0.5rem;
  background-color: var(--primary);
  border-radius: 9999px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.visual-container {
  position: relative;
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 42px;
}

.visual-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(to bottom right, #0D2616, #121212);
  border-radius: 42px;
  opacity: 0.7;
  z-index: 0;
}

.logo-wrapper {
  position: relative;
  border-radius: 42px;
  width: fit-content;
  /* height: auto; */
}

.phone-wrapper {
  position: absolute;
  bottom: 0;
  right: 0;
  max-width: 100%;
  height: auto;
  z-index: 3;
  overflow: hidden;
}

.logo {
  position: relative;
  opacity: 0.9;
  border-radius: 42px;
  z-index: -1;
}

.phone {
  position: relative;
  transform: translateY(18%);
  z-index: 3;
}
</style>
