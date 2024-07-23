<script lang="ts" setup>
import { BugIcon } from 'lucide-vue-next';
</script>

<template>
  <NuxtLayout name="404">
    <div class="flex h-full w-full flex-col items-center justify-center gap-8 bg-background px-4">
      <div class="flex max-w-md flex-col items-center justify-center gap-4 text-center">
        <BugIcon class="h-24 w-24 text-green" />
        <h1 class="text-4xl font-bold tracking-tighter sm:text-5xl">
          Page Not Found
        </h1>
        <p class="text-muted-foreground font font-medium">
          The page you are looking for does not exist or has been moved.
        </p>
        <NuxtLink to="/"
          class="inline-flex h-10 items-center justify-center rounded-md bg-primary px-8 text-sm font-medium text-primary-foreground shadow transition-colors hover:bg-primary/90 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-primary-foreground disabled:pointer-events-none disabled:opacity-50 dark:focus-visible:ring-primary-foreground">
          Go to Home
        </NuxtLink>
      </div>
    </div>
  </NuxtLayout>
</template>