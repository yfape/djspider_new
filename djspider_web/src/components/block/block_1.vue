<!--
 * @Author: your name
 * @Date: 2021-07-25 22:39:19
 * @LastEditTime: 2021-07-25 23:16:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\components\data\block_1.vue
-->
<template>
<q-item clickable v-ripple class="main bg-white">
    <q-item-section avatar>
        <slot name="icon"></slot>
    </q-item-section>
    <q-item-section>
        <div class="text-subtitle2 text-grey-9">{{title}}</div>
        <div class="text-h5 text-bold text-red-7">{{num}}</div>
    </q-item-section>
</q-item>
</template>
<script>

export default{
name:'block_1',
components:{},
props:['title', 'num'],
data(){return {

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
.main{width:200px;height:80px;border-radius:5px}
</style>