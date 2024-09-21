<script lang="ts" setup>
import type { Tab } from "@/utils/tabTypes.js";
import { Tabs, TabsContent, TabsTrigger, TabsList } from "@/components/ui/tabs";

defineProps<{ tabs: Tab[] }>();
</script>

<template>
  <Tabs default-value="all" class="flex flex-col h-full">
    <div class="w-full px-4">
      <TabsList class="w-fit">
        <TabsTrigger v-for="tab in tabs" :key="tab.value" :value="tab.value">
          {{ tab.name }}
        </TabsTrigger>
      </TabsList>
    </div>

    <TabsContent
      class="flex-grow overflow-hidden py-4"
      v-for="tab in tabs"
      :key="tab.value"
      :value="tab.value"
    >
      <component :is="tab.component" class="h-full w-full" />
    </TabsContent>
  </Tabs>
</template>
