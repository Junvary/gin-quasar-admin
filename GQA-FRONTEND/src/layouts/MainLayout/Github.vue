<template>
    <q-fab round glossy push padding="6px" color="primary" icon="fab fa-github" direction="down" :persistent="false">
        <q-fab-action color="primary" @click="openLink('https://github.com/Junvary/gin-quasar-admin')"
            icon="fab fa-github" external-label label="访问Github" label-position="left" />
        <q-fab-action color="primary" @click="openLink('https://gitee.com/junvary/gin-quasar-admin')"
            icon="fab fa-github" external-label label="访问Gitee" label-position="left" />
    </q-fab>
</template>

<script>
export default {
    name: 'Github',
    methods: {
        openLink(url) {
            window.open(url)
        },
    },
}
</script>