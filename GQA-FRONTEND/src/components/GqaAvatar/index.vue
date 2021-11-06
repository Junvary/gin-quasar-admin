<template>
    <q-avatar :size="size">
        <img :src="avatarSrc" />
    </q-avatar>
</template>

<script>
import { GqaLoginAvatar } from 'src/settings'
export default {
    name: 'GqaAvatar',
    props: {
        // 当前登录账号的头像，从Cookie里面取
        loginUser: {
            type: Boolean,
            required: false,
            default: false,
        },
        src: {
            type: String,
            required: false,
            default: '',
        },
        size: {
            type: String,
            required: false,
            default: 'md', // 可选：'xs', 'sm', 'md', 'lg', 'xl'， '**px'
        },
    },
    data() {
        return {
            loading: false,
            avatarSrc: '',
        }
    },
    created() {
        this.checkAvatar()
    },
    methods: {
        checkAvatar() {
            this.loading = true
            if (this.loginUser) {
                // 登录游湖的头像
                if (this.$q.cookies.get('gqa-avatar')) {
                    // 如果Cookies里拿到头像，就显示
                    this.avatarSrc = this.$q.cookies.get('gqa-avatar')
                } else {
                    // 否则显示默认头像
                    this.avatarSrc = GqaLoginAvatar
                }
            } else if (this.src === '') {
                // 非登录用户，没有头像配置
                this.avatarSrc = GqaLoginAvatar
            } else if (this.src.substring(0, 4) === 'http') {
                // 非登录游湖，头像为连接
                this.avatarSrc = this.src
            } else {
                // 其他类型，有的时候再说
                console.log('还未完成其他头像类型...')
                this.avatarSrc = this.src
            }
            this.loading = false
        },
    },
}
</script>