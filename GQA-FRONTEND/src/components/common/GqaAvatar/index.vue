<template>
    <q-avatar :size="size">
        <img :src="avatarSrc" />
    </q-avatar>
</template>

<script setup>
import { computed, toRefs } from 'vue';
import { useQuasar } from 'quasar';
import useConfig from 'src/composables/useConfig';

const { GqaDefaultAvatar } = useConfig()
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
        default: 'md', // 'xs', 'sm', 'md', 'lg', 'xl'， '**px'
    },
})
const { loginUser, src, size } = toRefs(props)

const avatarSrc = computed(() => {
    const cookieAvatar = $q.cookies.get('gqa-avatar')
    if (loginUser.value) {
        // Login user's avatar
        if (cookieAvatar) {
            // avatar in the cookies
            if (cookieAvatar.substring(0, 4) === 'http') {
                // avatar is a link
                return cookieAvatar
            } else if (cookieAvatar.substring(0, 11) === 'gqa-upload:') {
                // avatar is uploaded
                return process.env.API + cookieAvatar.substring(11)
            }
            return cookieAvatar
        } else {
            // default avatar
            return GqaDefaultAvatar
        }
    } else if (src.value === '') {
        // Non login user, no avatar configuration
        return GqaDefaultAvatar
    } else if (src.value.substring(0, 4) === 'http') {
        // avatar is a link
        return src.value
    } else if (src.value.substring(0, 11) === 'gqa-upload:') {
        // avatar is uploaded
        return process.env.API + src.value.substring(11)
    } else {
        return src.value
    }
})
</script>
