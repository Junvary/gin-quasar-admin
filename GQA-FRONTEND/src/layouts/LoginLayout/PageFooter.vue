<template>
    <footer>
        <div class="row justify-around">
            <div class="col">
                <div class="footer-box-cloumn">
                    <div class="footer-box">
                        <q-avatar class="gin-quasar-admin-logo">
                            <img src="gqa128.png" />
                        </q-avatar>
                        <span style="font-size: 25px; font-weight: bold">
                            Gin-Quasar-Admin
                        </span>
                    </div>

                    <p style="text-transform: capitalize;">
                        Be the change you want to see in the world.
                    </p>
                </div>
            </div>
            <div class="col">
                <div class="footer-box-cloumn">
                    <span class="footer-box-title">关于我们</span>
                    <a class="footer-box-item">系统开发科</a>
                    <a class="footer-box-item">系统开发室</a>
                    <a class="footer-box-item">项目组一</a>
                    <a class="footer-box-item">项目组二</a>
                </div>
            </div>
            <div class="col"></div>
            <div class="col">
                <div class="footer-box">
                    <div class="footer-computer"></div>
                </div>
            </div>
        </div>
        <q-separator />
        <div class="row justify-around">
            <div class="col">
                <div class="footer-box">
                    技术支持：{{thisYear}}
                </div>
            </div>
            <div class="col">
                <div class="footer-box">
                    访问仓库
                </div>
            </div>
        </div>

    </footer>
</template>

<script>
export default {
    name: 'PageFooter',
    data() {
        return {
            thisYear: new Date(),
        }
    },
    mounted() {
        this.getThisYear()
    },
    methods: {
        getThisYear() {
            const year = new Date()
            this.thisYear = year.getFullYear()
        },
    },
}
</script>

<style lang="scss" scoped>
footer {
    background: #212b36;
    padding-top: 50px;
    padding-bottom: 30px;
    position: relative;
    z-index: 1;
    overflow: hidden;
    color: #ffffff;
    .footer-box {
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .footer-box-cloumn {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }
    .footer-computer {
        background-image: url('~src/assets/login/computer.svg');
        background-size: cover;
        width: 40vh;
        height: 25vh;
        left: 30px;
        top: 240px;
    }
    .footer-box-title {
        font-size: 20px;
        font-weight: bold;
        margin-bottom: 30px;
    }
    .footer-box-item {
        margin: 8px 0;
        cursor: pointer;
        &:hover {
            color: #3056d3;
        }
    }
}
</style>