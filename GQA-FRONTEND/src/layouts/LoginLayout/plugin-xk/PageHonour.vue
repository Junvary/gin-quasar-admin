<template>
    <div class="row items-center justify-center gqa-honour" id="gqa-honour">
        <q-carousel swipeable animated thumbnails infinite :autoplay="autoplay" arrows transition-prev="slide-right"
            transition-next="slide-left" @mouseenter="autoplay = false" @mouseleave="autoplay = 2000" v-model="slide"
            style="width: 60%; height: 650px">
            <q-carousel-slide v-for="(item, index)  in dataList" :key="index" :name="index" :img-src="item.src" />
        </q-carousel>
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'

export default {
    name: 'PageHonour',
    mixins: [tableDataMixin],
    data() {
        return {
            slide: 1,
            autoplay: 2000,
            url: {
                list: 'public/plugin-xk/honour-list',
            },
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
        }
    },
    computed: {
        dataList() {
            if (this.tableData && this.tableData.length) {
                const honourList = []
                this.tableData.forEach((item) => {
                    if (item.attachment) {
                        for (let a of JSON.parse(item.attachment)) {
                            honourList.push({
                                src: process.env.API + a.fileUrl.substring(11),
                            })
                        }
                    }
                })
                return honourList
            } else {
                return [
                    {
                        id: 1,
                        src: 'https://tva3.sinaimg.cn/large/0072Vf1pgy1foxk3jkoywj31hc0u0tr3.jpg',
                        text: '嘿嘿',
                    },
                    {
                        id: 2,
                        src: 'https://tva1.sinaimg.cn/large/0072Vf1pgy1foxkj6err0j31kw0w0hb7.jpg',
                        text: '嘿嘿',
                    },
                    {
                        id: 3,
                        src: 'https://tva2.sinaimg.cn/large/0072Vf1pgy1foxkft70qtj31hc0u0tqy.jpg',
                        text: '嘿嘿',
                    },
                ]
            }
        },
    },
    created() {
        this.getTableData()
    },
}
</script>

<style lang="scss" scoped>
.gqa-honour {
    background: rgba(248, 192, 147, 1);
    width: 100%;
    padding: 100px 0;
}
</style>