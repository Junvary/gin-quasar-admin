<template>
    <div class="gqa-honour" id="gqa-honour">
        <waterfall :col="col" :autoResize="autoResize" :moveTransitionDuration="0.4" :fillBox="fillBox"
            :col-width="colWidth" :list="dataList" ref="waterfall" imgKey="src" style="width: 80%">
            <!-- 两种图片绑定模式
        1.指定图片的Key（ imgKey="src"）
        2.在img标签上加class（ class="waterfall-img"） -->
            <!-- img标签如果设置宽高会渲染的更快 -->
            <div class="waterfall-item" :class="{bounceIn:item.state == 'show'}" slot-scope="item" style="margin: 5px">
                <img v-if="item.data.src" style="width: 100%" class="waterfall-img" :src="item.data.src">
                {{item.index}}-{{item.data.text}}
            </div>
        </waterfall>
    </div>
</template>

<script>
import waterfall from 'vue-waterfall-rapid'

export default {
    name: 'PageHonour',
    components: {
        waterfall,
    },
    data() {
        return {
            //是否根据容器尺寸自动计算重绘
            autoResize: true,
            //是否始终填满容器
            fillBox: true,
            //列宽-有指定列数则此属性失效
            colWidth: window.innerWidth / 5,
            //列数
            col: 3,
            dataList: [
                {
                    id: 1,
                    src: 'https://image.watsons.com.cn//upload/13afe9a7.jpg',
                    href: 'www.baidu.com',
                    text: '嘿嘿',
                },
                {
                    id: 2,
                    src: 'https://image.watsons.com.cn//upload/5c3e51e4.jpg',
                    href: 'www.baidu.com',
                    text: '嘿嘿',
                },
                {
                    id: 3,
                    src: 'https://image.watsons.com.cn//upload/13afe9a7.jpg',
                    href: 'www.baidu.com',
                    text: '嘿嘿',
                },
                {
                    id: 4,
                    src: 'https://image.watsons.com.cn//upload/5c3e51e4.jpg',
                    href: 'www.baidu.com',
                    text: '嘿嘿',
                },
                {
                    id: 5,
                    src: 'https://image.watsons.com.cn//upload/13afe9a7.jpg',
                    href: 'www.baidu.com',
                    text: '嘿嘿',
                },
                {
                    id: 6,
                    src: 'https://image.watsons.com.cn//upload/5c3e51e4.jpg',
                    href: 'www.baidu.com',
                    text: '嘿嘿',
                },
            ],
        }
    },
}
</script>

<style lang="scss" scoped>
.gqa-honour {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 100px 0;
}
</style>