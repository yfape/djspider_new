/*
 * @Author: your name
 * @Date: 2021-07-23 10:21:02
 * @LastEditTime: 2021-08-12 16:10:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\utils\axios.js
 */
// 这个时axios的配置
import axios from "axios";
import main from "../main.js";
import store from "../store/index";
// import { config } from 'vue/types/umd';
// axios.defaults.baseURL = 'http://10.100.80.204:8001/v1';
axios.defaults.baseURL = "http://spider_api.***.com/v1";
axios.defaults.headers.post["Content-Type"] =
  "application/json,multipart/form-data";
// 错误信息处理
const errorHandle = (status, msg) => {
  if (status != 200 && status != 201) {
    main.config.globalProperties.$notify(msg);
  }
};
// 添加请求拦截器
axios.interceptors.request.use(
  function (config) {
    config.headers["Authorization"] = store.state.token;
    // 在发送请求之前做些什么
    // console.log(config);
    return config;
  },
  function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

// 添加响应拦截器
axios.interceptors.response.use(
  function (response) {
    // 对响应数据做点什么
    // console.log();
    // console.log(response.data.token);
    // response.headers['Authorization'] = response.data.token;
    return response.status === 200 || response.status === 201
      ? Promise.resolve(response.data)
      : Promise.reject(response);
  },
  function (error) {
    // 对响应错误做点什么
    const { response } = error;
    if (response) {
      errorHandle(response.status, response.data.msg);
    } else {
      main.config.globalProperties.$notify();
    }
    return Promise.reject(response.data);
  }
);
export default axios;
