<template>
  <div class="flex flex-col h-screen">
    <div class="flex h-[92%]">

      <!-- Sidebar -->
      <div class="w-64 h-full">
        <MusicSidebar />
      </div>

      <!-- Main Content -->
      <div class="w-full">
        <slot />
      </div>
    </div>

    <!-- Music Control Bar -->
    <MusicControlbar />
  </div>
</template>