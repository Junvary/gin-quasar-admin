<template>
    <div>
        <q-input :label="label || $t('Selete') + $t('User')" v-model="showName" :class="className" readonly>
            <template v-slot:append>
                <q-btn dense color="primary" :label="$t('Select')" @click="showSelectUserDialog" />
            </template>
        </q-input>
        <SelectUserDialog ref="selectUserDialog" @handleSelectUser="handleSelectUser" :selection="selection" />
    </div>
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
        selectUsername: {
            type: [String, Array],
            required: false,
        },
        label: {
            type: String,
            required: false,
            default: function () {
                return this.$t('Selete') + this.$t('User')
            },
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
            if (this.selection === 'multiple') {
                const usernameList = []
                for (let i of event) {
                    usernameList.push(i.username)
                }
                this.$emit('update:selectUser', event)
                this.$emit('update:selectUsername', usernameList)
            } else {
                if (event.length) {
                    this.$emit('update:selectUser', event[0])
                    this.$emit('update:selectUsername', event[0].username)
                } else {
                    this.$emit('update:selectUser', {})
                    this.$emit('update:selectUsername', '')
                }
            }
        },
    },
}
</script>
