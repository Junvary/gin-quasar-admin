<template>
    <div class="row items-center justify-center gqa-honour" id="gqa-honour">
        <div style="width: 60vw">
            <q-carousel ref="carousel" swipeable animated thumbnails infinite :autoplay="autoplay"
                transition-prev="slide-right" transition-next="slide-left" @mouseenter="autoplay = false"
                @mouseleave="autoplay = 5000" v-model="slide" v-model:fullscreen="fullscreen" height="85vh">
                <q-carousel-slide v-for="(item, index)  in dataList" :key="index" :name="index" :img-src="item.src"
                    style="padding: 0; background-color:rgba(248, 192, 147, 1);">
                    <q-img :src="item.src" style="height: 100%; width: 100%; background-color:rgba(248, 192, 147, 1);"
                        fit="scale-down" />
                </q-carousel-slide>
                <template v-slot:control>
                    <q-carousel-control position="bottom-right" :offset="[18, 18]">
                        <q-btn push round dense color="primary" text-color="white" icon="arrow_left"
                            @click="$refs.carousel.previous()" />
                        <q-btn push round dense color="primary" text-color="white"
                            :icon="fullscreen ? 'fullscreen_exit' : 'fullscreen'" @click="fullscreen = !fullscreen" />
                        <q-btn push round dense color="primary" text-color="white" icon="arrow_right"
                            @click="$refs.carousel.next()" />
                    </q-carousel-control>
                </template>
            </q-carousel>
        </div>

    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaShowName from 'src/components/GqaShowName'

export default {
    name: 'PageHonour',
    mixins: [tableDataMixin],
    components: {
        GqaShowName,
    },
    data() {
        return {
            slide: 1,
            autoplay: 5000,
            url: {
                list: 'public/plugin-xk/honour-list',
            },
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
            fullscreen: false,
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
                                title: item.title,
                                createdByUser: item.createdByUser,
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
                        title: 'Gin-Quasar-Admin',
                        createdByUser: {
                            username: 'Gin-Quasar-Admin',
                        },
                    },
                    {
                        id: 2,
                        src: 'https://tva1.sinaimg.cn/large/0072Vf1pgy1foxkj6err0j31kw0w0hb7.jpg',
                        title: 'Gin-Quasar-Admin',
                        createdByUser: {
                            username: 'Gin-Quasar-Admin',
                        },
                    },
                    {
                        id: 3,
                        src: 'https://tva2.sinaimg.cn/large/0072Vf1pgy1foxkft70qtj31hc0u0tqy.jpg',
                        title: 'Gin-Quasar-Admin',
                        createdByUser: {
                            username: 'Gin-Quasar-Admin',
                        },
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
.custom-caption {
    text-align: center;
    padding: 5px;
    color: white;
    background-color: rgba(0, 0, 0, 0.3);
}
</style>