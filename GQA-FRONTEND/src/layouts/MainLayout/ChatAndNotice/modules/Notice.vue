<template>
    <q-btn dense round glossy push color="primary" icon="notifications">
        <q-badge color="negative" floating>4</q-badge>
        <q-menu>
            <q-card>
                <q-tabs v-model="messageType" dense class="text-grey" active-color="primary" indicator-color="primary"
                    align="justify" narrow-indicator style="padding: 10px">
                    <q-tab name="system" :label="$t('NoticeSystem')">
                        <q-badge color="negative" floating>4</q-badge>
                    </q-tab>
                    <q-tab name="message" :label="$t('NoticeMessage')">
                        <q-badge color="negative" floating>4</q-badge>
                    </q-tab>

                    <q-tab name="todo" :label="$t('NoticeToDo')">
                        <q-badge color="negative" floating>4</q-badge>
                    </q-tab>
                </q-tabs>

                <q-separator />

                <q-tab-panels v-model="messageType" animated>
                    <q-tab-panel style="padding: 0" name="message">
                        <q-list bordered separator style="min-width: 300px">
                            <q-item clickable v-ripple v-for="(item, index) in message" :key="index">
                                <q-item-section avatar>
                                    <q-icon color="primary" name="message" />
                                </q-item-section>

                                <q-item-section>
                                    {{item.title}}
                                </q-item-section>
                            </q-item>
                        </q-list>
                        <q-item clickable v-ripple class="text-center">
                            <q-item-section>{{ $t('CheckAll') }}</q-item-section>
                        </q-item>
                    </q-tab-panel>

                    <q-tab-panel style="padding: 0" name="system">
                        <q-list bordered separator style="min-width: 300px">
                            <q-item clickable v-ripple v-for="(item, index) in system" :key="index">
                                <q-item-section avatar>
                                    <q-icon color="primary" name="notifications" />
                                </q-item-section>

                                <q-item-section>
                                    {{item.title}}
                                </q-item-section>
                            </q-item>
                        </q-list>
                        <q-item clickable v-ripple class="text-center">
                            <q-item-section>{{ $t('CheckAll') }}</q-item-section>
                        </q-item>
                    </q-tab-panel>

                    <q-tab-panel style="padding: 0" name="todo">
                        <q-list bordered separator style="min-width: 300px">
                            <q-item clickable v-ripple v-for="(item, index) in todo" :key="index">
                                <q-item-section avatar>
                                    <q-icon color="primary" name="list" />
                                </q-item-section>

                                <q-item-section>
                                    {{item.title}}
                                </q-item-section>
                            </q-item>
                        </q-list>
                        <q-item clickable v-ripple class="text-center">
                            <q-item-section>{{ $t('CheckAll') }}</q-item-section>
                        </q-item>
                    </q-tab-panel>
                </q-tab-panels>
            </q-card>

        </q-menu>
    </q-btn>

</template>

<script>
export default {
    name: 'Notice',
    data() {
        return {
            messageType: 'system',
            message: [],
            system: [],
            todo: [],
        }
    },
    created() {
        const m = {
            title: this.$t('NoticeNew'),
        }
        const s = {
            title: this.$t('NoticeSystemNew'),
        }
        const t = {
            title: this.$t('NoticeToDoNew'),
        }
        this.message.push(m, m, m, m)
        this.system.push(s, s, s, s)
        this.todo.push(t, t, t, t)
    },
}
</script>
