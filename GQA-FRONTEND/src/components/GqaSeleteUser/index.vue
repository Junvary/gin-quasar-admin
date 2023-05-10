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

<script setup>
import { computed, ref, toRefs } from 'vue';
import SelectUserDialog from './SelectUserDialog.vue'

const props = defineProps({
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
})
const { selectUser, selectUsername, label, className, selection } = toRefs(props)
const showName = computed(() => {
    if (selection === 'multiple') {
        let nickname = ''
        let realName = ''
        let id = ''
        for (let u of selectUser.value) {
            if (u.nickname) {
                nickname += u.nickname + ' '
            } else if (u.real_name) {
                realName += u.real_name + ' '
            } else if (u.id) {
                id += u.id + ' '
            } else {
                return ''
            }
        }
        return nickname || realName || id
    } else {
        if (selectUser.value?.nickname) {
            return selectUser.value.nickname
        } else if (selectUser.value?.real_name) {
            return selectUser.value.real_name
        } else if (selectUser.value?.id) {
            return selectUser.value.id
        } else {
            return ''
        }
    }
})
const selectUserDialog = ref(null)
const showSelectUserDialog = () => {
    selectUserDialog.value.show(selectUser.value)
}

const emit = defineEmits(['update:selectUser', 'update:selectUsername'])
const handleSelectUser = (event) => {
    if (selection === 'multiple') {
        const usernameList = []
        for (let i of event) {
            usernameList.push(i.username)
        }
        emit('update:selectUser', event)
        emit('update:selectUsername', usernameList)
    } else {
        if (event.length) {
            emit('update:selectUser', event[0])
            emit('update:selectUsername', event[0].username)
        } else {
            emit('update:selectUser', {})
            emit('update:selectUsername', '')
        }
    }
}
</script>
