<!--
 * @Author: your name
 * @Date: 2021-08-09 15:49:23
 * @LastEditTime: 2021-08-09 17:12:16
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\components\panel_log.vue
-->
<template>
<q-dialog v-model="show" @hide="hidePanel">
    <div v-if="inshow" class="bg-white q-pa-sm full-width">
        <div class="text-subtitle1 text-grey-7">今日日志</div>
        <q-separator class="q-mb-sm"/>
        <div class="bg-black text-white" style="white-space: pre-line;height:500px;overflow-y:auto;overflow-x:hidden" v-html="content"></div>
    </div>
    <q-spinner-bars v-else color="white" size="2em"/>
</q-dialog>
</template>
<script>
export default{
name: '',
data(){return {
    show:false,inshow:false,content:""
}},
mounted(){},
methods:{
    showPanel(){
        this.show = true
        this.axios.get("/log").then(res =>{
            this.content = res.data
            this.inshow = true
        }).catch(()=>{
            this.show = false
        })
    },
    hidePanel(){
        this.inshow = false
    }
},
}
</script>
<style scoped>
</style>
