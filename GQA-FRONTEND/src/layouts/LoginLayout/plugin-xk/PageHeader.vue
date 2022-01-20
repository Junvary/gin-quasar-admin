<template>
    <header>
        <div class="logo">
            <GqaAvatar size="xl" :src="gqaFrontend.gqaWebLogo" />
            <span style="margin-left: 10px">
                {{ gqaFrontend.gqaMainTitle }}
            </span>
        </div>
        <ul>
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-banner')" label="首页" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-news')" label="最新要闻" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-project')" label="项目进度" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-weapon')" label="武器库" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-honour')" label="荣誉认证" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-resource')" label="资源查阅" />
            <q-btn flat rounded push size="lg" text-color="white" label="版本信息">
                <GqaVersion />
            </q-btn>
        </ul>
    </header>
</template>

<script>
import { gqaFrontendMixin } from 'src/mixins/gqaFrontendMixin'
import GqaVersion from 'src/components/GqaVersion'
import GqaAvatar from 'src/components/GqaAvatar'

export default {
    name: 'PageHeader',
    mixins: [gqaFrontendMixin],
    components: {
        GqaVersion,
        GqaAvatar,
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
    beforeUnmount() {
        window.removeEventListener('scroll', () => {
            var header = document.querySelector('header')
            header.classList.toggle('sticky', window.scrollY > 0)
        })
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
    cursor: pointer;
}
header ul {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
}

header.sticky .logo,
header.sticky {
    color: black;
    .q-btn--rounded {
        color: black !important;
    }
}
</style>