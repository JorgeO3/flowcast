<script lang="ts" setup>
import { Search } from "lucide-vue-next";

const router = useRouter();
const search = useState("search", () => "");

const handleFocus = () => {
  router.push("/play/browse");
};
</script>

<template>
  <form class="flex-1 sm:flex-initial">
    <div class="relative">
      <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
      <Input
        type="search"
        v-model="search"
        @focusin="handleFocus"
        placeholder="What do you want to play?"
        class="pl-8 sm:w-[300px] md:w-[200px] lg:w-[300px]"
      />
    </div>
  </form>
</template>
