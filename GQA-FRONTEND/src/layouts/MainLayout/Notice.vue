<template>
    <q-tab name="notifications" icon="notifications" v-if="$q.screen.gt.sm">
        <q-badge color="red" floating>{{ itemsMenu.length }}</q-badge>
        <q-tooltip class="bg-blue">通知</q-tooltip>
        <q-menu
            fit
            anchor="bottom left"
            self="top middle"
            :offset="[93, 0]"
            @show="scrollTarget = $refs.scrollTargetRef"
        >
            <q-item-label header> 我的消息 </q-item-label>
            <q-list
                ref="scrollTargetRef"
                class="scroll"
                style="max-height: 250px; width: 300px"
            >
                <q-infinite-scroll
                    @load="onLoadMenu"
                    :offset="250"
                    :scroll-target="scrollTarget"
                >
                    <q-item
                        clickable
                        v-ripple
                        v-for="(item, index) in itemsMenu"
                        :key="index"
                    >
                        <q-item-section avatar>
                            <q-avatar color="primary" text-color="white">
                                R{{ index + 1 }}
                            </q-avatar>
                        </q-item-section>

                        <q-item-section>
                            <q-item-label
                                >Ruddy Jedrzej {{ index + 1 }}</q-item-label
                            >
                            <q-item-label caption lines="1"
                                >rjedrzej0@discuz{{
                                    index + 1
                                }}.net</q-item-label
                            >
                        </q-item-section>

                        <q-item-section side>
                            <q-icon name="chat_bubble" color="green" />
                        </q-item-section>
                    </q-item>

                    <template v-slot:loading>
                        <div class="text-center q-my-md">
                            <q-spinner-dots color="primary" size="40px" />
                        </div>
                    </template>
                </q-infinite-scroll>
            </q-list>
        </q-menu>
    </q-tab>
</template>

<script>
export default {
    name: "Notice",
    data() {
        return {
            scrollTarget: void 0,
            itemsMenu: [{}, {}, {}, {}, {}, {}, {}],
        };
    },
    methods: {
        onLoadMenu(index, done) {
            if (index > 1) {
                setTimeout(() => {
                    if (this.itemsMenu) {
                        this.itemsMenu.push({}, {}, {}, {}, {}, {}, {});
                        done();
                    }
                }, 2000);
            } else {
                setTimeout(() => {
                    done();
                }, 200);
            }
        },
    },
};
</script>