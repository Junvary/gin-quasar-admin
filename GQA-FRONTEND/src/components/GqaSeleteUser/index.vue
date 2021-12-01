<template>
    <q-input :label="label" v-model="showName" :class="className" readonly>
        <template v-slot:append>
            <q-btn dense color="primary" label="选择" @click="showSelectUserDialog" />
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
            type: [Number, Array],
            required: false,
        },
        label: {
            type: String,
            required: false,
            default: '选择用户',
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
            if (this.selection === 'multiple') {
                let nickname = ''
                let realName = ''
                let id = ''
                for (let u of this.selectUser) {
                    if (u.nickname) {
                        nickname += u.nickname + ' '
                    } else if (u.realName) {
                        realName += u.realName + ' '
                    } else if (u.id) {
                        id += u.id + ' '
                    } else {
                        return ''
                    }
                }
                return nickname || realName || id
            } else {
                if (this.selectUser.nickname) {
                    return this.selectUser.nickname
                } else if (this.selectUser.realName) {
                    return this.selectUser.realName
                } else if (this.selectUser.id) {
                    return this.selectUser.id
                } else {
                    return ''
                }
            }
        },
    },
    methods: {
        showSelectUserDialog() {
            this.$refs.selectUserDialog.show(this.selectUser)
        },
        handleSelectUser(event) {
            console.log(event)
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