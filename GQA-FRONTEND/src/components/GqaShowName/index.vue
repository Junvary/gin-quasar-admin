<template>
    <span>
        {{ trueName }}
    </span>
</template>

<script setup>
import { useQuasar } from 'quasar';
// import { GqaDefaultUsername } from 'src/settings'
import { computed, toRefs } from 'vue';
import useCommon from 'src/composables/useCommon';

const { GqaDefaultUsername } = useCommon()
const $q = useQuasar()
const props = defineProps({
    customNameString: {
        type: String,
        required: false,
        default: '',
    },
    customNameObject: {
        type: Object,
        required: false,
        default: () => {
            return {}
        },
    },
    showMyName: {
        type: Boolean,
        required: false,
        default: false,
    },
})
const { customNameString, customNameObject, showMyName } = toRefs(props)

const trueName = computed(() => {
    if (customNameString.value !== '') {
        // 自定义名字
        return customNameString.value
    } else if (JSON.stringify(customNameObject.value) !== '{}') {
        // 其他用户的名字
        if (customNameObject.value.nickname) {
            return customNameObject.value.nickname
        } else if (customNameObject.value.real_name) {
            return customNameObject.value.real_name
        } else {
            return customNameObject.value.username
        }
    } else if (showMyName.value) {
        const nickname = $q.cookies.get('gqa-nickname')
        const realName = $q.cookies.get('gqa-realName')
        if (nickname) {
            return nickname
        } else {
            return realName
        }
    } else {
        return GqaDefaultUsername
    }
})
</script>
