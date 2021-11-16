<template>
    <header>
        <a href="javascript:;" class="logo">
            <q-avatar class="gin-quasar-admin-logo">
                <img src="gqa128.png" />
            </q-avatar>
            <span style="margin-left: 10px">
                {{ gqaFrontend.gqaMainTitle }}
            </span>
        </a>
        <ul>
            <li>
                <a class="gqa-menu-scroll" href="javascript:;" @click="handleScroll('gqa-banner')">首页</a>
            </li>
            <li>
                <a class="gqa-menu-scroll" href="javascript:;" @click="handleScroll('gqa-news')">最新要闻</a>
            </li>
            <li>
                <a class="gqa-menu-scroll" href="javascript:;" @click="handleScroll('gqa-project')">项目进度</a>
            </li>
            <li>
                <a class="gqa-menu-scroll" href="javascript:;" @click="handleScroll('gqa-honour')">荣誉认证</a>
            </li>
            <li>
                <a class="gqa-menu-scroll" href="javascript:;" @click="handleScroll('gqa-weapon')">武器库</a>
            </li>
            <li>
                <a class="gqa-menu-scroll" href="javascript:;" @click="handleScroll('gqa-document')">相关文档</a>
            </li>
            <li>
                <a class="gqa-menu-scroll" href="javascript:;" @click="handleScroll('gqa-download')">常用下载</a>
            </li>
            <li>
                <a class="gqa-menu-scroll" href="javascript:;">
                    版本信息
                    <GqaVersion />
                </a>
            </li>
        </ul>
    </header>
</template>

<script>
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'
import GqaVersion from 'src/components/GqaVersion'

export default {
    name: 'PageHeader',
    mixins: [gqaFrontendMixin],
    components: {
        GqaVersion,
    },
    methods: {
        flow() {
            // 监视下滑，改变导航栏
            window.addEventListener('scroll', () => {
                var header = document.querySelector('header')
                header.classList.toggle('sticky', window.scrollY > 0)
            })
        },
        handleScroll(ele) {
            document.getElementById(ele).scrollIntoView({
                behavior: 'smooth',
            })
        },
    },
    mounted() {
        this.flow()
    },
}
</script>

<style lang="scss" scoped>
header {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    transition: 0.6s;
    padding: 40px 100px;
    z-index: 100;
    box-sizing: border-box;
}
header.sticky {
    padding: 10px 100px;
    background: rgba(255, 255, 255, 0.8);
    box-shadow: 0px 1px 5px #888888;
}
header .logo {
    position: relative;
    font-weight: 700;
    color: #ffffff;
    text-decoration: none;
    font-size: 25px;
    text-transform: uppercase;
    letter-spacing: 2px;
    transition: 0.6s;
    display: flex;
    align-items: center;
    user-select: none;
}
header ul {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
}
header ul li {
    position: relative;
    list-style: none;
    &:hover {
        a {
            text-shadow: #fff 1px 0 0, #fff 0 1px 0, #fff -1px 0 0, #fff 0 -1px 0;
            color: #3056d3;
        }
    }
}
header ul li a {
    position: relative;
    margin: 0 15px;
    text-decoration: none;
    color: #ffffff;
    letter-spacing: 2px;
    font-weight: bold;
    font-size: 20px;
    transition: 0.6s;
    user-select: none;
}
header.sticky .logo,
header.sticky ul li a {
    color: black;
    &:hover {
        text-shadow: #fff 1px 0 0, #fff 0 1px 0, #fff -1px 0 0, #fff 0 -1px 0;
        color: #3056d3;
    }
}
.gqa-menu-scroll {
    cursor: pointer;
}
</style>