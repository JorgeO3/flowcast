<script lang="ts" setup>
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { AppleIcon, FacebookIcon } from "lucide-vue-next"

definePageMeta({ layout: "auth" });
</script>

<template>
  <div class="w-full lg:grid lg:min-h-[600px] lg:grid-cols-2 xl:min-h-[800px]">
    <div class="flex items-center justify-center py-12">
      <div class="mx-auto grid w-[350px] gap-6">
        <div class="grid gap-2 text-center">
          <h1 class="text-3xl font-bold">Login</h1>
          <p class="text-balance text-muted-foreground">
            Enter your email below to login to your account
          </p>
        </div>
        <div class="grid gap-4">
          <div class="grid gap-2">
            <Label for="email" class="text-white">Email</Label>
            <Input id="username" type="text" placeholder="Correo electrónico o nombre de usuario"
              class="mt-1 bg-[#2a2a2a] border-0 text-white placeholder-gray-400" />
          </div>
          <div class="grid gap-2">
            <div class="flex items-center">
              <Label for="password" class="text-white">Password</Label>
              <NuxtLink href="#" class="ml-auto inline-block text-sm underline">
                Forgot your password?
              </NuxtLink>
            </div>
            <Input id="password" type="password" required placeholder="Contraseña"
              class="mt-1 bg-[#2a2a2a] border-0 text-white placeholder-gray-400" />
          </div>
          <Button type="submit" class="w-full bg-[#1DB954] hover:bg-[#31d969] text-black font-bold py-3"> Login
          </Button>
          <div class="mt-6">
            <NuxtLink href="#" class="text-white hover:underline block text-center">¿Has olvidado tu contraseña?
            </NuxtLink>
          </div>
          <div class="mt-6 flex items-center justify-center">
            <div class="border-t border-gray-600 w-full"></div>
            <span class="px-4 text-gray-400 text-sm">O</span>
            <div class="border-t border-gray-600 w-full"></div>
          </div>
          <div class="mt-6 space-y-4">
            <Button variant="outline" class="w-full bg-black text-white border-white hover:bg-gray-900">
              <AppleIcon class="mr-2 h-5 w-5" />
              Login with Apple
            </Button>
            <Button variant="outline" class="w-full bg-[#1877F2] text-white border-[#1877F2] hover:bg-[#0a66c2]">
              <FacebookIcon class="mr-2 h-5 w-5" />
              Login with Facebook
            </Button>
          </div>
        </div>
        <div class="mt-4 text-center text-sm">
          Don't have an account?
          <NuxtLink href="#" class="underline"> Sign up </NuxtLink>
        </div>
      </div>
    </div>
    <div class="hidden bg-muted lg:block">
      <div class="hidden bg-muted/40 lg:flex lg:flex-1 flex-col items-center justify-center text-center min-h-screen">
        <div class="flex items-center justify-center">
          <NuxtImg src="/logo.png" placeholder="Logo flowcast" width="200px" />
        </div>
        <h1 class="text-4xl font-bold py-6 dark:text-secondary-foreground">
          Flowcast
        </h1>
        <h3 class="font-medium text-sm text-muted-foreground dark:text-secondary-foreground">
          Musica para todos
        </h3>
      </div>
    </div>
  </div>
</template>

<style></style>
