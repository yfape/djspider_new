<!--
 * @Author: your name
 * @Date: 2021-07-22 22:23:34
 * @LastEditTime: 2021-08-03 13:53:58
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\components\user_avater.vue
-->
<template>
<q-item class="main shadow-2 cursor-pointer" clickable v-ripple>
    <q-item-section top avatar>
        <q-avatar rounded>
            <img :src="user.Headimg">
        </q-avatar>
    </q-item-section>
    <q-item-section class="text-left">
        <q-item-label>{{user.Name}}</q-item-label>
        <q-item-label caption>{{user.Email}}</q-item-label>
    </q-item-section>
    <q-item-section side top>
        <q-item-label caption>{{ isNaN(parseInt(user.Identify)) ? user.Identify : identifys[user.Identify]}}</q-item-label>
    </q-item-section>
</q-item>
</template>
<script>

export default{
name:'user_avater',
components:{},
props:['user'],
data(){return {
    identifys:["未知","用户","管理员"]
}},
mounted(){

},
methods:{

},
watch:{},
destroy(){},
}
</script>
<style scoped>
.main{width:300px;border-radius:5px;background:#fff}
</style>