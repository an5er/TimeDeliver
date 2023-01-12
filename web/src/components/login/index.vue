<template>
  <div class="loginPart">
    <h2>用户登录</h2>
  <el-form
      label-position="left"
      label-width="70px"
      :model="form"
      style="max-width: 300px;margin: 0 auto;"
  >
    <el-form-item label="用户名" class="form-item">
      <el-input v-model="form.uname" />
    </el-form-item>
    <el-form-item label="密码" class="form-item">
      <el-input v-model="form.passwd" />
    </el-form-item>
  </el-form>
  <el-row>
    <el-button class="login-button" type="primary" style="margin: 0 auto;" @click="ToLogin">登录</el-button>
  </el-row>
  </div>
</template>

<script>
import config_ from "../config/index.vue"
import {getCurrentInstance,} from "vue";
import {ElMessage} from "element-plus";


export default {
  name: "login-component",
  data() {
    return{
      form:{
        uname: "",
        passwd: "",
      }
    }
  },
  methods: {
    ToLogin: function (){
      console.log(this.axios)
      this.axios.post(`${config_.apiAddress}/user/login`,{
        "uname": this.form.uname,
        "passwd": this.form.passwd,
      },{
        withCredentials: true,
        credentials: 'include',
      }).then(function (data) {
        console.log(data.data.code === 20)
        if (data.data.code === 20){
          window.sessionStorage.setItem("token",data.data.data.token)
        }

        ElMessage({
          showClose: true,
          message: '登录成功，正在跳转到主页',
          type: 'success',
        })
        setTimeout(function() { location.href = "/home"; }, 500);
      })
    }
  },
  setup(){
    const getAxios = ()=>{
      return getCurrentInstance().appContext.config.globalProperties.$http
    }
    return {
      getAxios
    }
  },
  mounted() {
    this.axios = this.getAxios();
    this.axios.defaults.withCredentials = true
  }
}
</script>

<style>
.loginPart{
  position:absolute;
  /*定位方式绝对定位absolute*/
  top:50%;
  left:50%;
  /*顶和高同时设置50%实现的是同时水平垂直居中效果*/
  transform:translate(-50%,-50%);
  /*实现块元素百分比下居中*/
  width:450px;
  padding:50px;
  background: rgba(0,0,0,.5);
  box-sizing:border-box;
  box-shadow: 0px 15px 25px rgba(0,0,0,.5);
  /*边框阴影  水平阴影0 垂直阴影15px 模糊25px 颜色黑色透明度0.5*/
  border-radius:15px;
  /*边框圆角，四个角均为15px*/
}
.loginPart h2{
  margin:0 0 30px;
  padding:0;
  color: #fff;
  text-align:center;
}
.form-item .el-form-item__label{
  color: #fff;
}
</style>
