<template>
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
export default {
  name: "home",
  components: {HomeCalendar, HomeHeader}
}
</script>

<style scoped>
.home-page {
  height: 100%;
  padding: 0;
  margin: 0;
  background: #141414;
}
</style>
