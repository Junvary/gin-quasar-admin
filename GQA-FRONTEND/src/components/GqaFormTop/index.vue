<template>
    <div class="row">
        <q-input dense class="col" label="ID" v-model="recordDetail.value.id" stack-label readonly />
        <q-field dense class="col" :label="$t('CreatedAt')" stack-label readonly>
            <template v-slot:control>
                {{ showDateTime(recordDetail.value.created_at) }}
            </template>
        </q-field>
        <q-field dense class="col" :label="$t('CreatedBy')" stack-label readonly>
            <template v-slot:control>
                <GqaShowName v-if="recordDetail.value.created_by_user"
                    :customNameObject="recordDetail.value.created_by_user" />
            </template>
        </q-field>
        <q-field dense class="col" :label="$t('UpdatedAt')" stack-label readonly>
            <template v-slot:control>
                {{ showDateTime(recordDetail.value.updated_at) }}
            </template>
        </q-field>
        <q-field dense class="col" :label="$t('UpdatedBy')" stack-label readonly>
            <template v-slot:control>
                <GqaShowName v-if="recordDetail.value.updated_by_user"
                    :customNameObject="recordDetail.value.updated_by_user" />
            </template>
        </q-field>
    </div>
</template>

<script setup>
import { toRefs } from 'vue';
import GqaShowName from 'src/components/GqaShowName'
import useCommon from 'src/composables/useCommon'

const { showDateTime } = useCommon()
const props = defineProps({
    recordDetail: {
        type: Object,
        required: true,
    }
})
const { recordDetail } = toRefs(props);
</script>