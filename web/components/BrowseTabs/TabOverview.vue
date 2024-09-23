<script lang="ts" setup>
import {
  HeartIcon,
  CloudDownloadIcon,
  EllipsisVerticalIcon,
} from "lucide-vue-next";
import Button from "@/components/ui/button/Button.vue";
import MusicSection from "@/components/MusicSection.vue";

const likes = useState("likes", () => trackList.map((track) => track.liked));
const toggleLike = (i: number) => (likes.value[i] = !likes.value[i]);
</script>

<template>
  <div class="flex flex-col px-4 gap-10 overflow-y-auto">
    <!-- Sección de Títulos -->
    <div class="flex flex-col w-full">
      <div class="flex items-end justify-between">
        <div class="space-y-1">
          <h2 class="text-2xl font-semibold tracking-tight">
            {{ "Your Top Hits" }}
          </h2>
          <p class="text-sm text-muted-foreground">
            {{ "The songs you can't stop playing" }}
          </p>
        </div>

        <Button
          :variant="null"
          class="text-sm font-medium hover:text-foreground text-muted-foreground p-0 h-fit"
        >
          All
        </Button>
      </div>

      <!-- Lista de Canciones -->
      <div class="flex flex-col overflow-y-auto h-full w-full gap-y-2 py-5">
        <template v-for="(track, i) in trackList" :key="track.title">
          <div
            class="flex items-center justify-between py-2 border rounded-md text-muted-foreground text-sm px-4"
          >
            <div class="flex items-center w-2/5">
              <p class="pr-4">{{ i + 1 }}</p>
              <NuxtImg
                :src="track.cover"
                alt="Album cover"
                class="w-10 h-10 rounded-md"
              />
              <p class="pl-4">{{ track.title }}</p>
            </div>

            <div class="flex items-center w-1/4">
              <p>{{ track.artist }}</p>
            </div>

            <div class="flex items-center w-1/4">
              <p>{{ track.album }}</p>
            </div>

            <div class="flex gap-x-2 items-center w-1/5 justify-end">
              <button>
                <HeartIcon
                  @click="toggleLike(i)"
                  class="w-5 h-5 text-pink-500 hover:fill-current"
                  :fill="likes[i] ? 'currentColor' : 'none'"
                />
              </button>

              <button>
                <CloudDownloadIcon class="w-5 h-5" />
              </button>

              <p class="flex w-10 items-center justify-center">
                {{ track.duration }}
              </p>

              <button>
                <EllipsisVerticalIcon class="w-5 h-5" />
              </button>
            </div>
          </div>
        </template>
      </div>
    </div>

    <!-- Sección de Artistas -->
    <MusicSection :="trendingArtistSection" />
  </div>
</template>

<script lang="ts">
import type { Section } from "@/utils/tabAllSections.js";

const trackList = [
  {
    title: "The Less I Know The Better",
    artist: "Tame Impala",
    album: "Currents",
    duration: "3:39",
    cover: "https://picsum.photos/id/44/40",
    liked: true,
  },
  {
    title: "Let It Happen",
    artist: "Tame Impala",
    album: "Currents",
    duration: "7:46",
    cover: "https://picsum.photos/id/43/40",
    liked: false,
  },
  {
    title: "Elephant",
    artist: "Tame Impala",
    album: "Lonerism",
    duration: "3:32",
    cover: "https://picsum.photos/id/45/40",
    liked: false,
  },
  {
    title: "Feels Like We Only Go Backwards",
    artist: "Tame Impala",
    album: "Lonerism",
    duration: "3:12",
    cover: "https://picsum.photos/id/40/40",
    liked: true,
  },
  {
    title: "Borderline",
    artist: "Tame Impala",
    album: "The Slow Rush",
    duration: "4:33",
    cover: "https://picsum.photos/id/41/40",
    liked: false,
  },
  {
    title: "Lost In Yesterday",
    artist: "Tame Impala",
    album: "The Slow Rush",
    duration: "4:10",
    cover: "https://picsum.photos/id/42/40",
    liked: false,
  },
];

