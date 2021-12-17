<template>
    <q-menu v-model="showVersion" anchor="bottom right" self="top right">
        <q-list dense bordered separator>
            <q-item clickable v-close-popup v-ripple>
                <q-item-section>
                    <q-chip icon="img:icon/vue.svg">
                        Vue: {{ $vueVersion }}
                    </q-chip>
                </q-item-section>
            </q-item>
            <q-item clickable v-close-popup v-ripple>
                <q-item-section>
                    <q-chip icon="img:icon/quasar.svg">
                        Quasar: {{ $quasarVersion }}
                    </q-chip>
                </q-item-section>
            </q-item>
            <q-item clickable v-close-popup v-ripple>
                <q-item-section>
                    <q-chip icon="img:icon/golang.svg">
                        Go: {{ goVersion }}
                    </q-chip>
                </q-item-section>
            </q-item>
            <q-item clickable v-close-popup v-ripple>
                <q-item-section>
                    <q-chip icon="img:icon/gin.svg">
                        Gin: {{ ginVersion }}
                    </q-chip>
                </q-item-section>
            </q-item>
        </q-list>
    </q-menu>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
    name: 'GqaVersion',
    computed: {
        ...mapGetters({
            goVersion: 'storage/gqaGoVersion',
            ginVersion: 'storage/gqaGinVersion',
        }),
    },
    data() {
        return {
            showVersion: false,
        }
    },
}
</script>