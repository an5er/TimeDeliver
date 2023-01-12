<template>
  <Particles id="tsparticles" :particlesInit="particlesInit" :options="particlesOptions" />
</template>


<script>

import { loadStarsPreset } from "tsparticles-preset-stars";

export default {
  name: "particles-component",
  data: () => {
    return {
      particlesOptions: {

        preset: "stars",
      }
    }
  },
  methods: {
    particlesInit: async (engine) => {
      await loadStarsPreset(engine);
    },
  },
};
</script>

<style scoped>

</style>
