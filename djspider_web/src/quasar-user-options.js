/*
 * @Author: your name
 * @Date: 2021-07-22 15:38:25
 * @LastEditTime: 2021-07-30 16:38:39
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\quasar-user-options.js
 */

import './styles/quasar.sass'
import iconSet from 'quasar/icon-set/material-icons-outlined.js'
import lang from 'quasar/lang/zh-CN.js'
import '@quasar/extras/material-icons-outlined/material-icons-outlined.css'
import {Notify,Loading} from 'quasar'

// To be used on app.use(Quasar, { ... })
export default {
  config: {},
  plugins: {
    Notify,
    Loading
  },
  lang: lang,
  iconSet: iconSet
}