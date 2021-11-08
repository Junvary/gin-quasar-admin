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
    computed: {
        avatarSrc() {
            const cookieAvatar = this.$q.cookies.get('gqa-avatar')
            if (this.loginUser) {
                // 登录游湖的头像
                if (cookieAvatar) {
                    // 如果Cookies里拿到头像，就显示
                    if (cookieAvatar.substring(0, 4) === 'http') {
                        // 头像为链接
                        return cookieAvatar
                    } else if (cookieAvatar.substring(0, 11) === 'gqa-upload:') {
                        // 头像为上传
                        return process.env.API + cookieAvatar.substring(11)
                    }
                    return cookieAvatar
                } else {
                    // 否则显示默认头像
                    return GqaLoginAvatar
                }
            } else if (this.src === '') {
                // 非登录用户，没有头像配置
                return GqaLoginAvatar
            } else if (this.src.substring(0, 4) === 'http') {
                // 非登录上传，头像为链接
                return this.src
            } else if (this.src.substring(0, 11) === 'gqa-upload:') {
                // 非登录用户，头像为上传
                return process.env.API + this.src.substring(11)
            } else {
                // 其他类型，有的时候再说
                console.log('还未完成其他头像类型...')
                return this.src
            }
        },
    },
}
</script>