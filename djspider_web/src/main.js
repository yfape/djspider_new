/*
 * @Author: your name
 * @Date: 2021-07-22 15:37:24
 * @LastEditTime: 2021-07-25 20:49:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\main.js
 */
import { createApp } from 'vue'
import App from './App.vue'
import { Quasar } from 'quasar'
import quasarUserOptions from './quasar-user-options'
import router from './router'
import store from './store'
import axios from './utils/axios'
import VueAxios from 'vue-axios'

const app = createApp(App).use(store).use(router).use(VueAxios,axios).use(Quasar, quasarUserOptions);
app.config.globalProperties.$notify = function(msg='服务器繁忙，请稍后再试', type='negative',position="top"){
    this.$q.notify({
        type: type,
        position: position,
        message: msg,
        timeout:1500
    })
}
app.mount('#app')

export default app