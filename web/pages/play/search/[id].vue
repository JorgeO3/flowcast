<script lang="ts" setup>
definePageMeta({ layout: "play" });

const { params } = useRoute();
</script>

<template>
  <div>
    <h1>{{ params.id }}</h1>
  </div>
</template>
