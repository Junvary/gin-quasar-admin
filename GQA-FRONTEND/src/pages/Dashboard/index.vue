<template>
    <q-page>
        <q-img :src="randomImg" class="absolute-full">
            <div class="absolute-full text-subtitle2 flex flex-center">
                <span style="font-size: 40px">
                    {{ $t('WelcomeTo') }}
                    Gin-Quasar-Admin！
                </span>
            </div>
        </q-img>
    </q-page>
</template>

<script>
export default {
    name: 'Dashboard',
    data() {
        return {
            randomImg: 'https://api.ixiaowai.cn/api/api.php',
        }
    },
}
</script>
