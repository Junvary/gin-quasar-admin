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
    </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
    name: 'Error404',
})
</script>
