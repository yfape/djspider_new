<!--
 * @Author: your name
 * @Date: 2021-07-25 12:06:31
 * @LastEditTime: 2021-08-05 09:59:42
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\views\Articles.vue
-->
<template><div class='q-pa-md'>
    <div>
        <q-input @change="onRequest" dense class="full-width bg-white q-mb-md" square color="red-7" outlined v-model="search" placeholder="搜索">
            <template v-slot:append>
                <q-btn flat icon="o_search" />
            </template>
        </q-input>
    </div>
    
    <q-markup-table flat bordered>
      <thead>
        <tr>
          <th class="text-left" v-for="(item,key,index) in columns" :key="index">{{item.label}}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, key, index) in articles" :key="index">
          <td class="text-left">{{item.Title}}</td>
          <td class="text-left">{{item.Site}}</td>
          <td class="text-left">{{item.Post_time=='1970-01-01'?'':item.Post_time}}</td>
          <td class="text-left">{{item.Create_time}}</td>
          <td class="text-left">
              <q-btn flat dense color="red-8" icon="o_link" @click="GoUrl(item.Url)"/>
          </td>
        </tr>
      </tbody>
    </q-markup-table>
    <div class="row items-center justify-center">
        <q-pagination :disable="loading" class="q-mt-md" v-model="pagination.page" color="red-7" :max="pagination.max" :max-pages="6" boundary-numbers/>
    </div>
</div></template>
<script>

export default{
name:'Articles',
components:{},
props:[],
data(){return {
    articles:[],
    search:'',loading:false,lock:false,
    columns:[
        {name:'Title',field:'Title',align:'left',label:'标题'},
        {name:'Site',field:'Site',align:'left',label:'网站'},
        {name:'Post_time',field:'Post_time',align:'left',label:'发布时间'},
        {name:'Create_time',field:'Create_time',align:'left',label:'爬取时间'},
        {name:'Action',align:'center',label:'操作'}
    ],
    pagination: {
        page: 1,
        max: 1,
        row: 12
    },
}},
mounted(){
    this.onRequest()
},
methods:{
    onRequest(){
        this.loading = true
        this.axios.get(`/articles/${this.pagination.row}/${this.pagination.page}`,{
            params: { search: this.search }
        })
        .then(res => {            
            this.articles = Object.assign([],[],res.data)
            this.pagination.max = res.max
            this.pagination.page = res.page
            this.loading = false
        })
        .catch(()=>{
            this.loading = false
        })
    },
    GoUrl(url){
        window.open(url)
    }
},
watch:{
    "pagination.page"(){
        if(!this.loading){
            this.onRequest()
        }
    }
},
destroy(){},
}
</script>
<style scoped>

</style>