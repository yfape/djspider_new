/*
 * @Author: your name
 * @Date: 2021-07-22 15:40:36
 * @LastEditTime: 2021-08-19 17:36:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\router\index.js
 */
import { createRouter, createWebHashHistory } from 'vue-router'
import app from '../main'
import store from '../store/index'
import Login from '../views/Login'
import Index from '../views/Index'
import Center from '../views/Center'
import Articles from '../views/Articles'
import ColsAndKeyWords from '../views/ColsAndKeyWords'
import User from '../views/User'
import ManageUser from '../views/ManageUser'
import ManageCol from '../views/ManageCol'

const routes = [
  { path: '/login', name: 'Login', component: Login },
  { path: '/', name: 'Index', component: Index, children:[
    { path: '/', name: 'Center', component: Center},
    { path: '/articles', name: 'Articles', component: Articles },
    { path: '/colsandkeywords', name: 'ColsAndKeyWords', component: ColsAndKeyWords },
    { path: '/user', name: 'User', component: User },
    { path: '/manageUser', name: 'ManageUser', component: ManageUser },
    { path: '/manageCol', name: 'ManageCol', component: ManageCol },
  ]}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  if(to.name != 'Login' && !store.state.token){
    var token = localStorage.getItem('token')
    if (!token){
      next('/login')
    }else{
      app.axios.post('/loginuser', {token: token}).then( res => {
        store.state.user = res.user
        store.state.token = token
        next()
      }).catch(() => {
        next('/login')
      })
    }
  }else{
    next()
  }
})

export default router
