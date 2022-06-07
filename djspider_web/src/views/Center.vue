<!--
 * @Author: your name
 * @Date: 2021-07-24 08:14:24
 * @LastEditTime: 2021-08-09 17:12:29
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider_web\src\views\Center.vue
-->
<template>
<div v-show="show">
    <div class="q-pa-lg row q-gutter-lg desktop-only">
        <block1 title="入库文章数" :num="content.count">
            <template v-slot:icon><q-spinner-facebook color="red-6" size="45px"/></template>
        </block1>
        <block1 title="今日有效抓取" :num="content.newcount">
            <template v-slot:icon><q-spinner-box color="green-7" size="55px"/></template>
        </block1>
        <block1 title="抓取站点数" :num="content.webcount">
            <template v-slot:icon><q-icon name="o_drag_indicator" color="orange-8" size="45px"/></template>
        </block1>
        <block1 title="用户总数" :num="content.usercount">
            <template v-slot:icon><q-icon name="o_group" color="light-green-5" size="45px"/></template>
        </block1>
    </div>

    <div v-if="content.sites.length>=0" id="grap1" class="col-12 bg-white q-py-sm q-mx-lg row items-center justify-center desktop-only"></div>

    <div class="row items-start justify-center q-pa-lg q-col-gutter-lg">
        <div class="col-md-6 col-12">
            <q-item v-for="(item, key, index) in content.articles" :key="index" clickable v-ripple class="bg-white q-mb-md" @click="GoUrl(item.Url)">
                <q-item-section>
                <q-item-label class="text-left">{{item.Title}}</q-item-label>
                <q-item-label caption class="text-right">{{item.Site}} {{item.Time}}</q-item-label>
                </q-item-section>
            </q-item>
        </div>
        <div class="col-md-6 col-12">
            <q-list class="bg-white">
                <q-item>
                    <q-item-section class="text-left row">
                        <div class="row items-center">
                            <q-icon name="o_bug_report" color="red-10" size="18px" class="q-mr-sm"/>
                            <div>定时爬虫进程</div>
                            <div v-if="content.spider" class="text-caption text-grey q-ml-md">{{actiontime.min}}分{{actiontime.sec}}秒</div>
                        </div>
                    </q-item-section>
                    <q-item-section avatar>
                        <div class="row items-center">
                            <div v-if="content.spider" class="text-green-7">已运行</div>
                            <div v-else class="text-grey-6">已关闭</div>
                            <q-separator v-if="$store.state.user.Identify=='管理员'" vertical class="q-mx-sm"/>
                            <div v-if="$store.state.user.Identify=='管理员'">
                                <q-btn v-if="!loading1" flat dense color="red-7" icon="o_power_settings_new" :loading="loading1" @click="switchSpider"/>
                                <q-spinner-ios v-else color="red-7" size="32px"/>
                            </div>
                        </div>
                    </q-item-section>
                </q-item>
                <q-separator />
                <q-item>
                    <q-item-section class="text-left row">
                        <div class="row items-center"><q-icon name="o_public" color="light-green-7" size="18px" class="q-mr-sm"/>HTTP服务进程</div>
                    </q-item-section>
                    <q-item-section avatar>
                        <div class="row items-center">
                            <div v-if="content.http" class="text-green-7">已运行</div>
                            <div v-else class="text-grey-6">已关闭</div>
                            <q-separator v-if="$store.state.user.Identify=='管理员'" vertical class="q-mx-sm"/>
                            <div v-if="$store.state.user.Identify=='管理员'">
                                <q-btn v-if="!loading2" flat dense color="red-7" icon="o_power_settings_new" :loading="loading2" @click="switchHttp"/>
                                <q-spinner-ios v-else color="red-7" size="32px"/>
                            </div>
                        </div>
                    </q-item-section>
                </q-item>
                <q-separator v-if="$store.state.user.Identify=='管理员'"/>
                <q-item v-if="$store.state.user.Identify=='管理员'">
                    <q-item-section class="text-left row">
                        <div class="row items-center"><q-icon name="o_flash_on" color="blue-7" size="18px" class="q-mr-sm"/>单次爬取推送</div>
                    </q-item-section>
                    <q-item-section avatar>
                        <div class="row items-center">
                            <q-btn :loading="loading3" flat dense color="green-7" icon="o_play_arrow" @click="runonce"/>
                        </div>
                    </q-item-section>
                </q-item>
                <q-separator/>
                <q-item>
                    <q-item-section class="text-left row">
                        <div class="row items-center"><q-icon name="o_description" color="teal-7" size="18px" class="q-mr-sm"/>查看爬取日志</div>
                    </q-item-section>
                    <q-item-section avatar>
                        <div class="row items-center">
                            <q-btn :loading="loading3" flat dense color="green-7" icon="o_visibility" @click="showLogPanel"/>
                        </div>
                    </q-item-section>
                </q-item>
            </q-list>
        </div>
    </div>
