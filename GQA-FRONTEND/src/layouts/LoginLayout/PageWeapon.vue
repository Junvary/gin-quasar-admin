<template>
    <div class="gqa-weapon" id="gqa-weapon">
        <PageWeaponCode />
        <PageWeaponNode />
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
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 50px 0;
}
</style>