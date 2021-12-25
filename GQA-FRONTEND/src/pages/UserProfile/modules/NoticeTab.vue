<template>
    <q-tabs v-model="messageType" dense class="text-grey" active-color="primary" indicator-color="primary"
        align="justify" narrow-indicator style="width: 100%">
        <q-tab name="system" :label="$t('NoticeSystem')" />
        <q-tab name="message" :label="$t('NoticeMessage')" />
        <q-tab name="todoNote" :label="$t('TodoNote')" />
    </q-tabs>
    <q-tab-panels v-model="messageType" animated style="width: 100%">
        <q-tab-panel name="system">
            <NoticeSystem />
        </q-tab-panel>

        <q-tab-panel name="message">
            <NoticeMessage />
        </q-tab-panel>

        <q-tab-panel name="todoNote">
            <NoticeTodoNote />
        </q-tab-panel>
    </q-tab-panels>
</template>

<script>
import NoticeSystem from './NoticeSystem'
import NoticeMessage from './NoticeMessage'
import NoticeTodoNote from './NoticeTodoNote'

export default {
    name: 'NoticeTab',
    components: {
        NoticeSystem,
        NoticeMessage,
        NoticeTodoNote,
    },
    data() {
        return {
            messageType: 'system',
        }
    },
    methods: {
        changeMessageType(type) {
            this.messageType = type
        },
    },
}
</script>