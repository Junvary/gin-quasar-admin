<template>
    <div class="fullscreen bg-blue text-white text-center q-pa-md flex flex-center">
        <div>
            <div style="font-size: 30vh">
                404
            </div>

            <div class="text-h2" style="opacity:.4">
                恭喜你，这个页面似乎还没人发现！
            </div>

            <q-btn class="q-mt-xl" color="white" text-color="blue" unelevated to="/" label="带我回家" no-caps />
        </div>
    </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
    name: 'Error404',
})
</script>
