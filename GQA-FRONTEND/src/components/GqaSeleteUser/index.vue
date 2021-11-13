<template>
    <q-input :label="label" v-model="showName" :class="className" readonly>
        <template v-slot:append>
            <q-btn dense color="primary" :label="$t('ComponentSelectUser')" @click="showSelectUserDialog" />
        </template>
    </q-input>
    <SelectUserDialog ref="selectUserDialog" @handleSelectUser="handleSelectUser" :selection="selection" />
</template>

<script>
import SelectUserDialog from './SelectUserDialog'
export default {
    name: 'GqaSeleteUser',
    props: {
        selectUser: {
            type: [Object, Array],
            required: false,
        },
        selectId: {
            type: Number,
            required: false,
        },
        label: {
            type: String,
            required: false,
            default: this.$parent.$t('ComponentSelectUserLabel'),
        },
        className: {
            type: String,
            required: false,
            default: '',
        },
        selection: {
            type: String,
            required: false,
            default: 'single',
        },
    },
    components: {
        SelectUserDialog,
    },
    computed: {
        showName() {
            if (this.selectUser.nickname) {
                return this.selectUser.nickname
            } else if (this.selectUser.realName) {
                return this.selectUser.realName
            } else if (this.selectUser.id) {
                return this.selectUser.id
            } else {
                return ''
            }
        },
    },
    methods: {
        showSelectUserDialog() {
            this.$refs.selectUserDialog.show(this.selectUser)
        },
        handleSelectUser(event) {
            if (this.selection === 'multiple') {
                const ids = []
                for (let i of event) {
                    ids.push(i.id)
                }
                this.$emit('update:selectUser', event)
                this.$emit('update:selectId', ids)
            } else {
                this.$emit('update:selectUser', event[0])
                this.$emit('update:selectId', event[0].id)
            }
        },
    },
}
</script>