</div>
<div v-show="!show" class="text-center q-py-xl text-grey-7">
    等待数据加载
</div>
<Panellog ref="Panellog"></Panellog>
</template>
<script>
import block1 from '../components/block/block_1'
import Panellog from '../components/panel_log'
import * as echarts from 'echarts';
export default{
name:'Center',
components:{block1,Panellog},
props:[],
data(){return {
    show:false,content:{sites:[],articles:[],count:0,webcount:0,spidercount:0,usercount:0,newcount:0,spider:0,http:0,lastspider:0},
    loading1:false,loading2:false,loading3:false,actiontime:{min:0, sec:0},dataAxis:[],dataAxisdata:[]
}},
mounted(){
    this.ld()
},
methods:{
    showLogPanel(){
        this.$refs.Panellog.showPanel()
    },
    runonce(){
        this.loading3 = true
        this.axios.post('/runonce').then(res => {
            this.$notify(res.msg, "positive")
            this.loading3 = false
        }).catch(()=>{
            this.loading3 = false
        })
    },
    GoUrl(url){
        window.open(url)
    },
    getActiontime(){
        this.cAT()
        setInterval(() => {
            this.cAT()
        }, 999);
    },
    cAT(){
        var ntc = this.content.lastspider - Date.parse(new Date()) / 1000
        console.log(ntc)
        if (ntc < 0){
            setTimeout(() => {
                this.ld()
            }, 3000);
            return
        }
        this.actiontime.min = parseInt(ntc / 60)
        this.actiontime.sec = ntc % 60
    },
    switchHttp(){
        this.loading2 = true
        this.axios.post('/switchhttp').then(res =>{
            this.$notify(res.msg, 'positive')
            this.ld()
            this.loading2 = false
        }).catch(()=>{
            this.loading2 = false
        })
    },
    switchSpider(){
        this.loading1 = true
        this.axios.post('/switchspider').then(res =>{
            if(this.content.spider == 0){
                setTimeout(() => {
                    this.ld()
                    this.$notify(res.msg, 'positive')
                    this.loading1 = false
                }, 1000);
            }else{
                this.$notify(res.msg, 'positive')
                this.ld()
                this.loading1 = false
            }   
        }).catch(()=>{
            this.loading1 = false
        })
    },
    ld(){
        this.axios.get('/center').then(res => {
            this.content = Object.assign({},{}, res)
            this.setGrap1Data()
            this.show = true
            if(this.content.spider){
                this.getActiontime()
            } 
            setTimeout(() => {
                this.grap1()
            }, 300);
        })
    },
    setGrap1Data(){
        let array = this.content.sites
        this.dataAxis = []
        this.dataAxisdata = []
        for (let i = 0; i < array.length; i++) {
            this.dataAxis.push(array[i].Name) 
            this.dataAxisdata.push(array[i].Count)
        }
    },
    grap1(){
        var chartDom = document.getElementById('grap1');
        var myChart = echarts.init(chartDom);
        var option;
        var dataAxis = this.dataAxis;
        var data = this.dataAxisdata;
        var yMax = 500;
        var dataShadow = [];

        for (var i = 0; i < data.length; i++) {
            dataShadow.push(yMax);
        }
        myChart.resize()
        option = {
            title: { text: '各网站爬取情况', left: '5%', textAlign: 'center'},
            tooltip: { trigger: 'axis', axisPointer: {type: 'shadow'}},
            grid: {left: '3%', right: '4%', bottom: '3%', containLabel: true },
            xAxis: { 
                data: dataAxis, axisTick: { show: false },
                axisLine: { show: true }, z: 10
            },
            yAxis: {
                axisLine: { show: true },
                axisTick: { show: true },
                axisLabel: { textStyle: { color: '#999' } }
            },
            toolbox: { feature: { saveAsImage: {} } },
            dataZoom: [ { type: 'inside' } ],
            series: [
                {
                    type: 'bar',
                    showBackground: false,
                    itemStyle: {
                        color: new echarts.graphic.LinearGradient(
                            0, 0, 0, 1,
                            [
                                {offset: 0, color: '#FFEBEE'},
                                {offset: 1, color: '#E53935'}
                            ]
                        )
                    },
                    emphasis: {
                        itemStyle: {
                            color: new echarts.graphic.LinearGradient(
                                0, 0, 0, 1,
                                [
                                    {offset: 0, color: '#E8F5E9'},
                                    {offset: 1, color: '#43CACF'}
                                ]
                            )
                        }
                    },
                    data: data
                }
            ]
        };

        var zoomSize = 6;
        myChart.on('click', function (params) {
            myChart.dispatchAction({
                type: 'dataZoom',
                startValue: dataAxis[Math.max(params.dataIndex - zoomSize / 2, 0)],
                endValue: dataAxis[Math.min(params.dataIndex + zoomSize / 2, data.length - 1)]
            });
        });

        option && myChart.setOption(option);
        
    }
},
watch:{},
destroy(){},
}
</script>
<style scoped>
#grap1{height:300px}
</style>