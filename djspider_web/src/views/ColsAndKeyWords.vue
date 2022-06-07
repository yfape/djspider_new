<!--
 * @Author: your name
 * @Date: 2021-07-25 19:33:12
 * @LastEditTime: 2021-08-04 16:16:38
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\views\ColsAndKeyWords.vue
-->
<template><div class='q-pa-md'>
    <div class="bg-white q-pa-md q-mb-md">
        <div class="text-caption text-grey-7 text-left">
            文章标题与下列关键字匹配则邮件推送给您
        </div>
        <div class="row q-gutter-sm q-mt-xs">
            <div class="q-mr-sm">
                <q-btn dense padding="2px 10px" outline label="新增关键字" icon="o_add">
                    <q-popup-edit v-model="de">
                        <q-input outlined color="red-7" label="请输入关键字" dense v-model="editkeyword">
                            <template v-slot:append>
                                <q-btn flat icon="o_add" dense @click="addKeyword" v-close-popup/>
                            </template>
                        </q-input>
                    </q-popup-edit>
                </q-btn>
            </div>
            <q-chip v-for="(item,key,index) in keywords" :key="index" square removable :color="getColor()" @remove="removeKeyword(item)" text-color="white">
                {{item}}
            </q-chip>
        </div>
    </div>
    <div class="bg-white q-pa-md">
        <div class="text-caption text-grey-7 text-left">请勾选您需要邮件推送的网站和栏目</div>
        <q-tree class="col-12 col-sm-6"
        color="red-7"
        accordion
        :nodes="simple"
        node-key="Spi_id"
        label-key="Name"
        children-key="Children"
        tick-strategy="leaf"
        v-model:ticked="selected"
        @update:ticked="selectCol"
        />
    </div>
</div></template>
<script>

export default{
name:'ColsAndKeyWords',
components:{},
props:[],
data(){return {
    colors:['red-6','red-7','blue-7','blue-5','teal-6','cyan-6','cyan-7','green-6','yellow-8','pink-7','brown-6'],
    selected:[1,2,3,4,5,6],
    old_selected:[],
    simple: [],
    keywords:[],
    editkeyword:"",de:""
}},
mounted(){
    this.getCols()
},
methods:{
    getCols(){
        this.axios.get('/cols').then(res => {
            this.simple = Object.assign([],[],res.data)
            this.getUserCols()
        })
    },
    getUserCols(){
        this.axios.get('/user/colsandkeywords').then(res => {
            if(res.cols.length > 0){
                var colsarr = res.cols.split(",")
                colsarr = colsarr.map(item => {
                    return +item;  
                });  
                this.selected = Object.assign([],[],colsarr)
            }else{
                this.selected = Object.assign([],[],[])
            }
            if(res.keywords.length > 0){
                var keysarr = res.keywords.split(",")
                this.keywords = Object.assign([],[],keysarr)
            }else{
                this.keywords = Object.assign([],[],[])
            }
            this.old_selected = Object.assign([],[],this.selected)
        })
    },
    removeKeyword(word){
        this.$q.loading.show()
        this.axios.delete("/user/keyword/"+word).then(res => {
            this.keywords = Object.assign([],[],res.keywords)
            this.$q.loading.hide()
        }).catch(()=>{
            this.$q.loading.hide()
        })
    },
    getColor(){
        var rand = Math.floor(Math.random()*this.colors.length)
        return this.colors[rand]
    },
    addKeyword(){
        if(!this.editkeyword){
            return
        }
        this.$q.loading.show()
        this.axios.post("/user/keyword", {
            "keyword": this.editkeyword
        }).then(res => {
            this.keywords = Object.assign([],[],res.keywords)
            this.editkeyword = ""
            this.$q.loading.hide()
        }).catch(()=>{
            this.$q.loading.hide()
        })
    },
    selectCol(s){
        this.$q.loading.show()
        var rs = Object.assign([],[],s)
        for (let i = 0; i < s.length; i++) {
            rs[i] = rs[i] + ""
        }
        this.axios.post("/user/cols", {
            "cols": rs
        }).then(() => {
            this.old_selected = Object.assign([],[],this.selected) 
            this.$q.loading.hide()
        }).catch(()=>{
            this.selected = Object.assign([],[],this.old_selected) 
            this.$q.loading.hide()
        })
    }
},
watch:{},
destroy(){},
}
</script>
<style scoped>

</style>