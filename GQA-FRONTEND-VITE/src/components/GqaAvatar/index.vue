<template>
    <q-avatar :size="size">
        <img :src="avatarSrc" />
    </q-avatar>
</template>

<script setup>
import { computed, toRefs } from 'vue';
// import { GqaDefaultAvatar } from 'src/settings'
import { useQuasar } from 'quasar';
import useCommon from 'src/composables/useCommon';

const { GqaDefaultAvatar } = useCommon()
const $q = useQuasar();
const props = defineProps({
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
        default: 'md', // 可选:'xs', 'sm', 'md', 'lg', 'xl'， '**px'
    },
})
const { loginUser, src, size } = toRefs(props)

const avatarSrc = computed(() => {
    const cookieAvatar = $q.cookies.get('gqa-avatar')
    if (loginUser.value) {
        // 登录用户的头像
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
            return GqaDefaultAvatar
        }
    } else if (src.value === '') {
        // 非登录用户，没有头像配置
        return GqaDefaultAvatar
    } else if (src.value.substring(0, 4) === 'http') {
        // 非登录上传，头像为链接
        return src.value
    } else if (src.value.substring(0, 11) === 'gqa-upload:') {
        // 非登录用户，头像为上传
        return process.env.API + src.value.substring(11)
    } else {
        return src.value
    }
})
</script>
