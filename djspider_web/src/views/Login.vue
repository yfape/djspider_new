<!--
 * @Author: your name
 * @Date: 2021-07-22 21:44:20
 * @LastEditTime: 2021-08-05 14:52:21
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\views\Login.vue
-->
<template>
<div class="main">
    <div class="main row items-center justify-around">
        <div class="col-12 col-md-6 row items-center" style="height:230px">
            <div class="col row items-center justify-center">
                <q-avatar size="80px" rounded class="shadow-3">
                    <img src="/static/img/logo.jpg">
                </q-avatar>
                <div class="text-white q-ml-lg" style="height:80px;">
                    <div style="font-size:1.8rem;">党建热点推送平台</div>
                    <q-separator style="height:2px;" class="bg-white q-my-xs"/>
                    <div class="text-subtitle1 text-left text-thin">收集&nbsp;归类&nbsp;存储&nbsp;推送</div>
                </div>
            </div>
        </div>
        <div class="col-12 col-md-6 row items-center" style="height:230px">
            <div class="col-12 q-pt-md">
                <div v-if="!show" style="margin:auto">
                    <q-spinner-puff color="red-7" size="2.5rem"/>
                    <div class="text-caption q-mt-md text-grey-7">身份验证中</div>
                </div>
                <div v-else-if="mode==1" id="loginpanel" class="q-pa-md">
                    <q-input class="q-mb-md" outlined color="red-7" type="email" v-model="input.email" label="邮箱"/>
                    <q-input v-if="input.email||input.pass" outlined color="red-7" type="password" v-model="input.pass" label="密码"/>
                    <q-btn class="full-width q-mt-md" v-show="input.email&&input.pass" :loading="loading" flat color="red-7" size="lg" label="登&emsp;录" @click="login"/>
                </div>
                <div v-else style="margin:auto">
                    <Useravater style="margin:auto" :user="$store.state.user" @click="quicklogin"></Useravater>
                    <div class="text-caption q-mt-sm cursor-pointer" @click="smode">点击输入邮箱登录</div>
                </div>
            </div>
        </div>
    </div>
    <div class="desktop-only copyright q-pa-sm text-secondary text-caption">Copyright &copy; 2021 《四川党的建设》杂志社</div>
    <div class="desktop-only ani-1 ani-1-1"></div>
    <div class="desktop-only ani-1 ani-1-2"></div>
    <div class="mobile-only am ani-1-1"></div>
    <div class="mobile-only am ani-1-2"></div>
</div>
</template>
<script>
import Useravater from '@/components/user_avater.vue'
export default{
name: 'Login',
components:{Useravater},
data(){return {
    show:false,loading:false,mode:1,
    input:{email:'',pass:''},
}},
mounted(){
    var ehtml = document.getElementsByTagName('html')[0]
    var ebody = document.getElementsByTagName('body')[0]   
    ehtml.style.cssText = "overflow-y:hidden!important"
    ebody.style.cssText = "overflow-y:hidden!important"
    this.getUserByToken()    
},
methods:{
    getUserByToken(){
        var token = localStorage.getItem('token')
        if(!token){
            this.show = true
            return
        }
        this.axios.post('/loginuser', {token: token}).then( res => {
            this.$store.state.user = res.user
            this.mode = 2
            this.$store.state.token = token
            this.show = true
        }).catch(() => {
            this.show = true
        })
    },
    quicklogin(){
        this.$router.replace("/")
    },
    login(){
        if(!(this.input.email && this.input.pass)){
            return 
        }
        this.loading = true
        this.axios.post('/login', {
            email: this.input.email,
            pass: this.input.pass
        }).then(res => {
            this.loading = false
            this.$notify(res.msg, 'info')
            this.$store.state.token = res.token
            localStorage.setItem('token', res.token)
            this.$store.state.user = res.user
            this.$router.replace("/")
        }).catch(() =>{
            this.loading = false
        });
    },
    smode(){
        this.mode = this.mode == 1 ? 2 : 1
    }
},
}
</script>
<style scoped>
.ani-1{
    position:absolute;top:-80vh;left:-140vw;
    background:#E53935;
    height:199vw;width:200vw;
    border-radius:48%;
}
.ani-1-1{
    animation: ani 8s infinite linear;z-index:-1;
}
.ani-1-2{
    z-index:-2;background:#E87F7D!important;
    animation: ani 6s infinite linear;
}
@keyframes ani{
    from{transform: rotate(0deg)}
    to{transform: rotate(360deg)}
}
#loginpanel{width:400px;margin:auto}
.copyright{position:absolute;z-index:1;bottom:0;right:0;}
.am{position:fixed;height:299vh;width:300vh;top:-246vh;left:-120vh;border-radius:48%;background:#E53935;}
</style>