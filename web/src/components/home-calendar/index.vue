<template>
  <div class="top-con">
    <div class="top" v-for="item in top" :key="item">{{ item }}</div>
  </div>
  <div class="box-wrap">
   <div
        class="date"
        :class="[item.thisMonth,item.isToday]"
        v-for="(item, index) in visibleCalendar"
        :key="index"
    >
     <transition name="el-fade-in-linear">
      <div class="box-card" v-tilt="{glare: true, maxGlare: 1, max: 30, speed: 400}">
          <div class="card-header">
            <span>{{item.day}}</span>
          </div>
          <div class="card-body">
            <template v-for="thing in things">
              <el-tag v-if="thing.Deadline.toString().split('T')[0].split('-')[2] == add0(item.day)
              && thing.Deadline.toString().split('T')[0].split('-')[1] == add0(item.month+1)">
                {{thing.Thing}}
              </el-tag>
            </template>
          </div>
      </div>
     </transition>
    </div>


  </div>
</template>

<script>
import {getCurrentInstance} from "vue";
import config_ from "../config/index.vue";
import {ParticlesComponent} from "particles.vue3";

export default {
  name: "home-calendar",
  components: {ParticlesComponent},
  data() {
    return{
      top: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
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
    },
    generateCalendarClassName(index){
      if ((index+1)%7===0){
        return "sundaycard";
      }else {
        return "normalcard"
      }
    },
    add0(m){
      return m<10?'0'+m:m
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
  },
  computed: {
      visibleCalendar: function() {
      const today = new Date()
      today.setHours(0)
      today.setMinutes(0)
      today.setSeconds(0)
      today.setMilliseconds(0)

      const calendarArr = []
      const currentFirstDay = new Date(today.getFullYear(), today.getMonth(), 1)
      const weekDay =
          currentFirstDay.getDay() === 0 ? 7 : currentFirstDay.getDay()
      const startDay = currentFirstDay - (weekDay - 1) * 24 * 3600 * 1000
      for (let i = 0; i < 42; i++) {
        const date = new Date(startDay + i * 24 * 3600 * 1000)
        calendarArr.push({
          date: new Date(startDay + i * 24 * 3600 * 1000),
          year: date.getFullYear(),
          month: date.getMonth(),
          day: date.getDate(),
          thisMonth:
              date.getFullYear() === today.getFullYear() &&
              date.getMonth() === today.getMonth()
                  ? 'thisMonth'
                  : '',
          isToday:
              date.getFullYear() === today.getFullYear() &&
              date.getMonth() === today.getMonth() &&
              date.getDate() === today.getDate()
                  ? 'isToday'
                  : '',
        })
      }
      return calendarArr
    }
  }
}

</script>

<style>
@font-face {
  font-family: tulisans;
  src: url("/font/Tulisans.ttf");
  font-weight: normal;
  font-style: normal;
}
@font-face {
  font-family: couplevalentinettfpersonal;
  src: url("/font/CoupleValentineTTFPersonal.ttf");
  font-weight: normal;
  font-style: normal;
}
@font-face {
  font-family: harmonyossansregular;
  src: url("/font/HarmonyOS_Sans_Regular.ttf");
  font-weight: normal;
  font-style: normal;
}


.card-header{
  font-family: couplevalentinettfpersonal;
  width: 30px;
  height: 30px;
  border-radius:10px;
}
.card-body{
  font-family: harmonyossansregular;
  float:right;
  width:100%;
  height:100%;
}
.top-con {
  font-family: tulisans;
  display: flex;
  align-items: center;
}
.top-con .top {
  border-radius: 5px;
  width: 25%;
  opacity:0.9;
  background: #CFD3DC;
  padding: 10px 0;
  margin: 5px;
  text-align: center;
}
.box-card{
  width: 207px;
  height: 96px;
  background: white;
  opacity:0.6;
  border-radius:10px;
  box-sizing:border-box;
}
.thisMonth .box-card{
  background-color: rgb(169, 225, 243);
  opacity:0.8;
}
.thisMonth.isToday .box-card{
  background-color: rgb(125, 230, 245);
  opacity:0.8;
}
.date {
  display: flex;
  width: 210px;
  margin-right: 0px;
  margin-left: 5px;
}
.box-wrap{
  display: grid;
  grid-gap: 5px 0px;
  grid-template-columns: repeat(7, 219.5px);
}
</style>
