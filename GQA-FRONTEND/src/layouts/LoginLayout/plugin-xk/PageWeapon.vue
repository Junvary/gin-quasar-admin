<template>
    <div class="row items-center justify-around gqa-weapon" id="gqa-weapon">
        <PageWeaponCode class="col" />
        <PageWeaponNode class="col" />
    </div>
</template>

<script>
import PageWeaponCode from './PageWeaponCode'
import PageWeaponNode from './PageWeaponNode'

export default {
    name: 'PageWeapon',
    components: {
        PageWeaponCode,
        PageWeaponNode,
    },
}
</script>

<style lang="scss" scoped>
.gqa-weapon {
    width: 100%;
    padding: 50px 0;
}
</style>