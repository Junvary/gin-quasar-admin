<template>
    <div>
        <q-btn round glossy push color="primary" :icon="`${!right ? 'settings' : 'close'}`" @click="right = !right" />
        <q-dialog v-model="right" position="right">
            <q-card style="height: 100%; width: 400px">
                <q-card-section>
                    <div class="text-h6">系统设置</div>
                </q-card-section>
                <q-card-section>
                    <gqa-theme />
                </q-card-section>

            </q-card>
        </q-dialog>
    </div>
</template>

<script>
import GqaTheme from 'src/components/GqaTheme'

export default {
    name: 'Setting',
    components: {
        GqaTheme,
    },
    data() {
        return {
            right: false,
        }
    },
}
</script>