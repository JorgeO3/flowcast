export const useScroll = () => {
  const internalX = useState("x_axis_scroll", () => 0);
  const internalY = useState("y_axis_scroll", () => 0);

  const update = () => {
    internalX.value = window.scrollX;
    internalY.value = window.scrollY;
  };

  const x = computed(() => internalX.value);
  const y = computed(() => internalY.value);

  onMounted(() => window.addEventListener("scroll", update));
  onUnmounted(() => window.removeEventListener("scroll", update));

  return { x, y };
};