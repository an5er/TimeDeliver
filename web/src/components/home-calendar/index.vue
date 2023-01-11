<template>

  <el-calendar v-model="value" class="calendar">
<!--    <template #date-cell="{ data }">-->
<!--      <el-card shadow="always" v-for="(Thing) in this.things">-->
<!--        {{ data.day.split('-')[2] }}-->
<!--        {{data.day.split('-')[2] == Thing.Deadline.split('T')[0].split('-')[2] ? Thing.Thing : '' }}-->
<!--      </el-card>-->
<!--    </template>-->
    <template>
      <div >
        <p>hello world</p>
      </div>
    </template>
  </el-calendar>
</template>

<script>
import {getCurrentInstance} from "vue";
import config_ from "../config/index.vue";

export default {
  name: "home-calendar",
  data() {
    return{
      value:new Date(),
      things:[],
    }
  },
  methods: {
    GetThings: function (){
      let that = this
      let month = (this.value.getMonth()+1) < 10?'0'+ (this.value.getMonth()+1):this.value.getMonth()+1;
      this.axios.post(`${config_.apiAddress}/todo/monthtodo`,{
        "year": this.value.getFullYear().toString(),
        "month": month.toString(),
      },{
        withCredentials: true,
        credentials: 'include',
        headers: {"x-token":window.sessionStorage.getItem("token")}
      }).then(function (data) {
        that.things = data.data.data
        // console.log(that.things)
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
    let that = this
    this.axios = this.getAxios();
    this.axios.defaults.withCredentials = true
    that.GetThings()
  }
}

</script>

<style>
.is-selected {
  color: #1989fa;
}
.calendar {
  width: auto;
  height: 1000px;
}
</style>
