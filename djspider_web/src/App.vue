<!--
 * @Author: your name
 * @Date: 2021-07-22 15:37:24
 * @LastEditTime: 2021-07-26 11:10:49
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\App.vue
-->
<template>
  <router-view/>
</template>

<script>
// window.onbeforeunload = function(){ return window.event.returnValue = "222"; }
export default{
  name: 'App'
}
</script>

<style lang="scss">
html,body{height:100%;width:100%;padding:0;margin:0;overflow:hidden}
#app {
  font-family:'黑体', Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height:100%;width:100%;
}
.main{height:100%;width:100%;}
</style>
