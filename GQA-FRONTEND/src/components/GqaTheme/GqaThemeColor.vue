<template>
    <div>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span :style="{ color: primary }">
                    {{ $t('Theme') + $t('Color') }}
                </span>
                <span>
                    primary: {{ primary }}
                </span>
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round v-for="(item, index) in colorList" :key="index" :size="primary === item ? 'md' : 'xs'"
                        :style="{ border: '1px solid grey', background: item }" class="cursor"
                        @click="setBrand('primary', item)" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span :style="{ color: secondary }">
                    {{ $t('Theme') + $t('Color') }}
                </span>
                <span>
                    secondary: {{ secondary }}
                </span>
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round v-for="(item, index) in colorList" :key="index" :size="secondary === item ? 'md' : 'xs'"
                        :style="{ border: '1px solid grey', background: item }" class="cursor"
                        @click="setBrand('secondary', item)" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span :style="{ color: accent }">
                    {{ $t('Theme') + $t('Color') }}
                </span>
                <span>
                    accent: {{ accent }}
                </span>
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round v-for="(item, index) in colorList" :key="index" :size="accent === item ? 'md' : 'xs'"
                        :style="{ border: '1px solid grey', background: item }" class="cursor"
                        @click="setBrand('accent', item)" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span :style="{ color: positive }">
                    {{ $t('Theme') + $t('Color') }}
                </span>
                <span>
                    positive: {{ positive }}
                </span>
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round v-for="(item, index) in colorList" :key="index" :size="positive === item ? 'md' : 'xs'"
                        :style="{ border: '1px solid grey', background: item }" class="cursor"
                        @click="setBrand('positive', item)" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span :style="{ color: negative }">
                    {{ $t('Theme') + $t('Color') }}
                </span>
                <span>
                    negative: {{ negative }}
                </span>
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round v-for="(item, index) in colorList" :key="index" :size="negative === item ? 'md' : 'xs'"
                        :style="{ border: '1px solid grey', background: item }" class="cursor"
                        @click="setBrand('negative', item)" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span :style="{ color: info }">
                    {{ $t('Theme') + $t('Color') }}
                </span>
                <span>
                    info: {{ info }}
                </span>
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round v-for="(item, index) in colorList" :key="index" :size="info === item ? 'md' : 'xs'"
                        :style="{ border: '1px solid grey', background: item }" class="cursor"
                        @click="setBrand('info', item)" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span :style="{ color: warning }">
                    {{ $t('Theme') + $t('Color') }}
                </span>
                <span>
                    warning: {{ warning }}
                </span>
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round v-for="(item, index) in colorList" :key="index" :size="warning === item ? 'md' : 'xs'"
                        :style="{ border: '1px solid grey', background: item }" class="cursor"
                        @click="setBrand('warning', item)" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span :style="{ color: light }">
                    {{ $t('Theme') + $t('Color') }}
                </span>
                <span>
                    light: {{ light }}
                </span>
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round v-for="(item, index) in colorList" :key="index" :size="light === item ? 'md' : 'xs'"
                        :style="{ border: '1px solid grey', background: item }" class="cursor"
                        @click="setBrand('light', item)" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span :style="{ color: dark }">
                    {{ $t('Theme') + $t('Color') }}
                </span>
                <span>
                    dark: {{ dark }}
                </span>
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round v-for="(item, index) in colorList" :key="index" :size="dark === item ? 'md' : 'xs'"
                        :style="{ border: '1px solid grey', background: item }" class="cursor"
                        @click="setBrand('dark', item)" />
                </div>
            </template>
        </q-field>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';
import { setCssVar, getCssVar, Cookies } from 'quasar';
import { ThemeStyleQuasar, ThemeStyleElement, ThemeStyleAnt } from "src/config/default"

const colorList = computed(() => {
    const cl = []
    for (let q in ThemeStyleQuasar) {
        cl.push(ThemeStyleQuasar[q])
    }
    for (let e in ThemeStyleElement) {
        cl.push(ThemeStyleElement[e])
    }
    for (let a in ThemeStyleAnt) {
        cl.push(ThemeStyleAnt[a])
    }
    return [...new Set(cl)]
})

const primary = ref(Cookies.get('gqa-theme-primary') || getCssVar('primary'));
const secondary = ref(Cookies.get('gqa-theme-secondary') || getCssVar('secondary'));
const accent = ref(Cookies.get('gqa-theme-accent') || getCssVar('accent'));
const positive = ref(Cookies.get('gqa-theme-positive') || getCssVar('positive'));
const negative = ref(Cookies.get('gqa-theme-negative') || getCssVar('negative'));
const info = ref(Cookies.get('gqa-theme-info') || getCssVar('info'));
const warning = ref(Cookies.get('gqa-theme-warning') || getCssVar('warning'));
const dark = ref(Cookies.get('gqa-theme-dark') || getCssVar('dark'));
const light = ref(Cookies.get('gqa-theme-light') || getCssVar('light'));

const setBrand = (type, color) => {
    Cookies.set('gqa-theme-' + type, color)
    setCssVar(type, color)
    eval(type).value = color;
}
</script>

<style lang="scss" scoped>
.cursor {
    cursor: pointer;
}
</style>
