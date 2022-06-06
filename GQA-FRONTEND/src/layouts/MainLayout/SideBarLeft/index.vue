<template>
    <q-card class="absolute-top text-primary text-bold text-center text-h6 cursor-pointer row items-center"
        style="height: 50px;">
        <q-btn stretch flat @click="changeTopMenu(topMenuItemPlugin)" style="width: 100%;"
            :label="topMenuItemPlugin?.top?.title ? $t(topMenuItemPlugin?.top?.title) : $t('Plugin') + $t('Menu') + '(' + topMenu.filter(tm => tm?.top?.is_plugin === 'yes').length + ')'"
            :class="darkTheme" v-if="topMenu.filter(tm => tm?.top?.is_plugin === 'yes').length">
            <q-btn dense stretch flat round icon="install_desktop" text-color="text-grey-8">
                <q-menu transition-show="flip-right" transition-hide="flip-left">
                    <q-list>
                        <q-item clickable v-close-popup @click="changeTopMenuPlugin(item)"
                            v-for="item in topMenu.filter(tm => tm?.top?.is_plugin === 'yes')">
                            <q-item-section>
                                {{ $t(item.top.title) }}
                            </q-item-section>
                        </q-item>
                    </q-list>
                </q-menu>
            </q-btn>
        </q-btn>
        <q-btn v-else stretch flat style="width: 100%; font-size: 1.1rem; font-weight: bold;" :class="darkTheme">
            {{ $t(topMenuItem.top.title) }}
        </q-btn>
    </q-card>
    <q-scroll-area style="height: calc(100% - 50px); margin-top: 50px;">
        <q-list>
            <template v-for="(item, index) in topMenuItem.treeChildren ? topMenuItem.treeChildren : []" :key="index">
                <SideBarLeftItem :addRoutesItem="item" :initLevel="0" />
            </template>
        </q-list>
    </q-scroll-area>
</template>

<script setup>
import { toRefs } from 'vue';
import SideBarLeftItem from './SideBarLeftItem'
import useDarkTheme from 'src/composables/useDarkTheme';

const { darkTheme } = useDarkTheme()
const props = defineProps({
    topMenuItem: {
        type: Object,
        required: false,
        default: () => {
            return {}
        },
    },
    topMenu: {
        type: Object,
        required: true,

    },
    topMenuItemPlugin: {
        type: Object,
        required: false,
        default: () => {
            return {}
        },
    },
})
const { topMenuItem, topMenu } = toRefs(props)

const emit = defineEmits(['changeTopMenu', 'changeTopMenuPlugin'])
const changeTopMenu = (topMenuItemPlugin) => {
    emit('changeTopMenu', topMenuItemPlugin)
}
const changeTopMenuPlugin = (item) => {
    emit('changeTopMenuPlugin', item)
}
</script>
