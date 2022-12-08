<template>
    <div class="flex text-center text-white fullscreen bg-blue q-pa-md flex-center">
        <div>
            <div style="font-size: 30vh">
                404
            </div>
            <div class="text-h2" style="opacity:.4">
                {{ $t('PageError404Title') }}
            </div>
            <q-btn class="q-mt-xl" color="white" text-color="blue" unelevated to="/" :label="$t('PageError404Home')"
                no-caps />
        </div>
        <AchievementDialog ref="achievementDialog" />
    </div>
</template>

<script setup>
import AchievementDialog from 'src/plugins/Achievement/AchievementDialog.vue';
import { onMounted, ref } from 'vue';
import { postAction } from 'src/api/manage';
import { useUserStore } from 'src/stores/user';

const userStore = useUserStore();
const achievementDialog = ref(null)

onMounted(() => {
    postAction('/plugin-achievement/obtain-find', {
        category_code: 'QiYu-Find-404',
        username: userStore.GetUsername()
    }).then(res => {
        if (res.code === 1 /* && res.data?.get */) {
            achievementDialog.value.show({
                category: "奇遇",
                name: "未知领域！"
            })
        }
    })
})
</script>

