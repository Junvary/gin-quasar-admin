<template>
    <q-page style="margin: 0 16px">

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="searchParams.name" label="姓名" />
            <q-input style="width: 20%" v-model="searchParams.username" label="账号" />
            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="uuid" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增用户" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-avatar="props">
                <q-td :props="props">
                    <gqa-avatar :src="props.row.avatar" />
                </q-td>
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <gqa-status :status="props.row.status" />
                </q-td>
            </template>

            <!-- <template v-slot:body-cell-dept="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        {{ getDeptName(props.row.dept) }}
                    </div>
                </q-td>
            </template>

            <template v-slot:body-cell-role="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-chip color="primary" text-color="white" v-for="(item, index) in props.row.role" :key="index">
                            {{ getRoleName(item) }}
                        </q-chip>

                    </div>
                </q-td>
            </template> -->

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" label="编辑" />
                        <q-btn color="negative" @click="handleDelete(props.row)" label="删除" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <add-or-edit-dialog ref="addOrEditDialog" @emitAddOrEdit="emitAddOrEdit" @handleFinish="handleFinish" />
    </q-page>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import addOrEditDialog from './modules/addOrEditDialog'
import GqaAvatar from 'src/components/GqaAvatar'
import GqaStatus from 'src/components/GqaStatus'

export default {
    name: 'User',
    mixins: [tableDataMixin],
    components: {
        addOrEditDialog,
        GqaAvatar,
        GqaStatus,
    },
    data() {
        return {
            url: {
                list: 'user/user-list',
                delete: 'user/user-delete',
            },
            columns: [
                { name: 'sort', align: 'center', label: '排序', field: 'sort' },
                { name: 'avatar', align: 'center', label: '头像', field: 'avatar' },
                { name: 'username', align: 'center', label: '账号', field: 'username' },
                { name: 'nickname', align: 'center', label: '昵称', field: 'nickname' },
                { name: 'realName', align: 'center', label: '真实姓名', field: 'realName' },
                { name: 'gender', align: 'center', label: '性别', field: 'gender' },
                { name: 'mobile', align: 'center', label: '手机', field: 'mobile' },
                { name: 'email', align: 'center', label: '邮箱', field: 'email' },
                { name: 'dept', align: 'center', label: '部门', field: 'dept' },
                { name: 'status', align: 'center', label: '状态', field: 'status' },
                { name: 'actions', align: 'center', label: '操作', field: 'actions' },
            ],
        }
    },
}
</script>