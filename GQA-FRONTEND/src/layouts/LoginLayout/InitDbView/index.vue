<template>
    <q-card-section horizontal :class="darkThemeLoginHelp"
        style="width: 50%; height: 100%; background: #f3fbff; border-radius: 20px 0 0 20px;">
        <q-card-section class="row justify-center" style="width: 100%;">
            <InitDbViewForm @initDbSuccess="initDbSuccess" />
        </q-card-section>
    </q-card-section>
    <q-separator vertical dark />
    <q-card-section horizontal style="width: 50%;">
        <q-card-section class="row justify-center" style="width: 100%;">
            <InitDbViewHelp />
        </q-card-section>
    </q-card-section>
</template>


<script setup>
import { onMounted } from 'vue'
import { postAction } from 'src/api/manage'
import { useRouter } from 'vue-router'
import InitDbViewForm from './InitDbViewForm.vue'
import InitDbViewHelp from './InitDbViewHelp.vue'
import useTheme from 'src/composables/useTheme';

const { darkThemeLoginHelp } = useTheme()
const router = useRouter()

onMounted(() => {
    checkDb()
})
const checkDb = async () => {
    const res = await postAction('public/check-db')
    if (res.code === 1) {
        if (res.data.need_init === false) {
            router.push({ path: '/' })
        }
    }
}

const emit = defineEmits(['initDbSuccess'])
const initDbSuccess = () => {
    emit('initDbSuccess')
}
</script>