<template>
  <div class="main-layout">
  </div>
  <div class="box-area">
    <loginform />
  </div>
</template>

<script>
import Loginform from "../../components/login/index.vue";
export default {
  components: {Loginform}
}
</script>

<style scoped>
.main-layout{
  width: 100vw;
  height: 100vh;
  background: url("/861816.jpg") no-repeat center fixed;
  background-size:100% 100%;
  background-size: cover;
  overflow: hidden;
}

</style>
