<!--
 * @Author: your name
 * @Date: 2021-07-25 19:38:19
 * @LastEditTime: 2021-08-03 12:41:52
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\views\User.vue
-->
<template><div>
    <div class="row q-mt-md q-px-md">
        <div class="col-12 col-sm-6 col-md-4">
            <div class="q-px-md q-py-sm bg-white text-left">
                <div class="text-caption text-grey-6 br">头像</div>
                <div class="row items-center justify-start">
                    <div class="col-3 q-pr-md"><q-img class="rounded-borders" :ratio="1" :src="$store.state.user.Headimg"/></div>
                    <div class="col-9"><q-input type="textarea" autogrow label="头像网络地址" color="red-7" dense outlined v-model="$store.state.user.Headimg" /></div>
                </div>
            </div>
        </div>
    </div>
    <div class="row q-mt-md q-px-md">
        <div class="col-12 col-sm-6 col-md-4">
            <div class="q-px-md q-py-sm bg-white text-left">
                <div class="text-caption text-grey-6 br">基本信息</div>
                <q-input class="q-mb-sm" label="姓名" color="red-7" dense outlined v-model="$store.state.user.Name" />
                <q-input class="q-mb-sm" label="邮箱" color="red-7" dense outlined v-model="$store.state.user.Email" />
                <q-input label="权限" color="red-7" dense outlined readonly v-model="$store.state.user.Identify" />
            </div>
        </div>
    </div>
    <div class="row q-mt-md q-px-md">
        <div class="col-12 col-sm-6 col-md-4">
            <div class="q-px-md q-py-sm bg-white text-right">
                <div class="text-caption text-grey-6 br text-left">修改密码</div>
                <q-input clearable class="q-mb-sm" placeholder="新密码" color="red-7" type="password" dense filled v-model="password" />        
            </div>
        </div>
    </div>
    <div class="row q-mt-md q-px-md">
        <div class="col-12 col-sm-6 col-md-4">
            <div class="q-px-md q-py-sm bg-white text-left">
                <div class="text-caption text-grey-6 br">操作</div>
                <q-btn class="full-width q-mb-sm" color="green-5" label="保存信息" @click="saveUser"/>
            </div>
        </div>
    </div>
</div></template>
<script>
export default{
name:'User',
components:{},
props:[],
data(){return {
    password:"",
    user:{}
}},
mounted(){
    this.user = Object.assign({}, {}, this.$store.state.user)
},
methods:{
    saveUser(){
        this.$q.loading.show()
        this.axios.post("/user", {
            Headimg: this.$store.state.user.Headimg,
            Name: this.$store.state.user.Name,
            Email: this.$store.state.user.Email,
            Pass: this.password
        }).then(() => {
            this.$notify("修改成功", "positive")
            this.$q.loading.hide()
        }).catch(() => {
            this.$store.state.user = Object.assign({},{},this.user)
            this.$q.loading.hide()
        })
    }
},
watch:{},
destroy(){},
}
</script>
<style scoped>
.sheight{height:400px}
.myshadow{box-shadow: 5px 8px 10px #F5F5F5, -5px 0px 10px #F5F5F5}
.br:before{content:"_";height:20px;width:3px;background:#EF5350;margin-right:3px}
</style>