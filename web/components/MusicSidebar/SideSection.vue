<script lang="ts" setup>
import SideLink from './SideLink.vue';

interface SidebarLink {
  icon: Component;
  text: string;
  path: string;
}

interface SidebarSection {
  title: string;
  items: SidebarLink[];
}

defineProps<{ section?: SidebarSection }>();
</script>

<template>
  <div v-if="section" :class="[$attrs.class ? $attrs.class : 'px-3 py-2']">

    <!-- Header -->
    <slot name="header" :title="section.title">
      <h2 class="mb-2 px-4 text-lg font-semibold tracking-tight">
        {{ section.title }}
      </h2>
    </slot>

    <!-- Content -->
    <slot name="content" :items="section.items">
      <div class="space-y-1">
        <template v-for="(item, i) in section.items" :key="`${item}-${i}`">
          <SideLink :playlist="item" />
        </template>
      </div>
    </slot>

  </div>
</template>