<!--
 * @Author: your name
 * @Date: 2021-07-25 19:56:01
 * @LastEditTime: 2021-08-03 16:32:25
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\views\ManageUser.vue
-->
<template><div class='q-pa-md'>
    <div class="full-width bg-white q-pa-sm row items-center justify-start">
        <q-btn color="green-7" icon="o_add" label="新增用户" class="q-mr-md" @click="addUser"/>
        <q-space/>
        <q-input dense filled label="搜索" color="red-7" clearable v-model="search">
            <template v-slot:prepend>
                <q-icon name="o_search"/>
            </template>
        </q-input>
    </div>
    <div class="row items-center justify-start q-gutter-sm q-mt-sm" v-if="show">
        <UserAvater v-for="item in users" :user="item" :key="item.UserId" @click="openPanel(item)"/>
    </div>
    <div v-if="!show" class="q-ma-xl text-grey-7">暂无用户数据</div>
    <UserPanel ref="userpanel" @change="getUser"/>
</div></template>
<script>
import UserAvater from '../components/user_avater'
import UserPanel from '../components/user_panel'
export default{
name:'ManageUser',
components:{UserAvater,UserPanel},
props:[],
data(){return {
    search:"",users:[],show:false,
}},
mounted(){
    this.show = false
    this.getUser()
},
methods:{
    getUser(){
        this.axios.get("/users").then(res => {
            this.users = Object.assign([],[],res.users)
            this.show = true
        })
    },
    openPanel(user){
        this.$refs.userpanel.showPanel(user)
    },
    addUser(){
        this.$refs.userpanel.showPanel()
    }
},
watch:{},
destroy(){},
}
</script>
<style scoped>

</style>