<template>
    <div>
        <q-tab
            name="emoji_food_beverage"
            icon="emoji_food_beverage"
            v-if="$q.screen.gt.xs"
            @click="toolbar = !toolbar"
        >
        </q-tab>
        <q-dialog v-model="toolbar">
            <q-card class="wd-500">
                <img src="https://pic1.zhimg.com/v2-b3c3c61be3c8664a4d982d806648b651_r.jpg?source=1940ef5c" />

                <q-list>
                    <q-item clickable>
                        <q-item-section avatar>
                            <q-icon color="primary" name="local_bar" />
                        </q-item-section>

                        <q-item-section>
                            <q-item-label>Bar XYZ</q-item-label>
                            <q-item-label caption>Have a drink.</q-item-label>
                        </q-item-section>
                    </q-item>

                    <q-item clickable>
                        <q-item-section avatar>
                            <q-icon color="red" name="local_gas_station" />
                        </q-item-section>

                        <q-item-section>
                            <q-item-label>Gas Station</q-item-label>
                            <q-item-label caption
                                >Fill your gas tank.</q-item-label
                            >
                        </q-item-section>
                    </q-item>

                    <q-item clickable>
                        <q-item-section avatar>
                            <q-icon color="amber" name="local_movies" />
                        </q-item-section>

                        <q-item-section>
                            <q-item-label>Cinema XYZ</q-item-label>
                            <q-item-label caption>Watch a movie.</q-item-label>
                        </q-item-section>
                    </q-item>
                </q-list>
            </q-card>
        </q-dialog>
    </div>
</template>

<script>
export default {
    name: "Github",
    data() {
        return{
            toolbar: false
        }
        
    },
};
</script>