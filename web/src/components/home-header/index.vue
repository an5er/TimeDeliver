<template xmlns="http://www.w3.org/1999/html">
  <div class="home-page">
    <el-menu
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        :ellipsis="false"
        @select="handleSelect"
        style="background: #141414;border: outset #373738;margin: auto"
    >
      <div class="flex-grow" />
      <el-menu-item index="1" style="color: #fff;">TimeDeliver</el-menu-item>
    </el-menu>

  </div>

</template>

<script>
export default {
  name: "home-header"
}
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
</style>