const trendingArtistSection: Section = {
  id: "d12b75b4-2e66-4532-aadd-4b52898798ea",
  title: "Trending Artists",
  text: "Explore this carefully curated collection of musical gems.",
  link: "/play/section/d12b75b4-2e66-4532-aadd-4b52898798ea",
  type: "Album",
  boxes: [
    {
      id: "43d948ce-8a0b-4250-8f18-535c6047f543",
      type: "Act",
      cover: {
        src: "https://picsum.photos/id/77/200",
        type: "Square",
      },
      playlists: [
        {
          name: "Rock",
          id: "8fd1eec8-48c7-4833-9540-2c9259f1eebc",
        },
        {
          name: "Stage And Screen",
          id: "f2297d55-4b1e-4cde-9e54-92875506764e",
        },
        {
          name: "Non Music",
          id: "9d2238c8-d127-4d22-9603-b71c2fd0aa00",
        },
      ],
      title: "Brown Eyed Girl",
      text: "An emotional journey through heartfelt lyrics and powerful beats.",
      link: "/play/Act/43d948ce-8a0b-4250-8f18-535c6047f543",
    },
    {
      id: "bc62253c-84dd-4c4b-9300-e32b8da5cbbe",
      type: "Podcast",
      cover: {
        src: "https://picsum.photos/id/189/200",
        type: "Square",
      },
      playlists: [
        {
          name: "World",
          id: "6dfb0138-d64e-4d2a-b38d-91f8fd7050d6",
        },
        {
          name: "Country",
          id: "b4a43728-6797-4d6d-8f1c-f198461cb73b",
        },
        {
          name: "Non Music",
          id: "31600bb4-67b5-4f36-975a-fedca8648de6",
        },
      ],
      title: "Eye of the Tiger",
      text: "An eclectic mix of melodies for every mood.",
      link: "/play/show/bc62253c-84dd-4c4b-9300-e32b8da5cbbe",
    },
    {
      id: "1833b9da-7aea-4368-aa6e-45ade80ff8ff",
      type: "Radio",
      cover: {
        src: "https://picsum.photos/id/284/200",
        type: "Square",
      },
      playlists: [
        {
          name: "Jazz",
          id: "59156f6f-725e-4e95-8749-a9a1278e7454",
        },
        {
          name: "Electronic",
          id: "f4eea635-a280-438c-a849-6afe5e00512d",
        },
        {
          name: "Folk",
          id: "88773810-07ac-4d58-9d03-c2b935f473cc",
        },
      ],
      title: "Help Me",
      text: "Feel the rhythm of timeless classics.",
      link: "/play/Playlist/1833b9da-7aea-4368-aa6e-45ade80ff8ff",
    },
    {
      id: "d035ec91-8598-45cb-952b-41f234c3ff75",
      type: "Podcast",
      cover: {
        src: "https://picsum.photos/id/640/200",
        type: "Square",
      },
      playlists: [
        {
          name: "Soul",
          id: "ff829e49-1282-4069-97be-fd58a62ef291",
        },
      ],
      title: "You Light Up My Life",
      text: "Feel the rhythm of timeless classics.",
      link: "/play/show/d035ec91-8598-45cb-952b-41f234c3ff75",
    },
    {
      id: "ec52c6d3-74b3-44f2-addf-0e1c92d90999",
      type: "Album",
      cover: {
        src: "https://picsum.photos/id/982/200",
        type: "Circle",
      },
      playlists: [
        {
          name: "Latin",
          id: "e2e5f13b-b733-4086-b6fb-5ae4190f856b",
        },
        {
          name: "Latin",
          id: "a9e2141f-3f08-4d8d-8c37-be4a1c00da4f",
        },
        {
          name: "Latin",
          id: "beb77fe7-a3f1-4a5f-a268-01341b8947e3",
        },
        {
          name: "Country",
          id: "1e6e844a-c875-40ec-9e9b-77de2121759b",
        },
        {
          name: "World",
          id: "c9feea0e-f8a4-49f4-af9b-92a045e10fc2",
        },
      ],
      title: "Take a Bow",
      text: "An emotional journey through heartfelt lyrics and powerful beats.",
      link: "/play/Album/ec52c6d3-74b3-44f2-addf-0e1c92d90999",
    },
    {
      id: "9d010fad-1b7d-4ecb-91c5-2ece623c1418",
      type: "Radio",
      cover: {
        src: "https://picsum.photos/id/989/200",
        type: "Circle",
      },
      playlists: [
        {
          name: "Electronic",
          id: "397fcd95-eb22-4383-9c56-6b92be6e8319",
        },
        {
          name: "Funk",
          id: "a7e80283-6edb-45ef-b3e5-3a28614feed5",
        },
      ],
      title: "Come On-a My House",
      text: "Feel the rhythm of timeless classics.",
      link: "/play/Playlist/9d010fad-1b7d-4ecb-91c5-2ece623c1418",
    },
    {
      id: "b434d7d8-9458-4c34-b6d9-e5266e935650",
      type: "Podcast",
      cover: {
        src: "https://picsum.photos/id/459/200",
        type: "Square",
      },
      playlists: [
        {
          name: "Country",
          id: "9aa95290-d586-4a50-baea-3b0330309a08",
        },
        {
          name: "Blues",
          id: "22d958e5-8828-454b-8b2d-577ef0cfafd6",
        },
        {
          name: "Latin",
          id: "dce9c342-d389-4487-affe-da9e1f2db618",
        },
        {
          name: "Soul",
          id: "6df7e663-ee45-4915-a837-cd689f96c295",
        },
        {
          name: "Country",
          id: "d59a6928-2597-4805-a3ac-ceee08ee2aad",
        },
      ],
      title: "Over the Rainbow",
      text: "Experience the latest hits that are shaping the world of music.",
      link: "/play/show/b434d7d8-9458-4c34-b6d9-e5266e935650",
    },
    {
      id: "b5184ee4-12f2-4232-ab2d-0e64bfc54b14",
      type: "Album",
      cover: {
        src: "https://picsum.photos/id/447/200",
        type: "Circle",
      },
      playlists: [
        {
          name: "World",
          id: "f25f15e0-ecb8-46a6-9e02-3fd9f1261df9",
        },
        {
          name: "Metal",
          id: "2c50677d-e232-41ba-9222-49d8b4129564",
        },
      ],
      title: "Time of the Season",
      text: "An eclectic mix of melodies for every mood.",
      link: "/play/Album/b5184ee4-12f2-4232-ab2d-0e64bfc54b14",
    },
    {
      id: "73aa167a-fd1b-4be5-90a4-0c02bf70da2e",
      type: "Playlist",
      cover: {
        src: "https://picsum.photos/id/867/200",
        type: "Circle",
      },
      playlists: [
        {
          name: "Folk",
          id: "d6c6e324-bf41-4550-afa6-dc730110b3fe",
        },
      ],
      title: "Tutti Frutti",
      text: "Experience the latest hits that are shaping the world of music.",
      link: "/play/Playlist/73aa167a-fd1b-4be5-90a4-0c02bf70da2e",
    },
    {
      id: "9be1fd87-8b62-4bd7-ac5d-3e952f93af8f",
      type: "Album",
      cover: {
        src: "https://picsum.photos/id/274/200",
        type: "Square",
      },
      playlists: [
        {
          name: "Latin",
          id: "f63daa36-eafd-4918-af63-8e9102aef268",
        },
        {
          name: "Pop",
          id: "8314d30f-a702-49b8-a0fc-8b3ecf0402b6",
        },
        {
          name: "Electronic",
          id: "a4c77705-3786-4ea3-b5d3-33e7fa91c8e9",
        },
      ],
      title: "Rag Doll",
      text: "An eclectic mix of melodies for every mood.",
      link: "/play/Album/9be1fd87-8b62-4bd7-ac5d-3e952f93af8f",
    },
  ],
};
</script>
