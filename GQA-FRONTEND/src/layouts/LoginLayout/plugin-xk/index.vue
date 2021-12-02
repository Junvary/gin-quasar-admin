<template>
    <div>
        <page-header />
        <page-news />
        <page-project />
        <page-weapon />
        <page-honour />
        <page-resource />
    </div>
</template>

<script>
import PageHeader from './PageHeader'
import PageNews from './PageNews'
import PageProject from './PageProject'
import PageHonour from './PageHonour'
import PageWeapon from './PageWeapon'
import PageResource from './PageResource'

export default {
    name: 'plugin-xk',
    components: {
        PageHeader,
        PageNews,
        PageProject,
        PageHonour,
        PageWeapon,
        PageResource,
    },
}
</script>