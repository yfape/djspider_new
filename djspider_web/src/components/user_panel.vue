<template>
<q-dialog v-model="show">
   <div class="q-pa-md bg-white q-mb-md" style="width:400px;overflow:visible">
        <div style="height:40px;overflow:visible"><q-img class="headimg shadow-3 rounded-borders" :src="user.Headimg" :ratio="1" style="width:100px;"/></div>
        <div style="height:0px;overflow:visible" class="text-right q-mb-xs">
            <div class="text-subtitle2 text-bold info">{{user.Name}}</div>
            <div class="text-caption text-grey-7 info">{{user.Email}}</div>
        </div>
        <q-input class="q-mb-sm" color="red-7" filled dense label="姓名" v-model="user.Name"/>
        <q-input class="q-mb-sm" color="red-7" filled dense label="邮箱" v-model="user.Email"/>
        <q-select class="q-mb-md" color="red-7" filled dense
        v-model="user.Identify" 
        :options="options" label="权限" 
        emit-value map-options 
        option-value="id"
        option-label="msg"
        />
        <div class="row q-mb-sm" v-if="user.Userid!=0">
            <div class="col-8 q-pr-sm"><q-btn class="full-width" color="yellow-8" label="重置密码" @click="resetPass"/></div>
            <div class="col-4 q-pl-sm"><q-btn class="full-width" color="red-9" icon="o_delete" label="删除" @click="deleteUser"/></div>
        </div>
        <div class="row">
            <div class="col-4 q-pr-sm"><q-btn class="full-width" color="grey-8" icon="o_undo" label="取消" @click="hidePanel"/></div>
            <div class="col-8 q-pl-sm" v-if="user.Userid!=0"><q-btn class="full-width" color="green-7" label="保存" @click="saveUser"/></div>
            <div class="col-8 q-pl-sm" v-else><q-btn class="full-width" color="yellow-8" label="新增" @click="addUser"/></div>
        </div> 
    </div>
</q-dialog>
</template>
<script>
export default{
name: '',
data(){return {
    show:false,
    user:{},
    init_user:{
        Userid:0,
        Headimg:"https://www.womaimi.com/uploads/allimg/200302/1-20030209142A33.jpg",
        Name:"",
        Email:"",
        Identify:1
    },
    options:[
        { id: 1, msg: '用户' },
        { id: 2, msg: '管理员' }
    ],
    headimgs:[
        'https://img0.baidu.com/it/u=3155998395,3600507640&fm=26&fmt=auto&gp=0.jpg',
        'https://img0.baidu.com/it/u=3401732207,3726302783&fm=26&fmt=auto&gp=0.jpg',
        'https://img2.baidu.com/it/u=1678948314,1083480950&fm=26&fmt=auto&gp=0.jpg',
        'https://img1.baidu.com/it/u=3238897766,3909683226&fm=26&fmt=auto&gp=0.jpg',
        'https://img0.baidu.com/it/u=1114462421,1631208970&fm=26&fmt=auto&gp=0.jpg',
        'https://img1.baidu.com/it/u=1751475856,2552656879&fm=26&fmt=auto&gp=0.jpg',
        'https://img1.baidu.com/it/u=3682391888,1606422123&fm=26&fmt=auto&gp=0.jpg',
        'https://img0.baidu.com/it/u=2152368713,790260685&fm=26&fmt=auto&gp=0.jpg',
        'https://img0.baidu.com/it/u=2265845335,3413513235&fm=26&fmt=auto&gp=0.jpg',
        'https://2b.zol-img.com.cn/product/87_100x75/483/cekizFQWBmfOk.gif'
    ]
}},
mounted(){},
methods:{
    showPanel(user){
        if(user){
            this.user = Object.assign({},{},user)
        }else{
            this.user = Object.assign({},{},this.init_user)
            this.user.Headimg = this.headimgs[this.getRndInteger(0,9)]
        }
        this.show = true
    },
    hidePanel(){
        this.show = false
    },
    getRndInteger(min, max) {
        return Math.floor(Math.random() * (max - min) ) + min;
    },
    resetPass(){
        this.$q.loading.show()
        this.axios.patch("/admin/user/pass",{
            UserId: this.user.UserId
        }).then(()=>{
            this.$notify("重置成功", "info")
            this.$q.loading.hide()
        }).catch(()=>{
            this.$q.loading.hide()
        })
    },
    deleteUser(){
        this.$q.loading.show()
        this.axios.delete("/admin/user/"+this.user.UserId).then(()=>{
            this.$emit("change")
            this.hidePanel()
            this.$notify("删除成功", "info")
            this.$q.loading.hide()
        }).catch(()=>{
            this.$q.loading.hide()
        })
    },
    saveUser(){
        this.$q.loading.show()
        this.axios.patch("/admin/user",{
            UserId: this.user.UserId,
            Headimg: this.user.Headimg,
            Name:this.user.Name,
            Email: this.user.Email,
            Identify: this.user.Identify
        }).then(()=>{
            this.$emit("change")
            this.hidePanel()
            this.$notify("保存成功", "info")
            this.$q.loading.hide()
        }).catch(()=>{
            this.$q.loading.hide()
        })
    },
    addUser(){
        this.$q.loading.show()
        this.axios.post("/admin/user",{
            Headimg: this.user.Headimg,
            Name:this.user.Name,
            Email: this.user.Email,
            Identify: this.user.Identify
        }).then(()=>{
            this.$emit("change")
            this.hidePanel()
            this.$notify("新增成功", "info")
            this.$q.loading.hide()
        }).catch(()=>{
            this.$q.loading.hide()
        })
    }
},
}
</script>
<style scoped>
.headimg{position:relative;top:-70px;border:5px solid #fff;}
.info{position:relative;top:-50px;}
</style>