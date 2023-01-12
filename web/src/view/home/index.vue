<template>
  <div class="particles">
    <particles-component/>
  </div>
  <div class="home-page">
    <div class="home-header">
      <home-header/>
    </div>
    <div class="home-calendat">
      <home-calendar/>
    </div>
  </div>
</template>

<script>
import HomeHeader from "../../components/home-header/index.vue";
import HomeCalendar from "../../components/home-calendar/index.vue";
import ParticlesComponent from "../../components/particles/index.vue";
export default {
  name: "home",
  components: {ParticlesComponent, HomeCalendar, HomeHeader},
}
</script>

<style scoped>
.home-page {
  height: 100%;
  padding: 0;
  margin: 0;
}
.home-calendat{
  height: auto;
}
.particles{
  position:absolute;
  z-index:-1;
}
</style>
