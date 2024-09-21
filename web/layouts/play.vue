<template>
  <div class="flex flex-col h-screen">
    <!-- Área Principal: Sidebar y Contenido Principal -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar Izquierdo -->
      <div class="w-64 flex-shrink-0 hidden md:block">
        <MusicSidebar />
      </div>

      <!-- Contenido Principal -->
      <div class="flex-1 flex flex-col overflow-hidden">
        <!-- Slot para insertar contenido específico de la página -->
        <MusicHeader />
        <slot />
      </div>
    </div>

    <!-- Music Control Bar -->
    <div class="h-20 flex-shrink-0">
      <MusicControlbar />
    </div>
  </div>
</template>
