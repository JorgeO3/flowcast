import * as lucide from "lucide-vue-next";

export interface SidebarLink {
  icon: Component;
  text: string;
  path: string;
}

export interface SidebarSection {
  title: string;
  items: SidebarLink[];
}

const data: SidebarSection[] = [
  {
    title: 'Discover',
    items: [
      { icon: lucide.CirclePlayIcon, text: 'Listen Now', path: '/play' },
      { icon: lucide.LayoutGridIcon, text: 'Browse', path: '/play/browse' },
      { icon: lucide.RadioIcon, text: 'Radio', path: '/play/radio' },
    ],
  },
  {
    title: 'Library',
    items: [
      { icon: lucide.ListMusic, text: 'Playlists', path: '/play/playlists' },
      { icon: lucide.Music2Icon, text: 'Songs', path: '/play/songs' },
      { icon: lucide.UserIcon, text: 'Made for You', path: '/play/made-for-you' },
      { icon: lucide.MicVocalIcon, text: 'Artists', path: '/play/artists' },
      { icon: lucide.LibraryBigIcon, text: 'Albums', path: '/play/albums' },
    ],
  },
  {
    title: 'Playlists',
    items: [
      { icon: lucide.DiscIcon, text: 'Chill Mix', path: '/play/playlist/1' },
      { icon: lucide.DiscIcon, text: 'Workout Mix', path: '/play/playlist/2' },
      { icon: lucide.DiscIcon, text: 'Focus Mix', path: '/play/playlist/3' },
      { icon: lucide.DiscIcon, text: 'Party Mix', path: '/play/playlist/4' },
      { icon: lucide.DiscIcon, text: 'Study Mix', path: '/play/playlist/5' },
      { icon: lucide.DiscIcon, text: 'Sleep Mix', path: '/play/playlist/6' },
      { icon: lucide.DiscIcon, text: 'Relax Mix', path: '/play/playlist/7' },
      { icon: lucide.DiscIcon, text: 'Travel Mix', path: '/play/playlist/8' },
      { icon: lucide.DiscIcon, text: 'Road Trip Mix', path: '/play/playlist/9' },
      { icon: lucide.DiscIcon, text: 'Summer Mix', path: '/play/playlist/10' },
      { icon: lucide.DiscIcon, text: 'Winter Mix', path: '/play/playlist/11' },
    ],
  }
];

export const sectionsMusicSidebar = (): SidebarSection[]  => {
  return data;
}
