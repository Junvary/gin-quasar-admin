<template>
    <div class="q-pa-md row items-center justify-center gqa-honour" id="gqa-honour">
        <q-carousel swipeable animated thumbnails infinite :autoplay="autoplay" arrows transition-prev="slide-right"
            transition-next="slide-left" @mouseenter="autoplay = false" @mouseleave="autoplay = 2000" v-model="slide"
            style="width: 60%; height: 650px">
            <q-carousel-slide v-for="item in dataList" :key="item.id" :name="item.id" :img-src="item.src" />
        </q-carousel>
    </div>
</template>

<script>
export default {
    name: 'PageHonour',
    data() {
        return {
            slide: 1,
            autoplay: 2000,
            dataList: [
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
                {
                    id: 4,
                    src: 'https://tva1.sinaimg.cn/large/0072Vf1pgy1foxk7rv2gpj31hc0u04e6.jpg',
                    text: '嘿嘿',
                },
                {
                    id: 5,
                    src: 'https://tva3.sinaimg.cn/large/0072Vf1pgy1foxk7n4ffij31hc0u0wyt.jpg',
                    text: '嘿嘿',
                },
                {
                    id: 6,
                    src: 'https://tva3.sinaimg.cn/large/0072Vf1pgy1foxli9lpyoj31kw0w0qlm.jpg',
                    text: '嘿嘿',
                },
            ],
        }
    },
}
</script>

<style lang="scss" scoped>
.gqa-honour {
    background: rgba(248, 192, 147, 1);
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 100px 0;
}
</style>