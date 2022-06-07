<!--
 * @Author: your name
 * @Date: 2021-07-23 15:46:48
 * @LastEditTime: 2021-08-19 17:36:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\views\Center.vue
-->
<template>
<q-layout view="lHh lpR fFf" class="bg-grey-3">

<q-header bordered class="bg-white text-white" height-hint="98">
  <q-toolbar>
    <q-btn class="desktop-only" dense flat round icon="o_menu" color="red-7" @click="mini = !mini" />
    <q-btn class="mobile-only" dense flat round icon="o_menu" color="red-7" @click="left = !left" />
    <q-space />
    <q-btn-dropdown color="red-3" class="text-white q-ma-none q-pa-none" flat>
        <template v-slot:label>
            <q-item clickable v-ripple class="bg-red-5">
                <q-item-section top avatar>
                    <q-avatar rounded>
                        <img :src="$store.state.user.Headimg">
                    </q-avatar>
                </q-item-section>
                <q-item-section class="text-left text-white">
                    <q-item-label>{{$store.state.user.Name}}</q-item-label>
                    <q-item-label caption class="text-white">{{$store.state.user.Email}}</q-item-label>
                </q-item-section>
            </q-item>
        </template>
        <q-list>
            <q-item clickable v-close-popup to="/user" active-class="text-red-7">
            <q-item-section>
                <q-item-label>用户信息</q-item-label>
            </q-item-section>
            </q-item>

            <q-item clickable v-close-popup to="/colsandkeywords" active-class="text-red-7">
            <q-item-section>
                <q-item-label>关注栏目及关键字</q-item-label>
            </q-item-section>
            </q-item>

            <q-item clickable v-close-popup to="/login" active-class="text-red-7">
            <q-item-section>
                <q-item-label>退出</q-item-label>
            </q-item-section>
            </q-item>
        </q-list>
    </q-btn-dropdown>
    
  </q-toolbar>
</q-header>

<q-drawer v-model="left" side="left" class="bg-red-7" :width="250" dark :mini="mini">
    <div class="q-py-lg row items-center justify-center no-wrap">
        <q-avatar rounded>
            <img src="/static/img/logo.jpg">
        </q-avatar>
        <div v-if="!mini" class="text-h6 text-white q-ml-sm">党建热点推送平台</div>
    </div>
    <q-list padding class="text-white" dark>
        <q-item to="/" exact append class="ai1" active-class="text-red-7 bg-grey-3 ai shadow-3">
            <q-item-section avatar>
                <q-icon name="o_home" />
            </q-item-section>
            <q-item-section class="text-subtitle1">
                数据中心
            </q-item-section>
        </q-item>
        <q-item to="/articles" exact append class="ai1" active-class="text-red-7 bg-grey-3 ai shadow-3">
            <q-item-section avatar>
                <q-icon name="o_article" />
            </q-item-section>
            <q-item-section class="text-subtitle1">
                文章库
            </q-item-section>
        </q-item>
        <q-item to="/colsandkeywords" exact append class="ai1" active-class="text-red-7 bg-grey-3 ai shadow-3">
            <q-item-section avatar>
                <q-icon name="o_dashboard_customize" />
            </q-item-section>
            <q-item-section class="text-subtitle1">
                栏目及关键字
            </q-item-section>
        </q-item>
        <q-item to="/user" exact append class="ai1" active-class="text-red-7 bg-grey-3 ai shadow-3">
            <q-item-section avatar>
                <q-icon name="o_face" />
            </q-item-section>
            <q-item-section class="text-subtitle1">
                个人信息
            </q-item-section>
        </q-item>
        <q-separator v-if="$store.state.user.Identify=='管理员'"/>
        <q-item v-if="$store.state.user.Identify=='管理员'" to="/manageuser" exact append class="ai1" active-class="text-red-7 bg-grey-3 ai shadow-3">
            <q-item-section avatar>
                <q-icon name="o_group" />
            </q-item-section>
            <q-item-section class="text-subtitle1">
                用户管理
            </q-item-section>
        </q-item>
        <q-item v-if="$store.state.user.Identify=='管理员'" to="/managecol" exact append class="ai1" active-class="text-red-7 bg-grey-3 ai shadow-3">
            <q-item-section avatar>
                <q-icon name="o_dashboard" />
            </q-item-section>
            <q-item-section class="text-subtitle1">
                栏目管理
            </q-item-section>
        </q-item>
    </q-list>
    
</q-drawer>

<q-page-container>
    <transition name="fade">
        <router-view/>
    </transition>
</q-page-container>

</q-layout>
</template>
<script>

export default{
name: 'Index',
components:{},
data(){return {
    show:false,left:true,mini:false,link:''
}},
mounted(){
    var ehtml = document.getElementsByTagName('html')[0]
    var ebody = document.getElementsByTagName('body')[0]
    ehtml.style.cssText = "overflow-y:auto!important"
    ebody.style.cssText = "overflow-y:auto!important"
    this.show = true
},
methods:{},
}
</script>
<style>
html,body{overflow-y:auto!important;}
.ai1{position:relative;left:0%;width:100%;border-radius:0px 0px 0px 0px;transition:border-radius 0.3s,width 0.3s,left 0.3s}
.ai{border-radius:24px 0px 0px 24px;width:80%;left:20%;}
</style>