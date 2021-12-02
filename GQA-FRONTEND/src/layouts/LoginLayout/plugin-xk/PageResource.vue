<template>
    <div class="row items-center justify-around gqa-resource" id="gqa-resource">
        <PageResourceDocument class="col" />
        <PageResourceDownload class="col" />
    </div>
</template>

<script>
import PageResourceDocument from './PageResourceDocument'
import PageResourceDownload from './PageResourceDownload'
export default {
    name: 'PageResource',
    components: {
        PageResourceDocument,
        PageResourceDownload,
    },
}
</script>

<style lang="scss" scoped>
.gqa-resource {
    width: 100%;
    padding: 100px 0;
}
</style>