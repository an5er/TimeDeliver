<template xmlns="http://www.w3.org/1999/html">
  <div class="home-page">
    <el-menu
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        :ellipsis="false"
        @select="handleSelect"
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
@font-face {
  font-family: alohabrush;
  src: url("/font/AlohaBrush-DemoVersion-Regular.ttf");
  font-weight: normal;
  font-style: normal;
}

.flex-grow {
  flex-grow: 1;
}
.home-page{
  font-family: alohabrush;
}
.el-menu-demo{
  background: #414243;
  border: outset #373738;
  margin: auto;
  opacity:0.5;
}
</style>
