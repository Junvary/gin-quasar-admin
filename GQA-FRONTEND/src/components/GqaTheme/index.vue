<template>
    <div>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span>{{ $t('ComponentThemeColor') }} primary:{{ primary }}</span>
                <q-avatar round size="xs" color="primary" style="border: 1px solid grey; margin-left: 5px" />
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round :size="primary === '#1976D2' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1976D2" class="cursor"
                        @click="setBrand('primary', '#1976D2')" />
                    <q-avatar round :size="primary === '#26A69A' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #26A69A" class="cursor"
                        @click="setBrand('primary', '#26A69A')" />
                    <q-avatar round :size="primary === '#9C27B0' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #9C27B0" class="cursor"
                        @click="setBrand('primary', '#9C27B0')" />
                    <q-avatar round :size="primary === '#21BA45' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #21BA45" class="cursor"
                        @click="setBrand('primary', '#21BA45')" />
                    <q-avatar round :size="primary === '#C10015' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #C10015" class="cursor"
                        @click="setBrand('primary', '#C10015')" />
                    <q-avatar round :size="primary === '#31CCEC' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #31CCEC" class="cursor"
                        @click="setBrand('primary', '#31CCEC')" />
                    <q-avatar round :size="primary === '#F2C037' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #F2C037" class="cursor"
                        @click="setBrand('primary', '#F2C037')" />
                    <q-avatar round :size="primary === '#FFFFFF' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #FFFFFF" class="cursor"
                        @click="setBrand('primary', '#FFFFFF')" />
                    <q-avatar round :size="primary === '#1D1D1D' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1D1D1D" class="cursor"
                        @click="setBrand('primary', '#1D1D1D')" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span>{{ $t('ComponentThemeColor') }} secondary:{{ secondary }}</span>
                <q-avatar round size="xs" color="secondary" style="border: 1px solid grey; margin-left: 5px" />
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round :size="secondary === '#1976D2' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1976D2" class="cursor"
                        @click="setBrand('secondary', '#1976D2')" />
                    <q-avatar round :size="secondary === '#26A69A' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #26A69A" class="cursor"
                        @click="setBrand('secondary', '#26A69A')" />
                    <q-avatar round :size="secondary === '#9C27B0' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #9C27B0" class="cursor"
                        @click="setBrand('secondary', '#9C27B0')" />
                    <q-avatar round :size="secondary === '#21BA45' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #21BA45" class="cursor"
                        @click="setBrand('secondary', '#21BA45')" />
                    <q-avatar round :size="secondary === '#C10015' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #C10015" class="cursor"
                        @click="setBrand('secondary', '#C10015')" />
                    <q-avatar round :size="secondary === '#31CCEC' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #31CCEC" class="cursor"
                        @click="setBrand('secondary', '#31CCEC')" />
                    <q-avatar round :size="secondary === '#F2C037' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #F2C037" class="cursor"
                        @click="setBrand('secondary', '#F2C037')" />
                    <q-avatar round :size="secondary === '#FFFFFF' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #FFFFFF" class="cursor"
                        @click="setBrand('secondary', '#FFFFFF')" />
                    <q-avatar round :size="secondary === '#1D1D1D' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1D1D1D" class="cursor"
                        @click="setBrand('secondary', '#1D1D1D')" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span>{{ $t('ComponentThemeColor') }} accent:{{ accent }}</span>
                <q-avatar round size="xs" color="accent" style="border: 1px solid grey; margin-left: 5px" />
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round :size="accent === '#1976D2' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1976D2" class="cursor"
                        @click="setBrand('accent', '#1976D2')" />
                    <q-avatar round :size="accent === '#26A69A' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #26A69A" class="cursor"
                        @click="setBrand('accent', '#26A69A')" />
                    <q-avatar round :size="accent === '#9C27B0' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #9C27B0" class="cursor"
                        @click="setBrand('accent', '#9C27B0')" />
                    <q-avatar round :size="accent === '#21BA45' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #21BA45" class="cursor"
                        @click="setBrand('accent', '#21BA45')" />
                    <q-avatar round :size="accent === '#C10015' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #C10015" class="cursor"
                        @click="setBrand('accent', '#C10015')" />
                    <q-avatar round :size="accent === '#31CCEC' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #31CCEC" class="cursor"
                        @click="setBrand('accent', '#31CCEC')" />
                    <q-avatar round :size="accent === '#F2C037' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #F2C037" class="cursor"
                        @click="setBrand('accent', '#F2C037')" />
                    <q-avatar round :size="accent === '#FFFFFF' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #FFFFFF" class="cursor"
                        @click="setBrand('accent', '#FFFFFF')" />
                    <q-avatar round :size="accent === '#1D1D1D' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1D1D1D" class="cursor"
                        @click="setBrand('accent', '#1D1D1D')" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span>{{ $t('ComponentThemeColor') }} positive:{{ positive }}</span>
                <q-avatar round size="xs" color="positive" style="border: 1px solid grey; margin-left: 5px" />
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round :size="positive === '#1976D2' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1976D2" class="cursor"
                        @click="setBrand('positive', '#1976D2')" />
                    <q-avatar round :size="positive === '#26A69A' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #26A69A" class="cursor"
                        @click="setBrand('positive', '#26A69A')" />
                    <q-avatar round :size="positive === '#9C27B0' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #9C27B0" class="cursor"
                        @click="setBrand('positive', '#9C27B0')" />
                    <q-avatar round :size="positive === '#21BA45' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #21BA45" class="cursor"
                        @click="setBrand('positive', '#21BA45')" />
                    <q-avatar round :size="positive === '#C10015' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #C10015" class="cursor"
                        @click="setBrand('positive', '#C10015')" />
                    <q-avatar round :size="positive === '#31CCEC' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #31CCEC" class="cursor"
                        @click="setBrand('positive', '#31CCEC')" />
                    <q-avatar round :size="positive === '#F2C037' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #F2C037" class="cursor"
                        @click="setBrand('positive', '#F2C037')" />
                    <q-avatar round :size="positive === '#FFFFFF' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #FFFFFF" class="cursor"
                        @click="setBrand('positive', '#FFFFFF')" />
                    <q-avatar round :size="positive === '#1D1D1D' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1D1D1D" class="cursor"
                        @click="setBrand('positive', '#1D1D1D')" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span>{{ $t('ComponentThemeColor') }} negative:{{ negative }}</span>
                <q-avatar round size="xs" color="negative" style="border: 1px solid grey; margin-left: 5px" />
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round :size="negative === '#1976D2' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1976D2" class="cursor"
                        @click="setBrand('negative', '#1976D2')" />
                    <q-avatar round :size="negative === '#26A69A' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #26A69A" class="cursor"
                        @click="setBrand('negative', '#26A69A')" />
                    <q-avatar round :size="negative === '#9C27B0' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #9C27B0" class="cursor"
                        @click="setBrand('negative', '#9C27B0')" />
                    <q-avatar round :size="negative === '#21BA45' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #21BA45" class="cursor"
                        @click="setBrand('negative', '#21BA45')" />
                    <q-avatar round :size="negative === '#C10015' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #C10015" class="cursor"
                        @click="setBrand('negative', '#C10015')" />
                    <q-avatar round :size="negative === '#31CCEC' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #31CCEC" class="cursor"
                        @click="setBrand('negative', '#31CCEC')" />
                    <q-avatar round :size="negative === '#F2C037' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #F2C037" class="cursor"
                        @click="setBrand('negative', '#F2C037')" />
                    <q-avatar round :size="negative === '#FFFFFF' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #FFFFFF" class="cursor"
                        @click="setBrand('negative', '#FFFFFF')" />
                    <q-avatar round :size="negative === '#1D1D1D' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1D1D1D" class="cursor"
                        @click="setBrand('negative', '#1D1D1D')" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span>{{ $t('ComponentThemeColor') }} info:{{ info }}</span>
                <q-avatar round size="xs" color="info" style="border: 1px solid grey; margin-left: 5px" />
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round :size="info === '#1976D2' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1976D2" class="cursor"
                        @click="setBrand('info', '#1976D2')" />
                    <q-avatar round :size="info === '#26A69A' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #26A69A" class="cursor"
                        @click="setBrand('info', '#26A69A')" />
                    <q-avatar round :size="info === '#9C27B0' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #9C27B0" class="cursor"
                        @click="setBrand('info', '#9C27B0')" />
                    <q-avatar round :size="info === '#21BA45' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #21BA45" class="cursor"
                        @click="setBrand('info', '#21BA45')" />
                    <q-avatar round :size="info === '#C10015' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #C10015" class="cursor"
                        @click="setBrand('info', '#C10015')" />
                    <q-avatar round :size="info === '#31CCEC' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #31CCEC" class="cursor"
                        @click="setBrand('info', '#31CCEC')" />
                    <q-avatar round :size="info === '#F2C037' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #F2C037" class="cursor"
                        @click="setBrand('info', '#F2C037')" />
                    <q-avatar round :size="info === '#FFFFFF' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #FFFFFF" class="cursor"
                        @click="setBrand('info', '#FFFFFF')" />
                    <q-avatar round :size="info === '#1D1D1D' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1D1D1D" class="cursor"
                        @click="setBrand('info', '#1D1D1D')" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span>{{ $t('ComponentThemeColor') }} warning:{{ warning }}</span>
                <q-avatar round size="xs" color="warning" style="border: 1px solid grey; margin-left: 5px" />
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round :size="warning === '#1976D2' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1976D2" class="cursor"
                        @click="setBrand('warning', '#1976D2')" />
                    <q-avatar round :size="warning === '#26A69A' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #26A69A" class="cursor"
                        @click="setBrand('warning', '#26A69A')" />
                    <q-avatar round :size="warning === '#9C27B0' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #9C27B0" class="cursor"
                        @click="setBrand('warning', '#9C27B0')" />
                    <q-avatar round :size="warning === '#21BA45' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #21BA45" class="cursor"
                        @click="setBrand('warning', '#21BA45')" />
                    <q-avatar round :size="warning === '#C10015' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #C10015" class="cursor"
                        @click="setBrand('warning', '#C10015')" />
                    <q-avatar round :size="warning === '#31CCEC' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #31CCEC" class="cursor"
                        @click="setBrand('warning', '#31CCEC')" />
                    <q-avatar round :size="warning === '#F2C037' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #F2C037" class="cursor"
                        @click="setBrand('warning', '#F2C037')" />
                    <q-avatar round :size="warning === '#FFFFFF' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #FFFFFF" class="cursor"
                        @click="setBrand('warning', '#FFFFFF')" />
                    <q-avatar round :size="warning === '#1D1D1D' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1D1D1D" class="cursor"
                        @click="setBrand('warning', '#1D1D1D')" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span>{{ $t('ComponentThemeColor') }} light:{{ light }}</span>
                <q-avatar round size="xs" color="light" style="border: 1px solid grey; margin-left: 5px" />
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round :size="light === '#1976D2' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1976D2" class="cursor"
                        @click="setBrand('dark', '#1976D2')" />
                    <q-avatar round :size="light === '#26A69A' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #26A69A" class="cursor"
                        @click="setBrand('light', '#26A69A')" />
                    <q-avatar round :size="light === '#9C27B0' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #9C27B0" class="cursor"
                        @click="setBrand('light', '#9C27B0')" />
                    <q-avatar round :size="light === '#21BA45' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #21BA45" class="cursor"
                        @click="setBrand('light', '#21BA45')" />
                    <q-avatar round :size="light === '#C10015' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #C10015" class="cursor"
                        @click="setBrand('light', '#C10015')" />
                    <q-avatar round :size="light === '#31CCEC' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #31CCEC" class="cursor"
                        @click="setBrand('light', '#31CCEC')" />
                    <q-avatar round :size="light === '#F2C037' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #F2C037" class="cursor"
                        @click="setBrand('light', '#F2C037')" />
                    <q-avatar round :size="light === '#FFFFFF' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #FFFFFF" class="cursor"
                        @click="setBrand('light', '#FFFFFF')" />
                    <q-avatar round :size="light === '#1D1D1D' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1D1D1D" class="cursor"
                        @click="setBrand('light', '#1D1D1D')" />
                </div>
            </template>
        </q-field>
        <q-field label-slot stack-label>
            <template v-slot:label>
                <span>{{ $t('ComponentThemeColor') }} dark:{{ dark }}</span>
                <q-avatar round size="xs" color="dark" style="border: 1px solid grey; margin-left: 5px" />
            </template>
            <template v-slot:control>
                <div class="q-gutter-md q-ma-none">
                    <q-avatar round :size="dark === '#1976D2' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1976D2" class="cursor"
                        @click="setBrand('dark', '#1976D2')" />
                    <q-avatar round :size="dark === '#26A69A' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #26A69A" class="cursor"
                        @click="setBrand('dark', '#26A69A')" />
                    <q-avatar round :size="dark === '#9C27B0' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #9C27B0" class="cursor"
                        @click="setBrand('dark', '#9C27B0')" />
                    <q-avatar round :size="dark === '#21BA45' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #21BA45" class="cursor"
                        @click="setBrand('dark', '#21BA45')" />
                    <q-avatar round :size="dark === '#C10015' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #C10015" class="cursor"
                        @click="setBrand('dark', '#C10015')" />
                    <q-avatar round :size="dark === '#31CCEC' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #31CCEC" class="cursor"
                        @click="setBrand('dark', '#31CCEC')" />
                    <q-avatar round :size="dark === '#F2C037' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #F2C037" class="cursor"
                        @click="setBrand('dark', '#F2C037')" />
                    <q-avatar round :size="dark === '#FFFFFF' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #FFFFFF" class="cursor"
                        @click="setBrand('dark', '#FFFFFF')" />
                    <q-avatar round :size="dark === '#1D1D1D' ? 'sm' : 'xs'"
                        style="border: 1px solid grey; background: #1D1D1D" class="cursor"
                        @click="setBrand('dark', '#1D1D1D')" />
                </div>
            </template>
        </q-field>

    </div>

</template>

<script setup>
import { ref } from 'vue';
import { setCssVar, getCssVar, Cookies } from 'quasar';

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
