export interface Cover {
  src: string;
  type: "Circle" | "Square";
}

export interface Playlist {
  name: string;
  id: string;
}

export interface MediaContent {
  type: 'Album' | 'Act' | 'Playlist' | 'Podcast' | 'Radio';
  cover: Cover;
  playlists: Playlist[];
  title: string;
  text: string;
  id: string;
  link: string;
}

export interface Section {
  title: string;
  text: string;
  link: string;
  type: 'Album' | 'Act' | 'Playlist' | 'Podcast' | 'Radio' | 'Mixed';
  id: string;
  boxes: MediaContent[];
}

export const sections: Section[] = [
  {
    "id": "d12b75b4-2e66-4532-aadd-4b52898798ea",
    "title": "Journey Through Beats",
    "text": "Explore this carefully curated collection of musical gems.",
    "link": "/play/section/d12b75b4-2e66-4532-aadd-4b52898798ea",
    "type": "Album",
    "boxes": [
      {
        "id": "43d948ce-8a0b-4250-8f18-535c6047f543",
        "type": "Act",
        "cover": {
          "src": "https://picsum.photos/id/77/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Rock",
            "id": "8fd1eec8-48c7-4833-9540-2c9259f1eebc"
          },
          {
            "name": "Stage And Screen",
            "id": "f2297d55-4b1e-4cde-9e54-92875506764e"
          },
          {
            "name": "Non Music",
            "id": "9d2238c8-d127-4d22-9603-b71c2fd0aa00"
          }
        ],
        "title": "Brown Eyed Girl",
        "text": "An emotional journey through heartfelt lyrics and powerful beats.",
        "link": "/play/Act/43d948ce-8a0b-4250-8f18-535c6047f543"
      },
      {
        "id": "bc62253c-84dd-4c4b-9300-e32b8da5cbbe",
        "type": "Podcast",
        "cover": {
          "src": "https://picsum.photos/id/189/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "World",
            "id": "6dfb0138-d64e-4d2a-b38d-91f8fd7050d6"
          },
          {
            "name": "Country",
            "id": "b4a43728-6797-4d6d-8f1c-f198461cb73b"
          },
          {
            "name": "Non Music",
            "id": "31600bb4-67b5-4f36-975a-fedca8648de6"
          }
        ],
        "title": "Eye of the Tiger",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/show/bc62253c-84dd-4c4b-9300-e32b8da5cbbe"
      },
      {
        "id": "1833b9da-7aea-4368-aa6e-45ade80ff8ff",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/284/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Jazz",
            "id": "59156f6f-725e-4e95-8749-a9a1278e7454"
          },
          {
            "name": "Electronic",
            "id": "f4eea635-a280-438c-a849-6afe5e00512d"
          },
          {
            "name": "Folk",
            "id": "88773810-07ac-4d58-9d03-c2b935f473cc"
          }
        ],
        "title": "Help Me",
        "text": "Feel the rhythm of timeless classics.",
        "link": "/play/Playlist/1833b9da-7aea-4368-aa6e-45ade80ff8ff"
      },
      {
        "id": "d035ec91-8598-45cb-952b-41f234c3ff75",
        "type": "Podcast",
        "cover": {
          "src": "https://picsum.photos/id/640/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Soul",
            "id": "ff829e49-1282-4069-97be-fd58a62ef291"
          }
        ],
        "title": "You Light Up My Life",
        "text": "Feel the rhythm of timeless classics.",
        "link": "/play/show/d035ec91-8598-45cb-952b-41f234c3ff75"
      },
      {
        "id": "ec52c6d3-74b3-44f2-addf-0e1c92d90999",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/982/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Latin",
            "id": "e2e5f13b-b733-4086-b6fb-5ae4190f856b"
          },
          {
            "name": "Latin",
            "id": "a9e2141f-3f08-4d8d-8c37-be4a1c00da4f"
          },
          {
            "name": "Latin",
            "id": "beb77fe7-a3f1-4a5f-a268-01341b8947e3"
          },
          {
            "name": "Country",
            "id": "1e6e844a-c875-40ec-9e9b-77de2121759b"
          },
          {
            "name": "World",
            "id": "c9feea0e-f8a4-49f4-af9b-92a045e10fc2"
          }
        ],
        "title": "Take a Bow",
        "text": "An emotional journey through heartfelt lyrics and powerful beats.",
        "link": "/play/Album/ec52c6d3-74b3-44f2-addf-0e1c92d90999"
      },
      {
        "id": "9d010fad-1b7d-4ecb-91c5-2ece623c1418",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/989/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Electronic",
            "id": "397fcd95-eb22-4383-9c56-6b92be6e8319"
          },
          {
            "name": "Funk",
            "id": "a7e80283-6edb-45ef-b3e5-3a28614feed5"
          }
        ],
        "title": "Come On-a My House",
        "text": "Feel the rhythm of timeless classics.",
        "link": "/play/Playlist/9d010fad-1b7d-4ecb-91c5-2ece623c1418"
      },
      {
        "id": "b434d7d8-9458-4c34-b6d9-e5266e935650",
        "type": "Podcast",
        "cover": {
          "src": "https://picsum.photos/id/459/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Country",
            "id": "9aa95290-d586-4a50-baea-3b0330309a08"
          },
          {
            "name": "Blues",
            "id": "22d958e5-8828-454b-8b2d-577ef0cfafd6"
          },
          {
            "name": "Latin",
            "id": "dce9c342-d389-4487-affe-da9e1f2db618"
          },
          {
            "name": "Soul",
            "id": "6df7e663-ee45-4915-a837-cd689f96c295"
          },
          {
            "name": "Country",
            "id": "d59a6928-2597-4805-a3ac-ceee08ee2aad"
          }
        ],
        "title": "Over the Rainbow",
        "text": "Experience the latest hits that are shaping the world of music.",
        "link": "/play/show/b434d7d8-9458-4c34-b6d9-e5266e935650"
      },
      {
        "id": "b5184ee4-12f2-4232-ab2d-0e64bfc54b14",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/447/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "World",
            "id": "f25f15e0-ecb8-46a6-9e02-3fd9f1261df9"
          },
          {
            "name": "Metal",
            "id": "2c50677d-e232-41ba-9222-49d8b4129564"
          }
        ],
        "title": "Time of the Season",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Album/b5184ee4-12f2-4232-ab2d-0e64bfc54b14"
      },
      {
        "id": "73aa167a-fd1b-4be5-90a4-0c02bf70da2e",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/867/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Folk",
            "id": "d6c6e324-bf41-4550-afa6-dc730110b3fe"
          }
        ],
        "title": "Tutti Frutti",
        "text": "Experience the latest hits that are shaping the world of music.",
        "link": "/play/Playlist/73aa167a-fd1b-4be5-90a4-0c02bf70da2e"
      },
      {
        "id": "9be1fd87-8b62-4bd7-ac5d-3e952f93af8f",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/274/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Latin",
            "id": "f63daa36-eafd-4918-af63-8e9102aef268"
          },
          {
            "name": "Pop",
            "id": "8314d30f-a702-49b8-a0fc-8b3ecf0402b6"
          },
          {
            "name": "Electronic",
            "id": "a4c77705-3786-4ea3-b5d3-33e7fa91c8e9"
          }
        ],
        "title": "Rag Doll",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Album/9be1fd87-8b62-4bd7-ac5d-3e952f93af8f"
      }
    ]
  },
  {
    "id": "2d29783c-6f2f-4bbb-95e7-faab2a8d512d",
    "title": "Classics Reimagined",
    "text": "Feel the power of music through every beat.",
    "link": "/play/section/2d29783c-6f2f-4bbb-95e7-faab2a8d512d",
    "type": "Podcast",
    "boxes": [
      {
        "id": "c1721dcb-45c8-4baa-8eeb-93b9a9f28f5d",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/531/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Metal",
            "id": "5afa58c5-bf8a-4520-b20d-73db467ab9e4"
          }
        ],
        "title": "Crazy",
        "text": "An emotional journey through heartfelt lyrics and powerful beats.",
        "link": "/play/Playlist/c1721dcb-45c8-4baa-8eeb-93b9a9f28f5d"
      },
      {
        "id": "abb60d80-2fcf-4105-bc2a-8c9022cc9a7d",
        "type": "Podcast",
        "cover": {
          "src": "https://picsum.photos/id/173/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Soul",
            "id": "ff41baf0-cf39-4bc8-ab2c-6090048785d8"
          },
          {
            "name": "Hip Hop",
            "id": "d2e51f67-782a-45ee-b873-05e7ad9c355a"
          }
        ],
        "title": "Down Hearted Blues",
        "text": "Immerse yourself in the sounds that defined an era.",
        "link": "/play/show/abb60d80-2fcf-4105-bc2a-8c9022cc9a7d"
      },
      {
        "id": "87d7bf90-81f7-4e50-a450-5fb5c6984962",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/859/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Jazz",
            "id": "4f127a29-c4cc-4078-aff1-3b0d5725b9ab"
          },
          {
            "name": "Pop",
            "id": "429dfa4d-a36e-4ff0-8909-8b556b9e9e6e"
          }
        ],
        "title": "Never Gonna Give You Up",
        "text": "Discover the artists pushing the boundaries of music.",
        "link": "/play/Playlist/87d7bf90-81f7-4e50-a450-5fb5c6984962"
      },
      {
        "id": "5c0dc219-4843-419d-b3a0-c1a39cd350a5",
        "type": "Act",
        "cover": {
          "src": "https://picsum.photos/id/858/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Funk",
            "id": "951cad11-b8f2-498c-a955-cb61d3274583"
          }
        ],
        "title": "Time of the Season",
        "text": "Feel the rhythm of timeless classics.",
        "link": "/play/Act/5c0dc219-4843-419d-b3a0-c1a39cd350a5"
      },
      {
        "id": "330e719d-74e7-4041-9b75-1daa781934fc",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/9/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Blues",
            "id": "72561cc6-fff4-4cd2-942c-34be12dd64e5"
          },
          {
            "name": "Folk",
            "id": "8610ab2b-fd31-4334-a5a4-afc7c637babd"
          },
          {
            "name": "Country",
            "id": "62695252-1f8a-4f15-97a5-7a21daa28920"
          },
          {
            "name": "Reggae",
            "id": "7b65809c-c673-4d3d-a9e4-4c09a3cc8fa7"
          }
        ],
        "title": "House of the Rising Sun",
        "text": "Feel the rhythm of timeless classics.",
        "link": "/play/Playlist/330e719d-74e7-4041-9b75-1daa781934fc"
      },
      {
        "id": "a43bccf4-59e8-4dd6-ac98-857959000257",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/939/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "World",
            "id": "67c99780-e22b-423d-9aca-91839ee01684"
          },
          {
            "name": "Reggae",
            "id": "17cca84e-30f5-4aba-bf4c-0d699db38dc4"
          },
          {
            "name": "Reggae",
            "id": "86df6951-e2ce-4912-a1f8-45d9023786b4"
          },
          {
            "name": "Electronic",
            "id": "52c1663a-dfb4-40e8-a92f-602f76ac13bc"
          }
        ],
        "title": "Honky Tonk",
        "text": "Immerse yourself in the sounds that defined an era.",
        "link": "/play/Playlist/a43bccf4-59e8-4dd6-ac98-857959000257"
      },
      {
        "id": "08c46040-96a7-4521-b0a7-1b192284669f",
        "type": "Act",
        "cover": {
          "src": "https://picsum.photos/id/402/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Jazz",
            "id": "9646fa08-e876-45f6-aa46-c2e1cacead5e"
          },
          {
            "name": "Rock",
            "id": "b02df468-2e99-498b-a28b-bee8176fdfd4"
          },
          {
            "name": "Non Music",
            "id": "69888559-f05a-4974-970e-60975448df7c"
          },
          {
            "name": "Electronic",
            "id": "da519300-696c-433c-85da-001cb8f83c58"
          }
        ],
        "title": "Abracadabra",
        "text": "Feel the rhythm of timeless classics.",
        "link": "/play/Act/08c46040-96a7-4521-b0a7-1b192284669f"
      },
      {
        "id": "203af97f-fda5-4ef8-a79d-3780e1d28c50",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/953/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Funk",
            "id": "0d68fb72-af3a-4847-a34b-d0e9671e7675"
          },
          {
            "name": "Non Music",
            "id": "07418688-084e-4eed-8b18-0e5b96c5cfe3"
          },
          {
            "name": "Metal",
            "id": "e58e8e91-9914-4151-bde2-81098824876f"
          },
          {
            "name": "Metal",
            "id": "48f3b2ac-a50c-4e77-a7df-e39d8d6e57a5"
          }
        ],
        "title": "This Land is Your Land",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Album/203af97f-fda5-4ef8-a79d-3780e1d28c50"
      },
      {
        "id": "98ea5541-e93e-4f22-9235-d4cd69494076",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/687/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Funk",
            "id": "53890d7a-1bcc-417c-ad79-407900ba2108"
          },
          {
            "name": "Classical",
            "id": "25744e50-d8e5-41e6-8158-8b7661853929"
          },
          {
            "name": "Non Music",
            "id": "6c5a7d6a-25ba-4220-b8fb-b760d459614c"
          },
          {
            "name": "Electronic",
            "id": "6f623486-2251-4a68-af11-b1f49f8a8ceb"
          },
          {
            "name": "Rock",
            "id": "9e6e1586-188a-4463-93eb-c434653b2b3b"
          }
        ],
        "title": "Stormy Weather (Keeps Rainin' All the Time)",
        "text": "An emotional journey through heartfelt lyrics and powerful beats.",
        "link": "/play/Playlist/98ea5541-e93e-4f22-9235-d4cd69494076"
      },
      {
        "id": "646ef806-6e96-48e0-9585-68c759e9ce13",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/279/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Folk",
            "id": "d1bf3a0a-e9a1-4631-b571-8b740c788fe9"
          },
          {
            "name": "Hip Hop",
            "id": "bdaa3cd0-04a1-4058-a293-900d13b1a67a"
          },
          {
            "name": "Folk",
            "id": "6e6e2bcb-c7a8-4b2c-b23b-8bbd73a0fbe5"
          },
          {
            "name": "Electronic",
            "id": "b067ed77-e8b0-4bd6-83e1-ad5881104cce"
          }
        ],
        "title": "I'm Sorry",
        "text": "Experience the latest hits that are shaping the world of music.",
        "link": "/play/Album/646ef806-6e96-48e0-9585-68c759e9ce13"
      }
    ]
  },
  {
    "id": "1a0c6c81-8abb-40b2-9ce8-2c524a0c95ad",
    "title": "Timeless Sounds",
    "text": "Dive into a world of melodies and rhythms.",
    "link": "/play/section/1a0c6c81-8abb-40b2-9ce8-2c524a0c95ad",
    "type": "Album",
    "boxes": [
      {
        "id": "6dbb28e7-ca2b-4c10-b2bf-a5e1867f85f7",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/450/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Soul",
            "id": "3f28831c-8e61-44a5-bde1-2be25f58118c"
          },
          {
            "name": "Pop",
            "id": "d286c9fc-7ff1-4df0-8df4-68d9df15b460"
          }
        ],
        "title": "Cherish",
        "text": "An emotional journey through heartfelt lyrics and powerful beats.",
        "link": "/play/Album/6dbb28e7-ca2b-4c10-b2bf-a5e1867f85f7"
      },
      {
        "id": "7db28149-7127-4806-bc13-9b66f737bf9c",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/883/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Rock",
            "id": "1c2c389e-858d-4940-8c1c-892071936065"
          }
        ],
        "title": "Body & Soul",
        "text": "An emotional journey through heartfelt lyrics and powerful beats.",
        "link": "/play/Playlist/7db28149-7127-4806-bc13-9b66f737bf9c"
      },
      {
        "id": "ab803845-986a-4011-9713-437d6d3c0375",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/145/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Classical",
            "id": "67138c3a-1069-4c27-97cf-d5ee3c1bcc2b"
          },
          {
            "name": "Electronic",
            "id": "5870930e-f4f7-41ab-91fc-8ab5788741f5"
          }
        ],
        "title": "It's Still Rock 'n' Roll to Me",
        "text": "Feel the rhythm of timeless classics.",
        "link": "/play/Playlist/ab803845-986a-4011-9713-437d6d3c0375"
      },
      {
        "id": "2d974874-eefd-4a04-ab5a-8d754c867ebc",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/12/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Latin",
            "id": "62d2e3be-4d50-4346-8a4e-072a4c0529b4"
          },
          {
            "name": "Folk",
            "id": "015120b2-ead9-420b-a6b8-973af43ce346"
          },
          {
            "name": "Latin",
            "id": "81bd6152-ffd2-4603-991b-53cb605409cb"
          },
          {
            "name": "Metal",
            "id": "6ffe9d8a-3c7c-4a71-a46b-1bce882a03b8"
          }
        ],
        "title": "Another Brick in the Wall (part 2)",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Album/2d974874-eefd-4a04-ab5a-8d754c867ebc"
      },
      {
        "id": "e1c77d83-9adf-4fdd-915e-0a5657440a9f",
        "type": "Act",
        "cover": {
          "src": "https://picsum.photos/id/974/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Non Music",
            "id": "e7358308-4db3-4df5-b247-d4967d502a54"
          },
          {
            "name": "Non Music",
            "id": "dcd71eba-a89b-4799-8240-a5d5f92d58c2"
          },
          {
            "name": "Classical",
            "id": "5fdcbd7d-1736-4c12-9b52-56cb40802bdc"
          },
          {
            "name": "Blues",
            "id": "f77ce297-b2d0-4a02-8899-e8f254758114"
          },
          {
            "name": "Latin",
            "id": "e286292e-924c-4e5e-9fa8-f040d865079c"
          }
        ],
        "title": "Honky Tonk",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Act/e1c77d83-9adf-4fdd-915e-0a5657440a9f"
      },
      {
        "id": "5d3646db-9502-4d6f-a22a-7fbb3422a65a",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/772/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Stage And Screen",
            "id": "dafe4c56-35a2-4b83-b31c-435fa7015282"
          },
          {
            "name": "Metal",
            "id": "a788af02-9019-4ee5-b0d2-3f09858bd7e7"
          },
          {
            "name": "Stage And Screen",
            "id": "f06ba83e-39a8-4c00-a855-2a642a4a7b58"
          },
          {
            "name": "Non Music",
            "id": "ccb592d6-679e-4b8b-8b9d-ef0f0e2f9376"
          },
          {
            "name": "Country",
            "id": "aefd3785-5569-4199-afe3-188170ea8bee"
          }
        ],
        "title": "Up Up & Away",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Playlist/5d3646db-9502-4d6f-a22a-7fbb3422a65a"
      },
      {
        "id": "f4101ccc-7913-4ff9-a5bd-c69867ee5138",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/330/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Stage And Screen",
            "id": "e954f451-889d-4ca1-a8e7-60999fe033fb"
          },
          {
            "name": "Rap",
            "id": "5a0415de-9d33-4e7d-82ec-781c1d8c785e"
          },
          {
            "name": "Pop",
            "id": "b11c317c-dacc-41f2-beb1-5444a4527d58"
          }
        ],
        "title": "One of These Nights",
        "text": "Discover the artists pushing the boundaries of music.",
        "link": "/play/Album/f4101ccc-7913-4ff9-a5bd-c69867ee5138"
      },
      {
        "id": "7a20b183-9935-484d-bf8a-f163f4fbf21d",
        "type": "Act",
        "cover": {
          "src": "https://picsum.photos/id/344/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Rock",
            "id": "230e3321-9fa9-4580-b8bc-9449144ffaa1"
          },
          {
            "name": "Jazz",
            "id": "fde7abb7-c786-4cd0-aca3-7d2e124f11b2"
          },
          {
            "name": "Pop",
            "id": "3e14bda0-34a3-454d-bbda-6995cc5d5a7c"
          },
          {
            "name": "Electronic",
            "id": "540b5e0a-0996-43d0-abc5-948e1d4959f0"
          }
        ],
        "title": "The Boys of Summer",
        "text": "Experience the latest hits that are shaping the world of music.",
        "link": "/play/Act/7a20b183-9935-484d-bf8a-f163f4fbf21d"
      },
      {
        "id": "b7d227b7-9def-4548-b177-54c74a5d464a",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/736/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Jazz",
            "id": "f675abf6-049a-4d85-addf-069db7119173"
          },
          {
            "name": "Pop",
            "id": "9961fce9-590c-4b0f-834e-e0252ac908d1"
          },
          {
            "name": "Rap",
            "id": "f2b7f927-d70c-4588-8812-118990d0acfa"
          },
          {
            "name": "Jazz",
            "id": "cdd7e136-590c-4160-9831-06d8933651d2"
          },
          {
            "name": "Hip Hop",
            "id": "01bb42a5-a96b-4433-b7b5-6392b107a28c"
          }
        ],
        "title": "Paper Doll",
        "text": "Feel the rhythm of timeless classics.",
        "link": "/play/Playlist/b7d227b7-9def-4548-b177-54c74a5d464a"
      },
      {
        "id": "4503e117-293e-463b-aab5-d7dca193814a",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/580/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Stage And Screen",
            "id": "b8e4675a-1df6-4a5f-b0dc-9eb6a86d7b02"
          },
          {
            "name": "Rap",
            "id": "deaeb6ec-bbc8-42df-8e41-4baf2ff7a7ab"
          },
          {
            "name": "Funk",
            "id": "b0249744-2dda-437c-aec5-38c6b237d348"
          }
        ],
        "title": "Nature Boy",
        "text": "Experience the latest hits that are shaping the world of music.",
        "link": "/play/Playlist/4503e117-293e-463b-aab5-d7dca193814a"
      }
    ]
  },
  {
    "id": "1cce2930-859a-48bd-82f6-49f5f1ec6e64",
    "title": "Hits of the Moment",
    "text": "Rediscover your favorite songs and artists in a new way.",
    "link": "/play/section/1cce2930-859a-48bd-82f6-49f5f1ec6e64",
    "type": "Playlist",
    "boxes": [
      {
        "id": "5ceeb6f4-b91e-4274-8915-67ffb9708a7f",
        "type": "Podcast",
        "cover": {
          "src": "https://picsum.photos/id/411/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Reggae",
            "id": "57901190-4b98-4bb0-91b2-09ba61e33da6"
          }
        ],
        "title": "The Letter",
        "text": "Discover the artists pushing the boundaries of music.",
        "link": "/play/show/5ceeb6f4-b91e-4274-8915-67ffb9708a7f"
      },
      {
        "id": "f450370a-e3e6-4434-8829-8c38e44cbcae",
        "type": "Podcast",
        "cover": {
          "src": "https://picsum.photos/id/339/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Pop",
            "id": "da497e53-c9c3-4412-a255-583857c9d85b"
          },
          {
            "name": "Soul",
            "id": "c6ea6dc2-8cc1-4e77-904f-65f660f92623"
          },
          {
            "name": "Hip Hop",
            "id": "b294caef-f323-4a19-9d0d-ba2dd6a36769"
          },
          {
            "name": "Electronic",
            "id": "6f8041a1-7cc9-42ba-978a-887617b8b500"
          },
          {
            "name": "Metal",
            "id": "000aa1e2-2d1b-4d67-96ab-2d9b7c343301"
          }
        ],
        "title": "Love Hangover",
        "text": "Immerse yourself in the sounds that defined an era.",
        "link": "/play/show/f450370a-e3e6-4434-8829-8c38e44cbcae"
      },
      {
        "id": "9b94184c-64bf-4334-92d2-44cca6d4f49a",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/183/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Metal",
            "id": "ba88e4fe-a8a9-4a8a-b2e5-809affcab7bb"
          },
          {
            "name": "Folk",
            "id": "c2ea30b1-629c-4b52-b225-d7be845a7d59"
          },
          {
            "name": "Pop",
            "id": "b4baf95d-8d6d-4d35-bdef-80acb3808dc8"
          },
          {
            "name": "Folk",
            "id": "7fdc8bba-52ab-4a78-beb5-ede553e6af77"
          },
          {
            "name": "Rock",
            "id": "db4f57f1-cfcc-4eba-81d6-0d4ee9cf563b"
          }
        ],
        "title": "Cherry Pink & Apple Blossom White",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Playlist/9b94184c-64bf-4334-92d2-44cca6d4f49a"
      },
      {
        "id": "f0705c14-caf7-468d-b61c-0d77025cfecf",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/392/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Non Music",
            "id": "e743b8b4-8126-45af-9d6b-a8748f078824"
          },
          {
            "name": "Hip Hop",
            "id": "149fb985-ffe3-4c6d-89b7-30b409ec1834"
          }
        ],
        "title": "Born in the USA",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Playlist/f0705c14-caf7-468d-b61c-0d77025cfecf"
      },
      {
        "id": "1594c874-8e3d-440c-889a-4dba1220ab07",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/888/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Funk",
            "id": "057326bd-c258-4efb-99e3-eca8a79a2ad6"
          },
          {
            "name": "Non Music",
            "id": "b404e60f-91f7-4c8d-ae0c-f810bcf415f9"
          },
          {
            "name": "Reggae",
            "id": "3bdbf5d5-be53-4eca-aec7-b00ff0d12922"
          },
          {
            "name": "Hip Hop",
            "id": "64e85acc-e109-40b4-9190-0f52b0fe8c48"
          },
          {
            "name": "Jazz",
            "id": "c862ab89-a24e-417a-8dfc-07cb69a2f259"
          }
        ],
        "title": "Thank You (Falettinme be Mice Elf Again)",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Playlist/1594c874-8e3d-440c-889a-4dba1220ab07"
      },
      {
        "id": "c458cd91-7369-40cb-aff0-7804de8e056e",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/936/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Classical",
            "id": "91dfe6e0-cfb6-466d-9b52-537e37315969"
          },
          {
            "name": "Latin",
            "id": "e01e9494-30a8-4799-bd55-93e3f246e3d5"
          }
        ],
        "title": "The Long & Winding Road",
        "text": "Immerse yourself in the sounds that defined an era.",
        "link": "/play/Album/c458cd91-7369-40cb-aff0-7804de8e056e"
      },
      {
        "id": "a35683e2-d46a-43fb-80a2-8c8599972e22",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/381/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Reggae",
            "id": "b5a856d9-b4bb-43b5-b8f6-d56357cc1dfd"
          },
          {
            "name": "Blues",
            "id": "783ab228-3d78-457a-9bdc-bc3c40d397c7"
          },
          {
            "name": "Reggae",
            "id": "5ff3ea3e-934f-437c-9670-4e05379b7d07"
          },
          {
            "name": "Pop",
            "id": "5c146eb4-aeed-4b28-8c09-bde14b80554f"
          },
          {
            "name": "Latin",
            "id": "c1125bf3-a4b7-4248-83e1-24daefd28963"
          }
        ],
        "title": "All Out of Love",
        "text": "An emotional journey through heartfelt lyrics and powerful beats.",
        "link": "/play/Album/a35683e2-d46a-43fb-80a2-8c8599972e22"
      },
      {
        "id": "454c59d9-4fd6-4e51-adaf-0ac52ff93e53",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/636/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Non Music",
            "id": "d4c14705-3f8f-44f6-a9d7-9ffde9060d03"
          },
          {
            "name": "Reggae",
            "id": "1acdd555-90e5-45dd-81d8-54d802bc4a3d"
          },
          {
            "name": "Pop",
            "id": "02ba7d71-2351-4ec4-baa3-db7ab7d4083b"
          },
          {
            "name": "Reggae",
            "id": "c8346c8d-5d18-40dd-b894-cd0dc096d9f6"
          }
        ],
        "title": "White Rabbit",
        "text": "Experience the latest hits that are shaping the world of music.",
        "link": "/play/Playlist/454c59d9-4fd6-4e51-adaf-0ac52ff93e53"
      },
      {
        "id": "d084a9ac-d290-4472-b953-8f49d6fbcf57",
        "type": "Act",
        "cover": {
          "src": "https://picsum.photos/id/552/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "World",
            "id": "e62149b0-96ab-4796-b7ae-af787310ac7f"
          },
          {
            "name": "Hip Hop",
            "id": "be2ff222-87c7-432c-858a-d1be36813742"
          }
        ],
        "title": "It's Too Late",
        "text": "Discover the artists pushing the boundaries of music.",
        "link": "/play/Act/d084a9ac-d290-4472-b953-8f49d6fbcf57"
      },
      {
        "id": "a3f18f48-5fbd-4e0a-8dab-a6b121d7b748",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/240/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Jazz",
            "id": "c2f8b50a-16cd-443e-801e-7b3b2b2d1090"
          },
          {
            "name": "Hip Hop",
            "id": "55bc3abb-af90-44ae-b982-f66452bad78c"
          },
          {
            "name": "Rock",
            "id": "c77f4f3b-abce-4180-b782-79162c0aa642"
          },
          {
            "name": "Latin",
            "id": "8e68a65d-9256-4f1a-b0ef-93f4fced5fe9"
          },
          {
            "name": "Rock",
            "id": "351f0113-0d24-4596-871e-316cabf02889"
          }
        ],
        "title": "Best of My Love",
        "text": "An eclectic mix of melodies for every mood.",
        "link": "/play/Album/a3f18f48-5fbd-4e0a-8dab-a6b121d7b748"
      }
    ]
  },
  {
    "id": "3a6851f8-fc18-406d-8ed5-f0af3c9b8ec1",
    "title": "Sounds of the Future",
    "text": "Rediscover your favorite songs and artists in a new way.",
    "link": "/play/section/3a6851f8-fc18-406d-8ed5-f0af3c9b8ec1",
    "type": "Album",
    "boxes": [
      {
        "id": "a8b995e0-ae36-4115-84f2-4ef7f26b65b7",
        "type": "Act",
        "cover": {
          "src": "https://picsum.photos/id/660/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Classical",
            "id": "7493381b-5015-46cb-8ab0-f7ee51841ee6"
          },
          {
            "name": "Metal",
            "id": "5b88c0ce-0a77-423d-9d0f-3a4a9d2a962c"
          },
          {
            "name": "World",
            "id": "edc43ba2-d67f-4852-80be-589aacd4922f"
          },
          {
            "name": "Country",
            "id": "6515c678-d1d3-4789-856f-b8930f615207"
          },
          {
            "name": "Non Music",
            "id": "79bac415-9ff6-4e0c-a937-2190b3036a69"
          }
        ],
        "title": "Penny Lane",
        "text": "Experience the latest hits that are shaping the world of music.",
        "link": "/play/Act/a8b995e0-ae36-4115-84f2-4ef7f26b65b7"
      },
      {
        "id": "243d82c4-7ab6-477d-acc6-1c27e129558a",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/913/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Hip Hop",
            "id": "4ab4c345-4135-4784-8229-15311684ccc5"
          },
          {
            "name": "Metal",
            "id": "4e41626f-4746-469e-89d1-e964bdf8fc4b"
          },
          {
            "name": "Folk",
            "id": "738318f9-94c9-4e67-8eb6-be509a5652b1"
          },
          {
            "name": "Hip Hop",
            "id": "bb05c3d9-df0b-4c2a-9a3d-8f64321344b4"
          },
          {
            "name": "Pop",
            "id": "dd1d1311-00fa-473b-8337-933d8cec4de2"
          }
        ],
        "title": "Cold",
        "text": "Discover the artists pushing the boundaries of music.",
        "link": "/play/Playlist/243d82c4-7ab6-477d-acc6-1c27e129558a"
      },
      {
        "id": "8e8a3c27-0984-4ef3-a6ab-2f8ff5d63ac2",
        "type": "Act",
        "cover": {
          "src": "https://picsum.photos/id/418/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Pop",
            "id": "99345e53-1824-4bb0-9890-949038640cc9"
          },
          {
            "name": "Rock",
            "id": "d9b39c89-d77a-4bbe-9e15-db15e3b29690"
          },
          {
            "name": "World",
            "id": "34853784-14a8-4c19-9fb9-8b7a029ffe85"
          },
          {
            "name": "Hip Hop",
            "id": "de33f099-97c5-4572-b139-3aa2474fbca5"
          },
          {
            "name": "Reggae",
            "id": "34cae5ce-56d6-4a51-8d22-3b08b2ddd7c3"
          }
        ],
        "title": "Mama Told Me Not to Come",
        "text": "Experience the latest hits that are shaping the world of music.",
        "link": "/play/Act/8e8a3c27-0984-4ef3-a6ab-2f8ff5d63ac2"
      },
      {
        "id": "d5675376-3d45-4db0-a440-24cc9da72fbb",
        "type": "Act",
        "cover": {
          "src": "https://picsum.photos/id/357/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Jazz",
            "id": "ee7c366a-189b-45ce-964c-f7964cc1ca5a"
          },
          {
            "name": "Funk",
            "id": "084ebc96-a2fb-4406-9a0b-36c10d29d063"
          },
          {
            "name": "Reggae",
            "id": "d265e145-a79b-492b-9d28-2720852319c3"
          },
          {
            "name": "World",
            "id": "b476c624-1233-44de-8fc9-03e8ed896d69"
          }
        ],
        "title": "(Sittin' On) the Dock of the Bay",
        "text": "Feel the rhythm of timeless classics.",
        "link": "/play/Act/d5675376-3d45-4db0-a440-24cc9da72fbb"
      },
      {
        "id": "c413961b-2c46-4b45-b44c-0f3569d05096",
        "type": "Album",
        "cover": {
          "src": "https://picsum.photos/id/527/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Metal",
            "id": "db56b910-1a30-461d-ad38-c009f5f28612"
          },
          {
            "name": "Metal",
            "id": "6f61c54f-9577-45ae-b825-78a537b05349"
          },
          {
            "name": "Latin",
            "id": "2881799b-1f83-4037-bb82-26ddf14be65b"
          },
          {
            "name": "Non Music",
            "id": "71ac06a6-44aa-45d4-87f0-0a851cf82aab"
          }
        ],
        "title": "Come On Eileen",
        "text": "An emotional journey through heartfelt lyrics and powerful beats.",
        "link": "/play/Album/c413961b-2c46-4b45-b44c-0f3569d05096"
      },
      {
        "id": "6a423f63-7b4a-46cf-9a8b-16c33a94a817",
        "type": "Playlist",
        "cover": {
          "src": "https://picsum.photos/id/851/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Funk",
            "id": "dd434291-c62a-45ae-a374-2ccc549b275a"
          },
          {
            "name": "Rock",
            "id": "62b58f76-3f41-44d3-81a6-5c9ddeefaf6e"
          },
          {
            "name": "Funk",
            "id": "b3529380-7d72-4525-b413-407bb0f850bc"
          },
          {
            "name": "Blues",
            "id": "6897ddeb-af0f-40ea-b5f3-6541377da1c1"
          }
        ],
        "title": "Call Me",
        "text": "Discover the artists pushing the boundaries of music.",
        "link": "/play/Playlist/6a423f63-7b4a-46cf-9a8b-16c33a94a817"
      },
      {
        "id": "afe269cc-f839-4c0d-aa35-881ec5f6d166",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/242/200",
          "type": "Circle"
        },
        "playlists": [
          {
            "name": "Classical",
            "id": "2d3b1205-2175-4771-b233-a66ae38b5f5d"
          },
          {
            "name": "Non Music",
            "id": "b47333e3-ce62-4624-b67a-cdcb0b39b8cd"
          },
          {
            "name": "World",
            "id": "57d2f24c-8948-4280-a795-9925a0a01011"
          },
          {
            "name": "Non Music",
            "id": "74e54134-4d74-42a0-b288-b4537ef3b2b5"
          },
          {
            "name": "Soul",
            "id": "90bca743-4574-480b-87f1-f824c1ab4373"
          }
        ],
        "title": "I Love Rock 'n' Roll",
        "text": "Immerse yourself in the sounds that defined an era.",
        "link": "/play/Playlist/afe269cc-f839-4c0d-aa35-881ec5f6d166"
      },
      {
        "id": "04bdea86-b4b6-40d9-8cdb-2fa53bcea854",
        "type": "Podcast",
        "cover": {
          "src": "https://picsum.photos/id/396/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "World",
            "id": "ca155870-677d-4036-a9ac-1b6a7b375d29"
          }
        ],
        "title": "Frenesi",
        "text": "Immerse yourself in the sounds that defined an era.",
        "link": "/play/show/04bdea86-b4b6-40d9-8cdb-2fa53bcea854"
      },
      {
        "id": "d585941c-2239-41d9-94f1-42798a67e6bc",
        "type": "Podcast",
        "cover": {
          "src": "https://picsum.photos/id/478/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Country",
            "id": "cf13802e-a268-4c4c-8529-cfaa9442e49b"
          },
          {
            "name": "Soul",
            "id": "ba047d38-5254-4f67-b2a8-a2b167fa71bb"
          },
          {
            "name": "Country",
            "id": "6142903d-6daa-4608-8006-de8b7ff9b86a"
          },
          {
            "name": "Funk",
            "id": "381d8662-282c-4143-901a-3a377750e788"
          },
          {
            "name": "Electronic",
            "id": "6b882fee-53f5-4a85-bb38-c33cca2ffe03"
          }
        ],
        "title": "50 Ways to Leave Your Lover",
        "text": "Discover the artists pushing the boundaries of music.",
        "link": "/play/show/d585941c-2239-41d9-94f1-42798a67e6bc"
      },
      {
        "id": "6ee1c81c-ebfc-4351-9021-098ef17f682c",
        "type": "Radio",
        "cover": {
          "src": "https://picsum.photos/id/559/200",
          "type": "Square"
        },
        "playlists": [
          {
            "name": "Blues",
            "id": "d869d24b-febb-473f-9384-d7221b9f6954"
          }
        ],
        "title": "Sweet Dreams (Are Made of This)",
        "text": "Discover the artists pushing the boundaries of music.",
        "link": "/play/Playlist/6ee1c81c-ebfc-4351-9021-098ef17f682c"
      }
    ]
  }
];
