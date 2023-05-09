<template>
    <q-fab dense round flat padding="6px" icon="bi-github" direction="down" :persistent="false">
        <q-fab-action color="primary" @click="openLink('https://github.com/Junvary/gin-quasar-admin')"
            icon="img:icon/github.svg" external-label label="Github" label-position="left" />
        <q-fab-action color="primary" @click="openLink('https://gitee.com/junvary/gin-quasar-admin')"
            icon="img:icon/gitee.svg" external-label label="Gitee" label-position="left" />
    </q-fab>
</template>

<script setup>
const openLink = (url) => {
    window.open(url)
}
</script>
