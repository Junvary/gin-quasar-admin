<template>
    <q-page padding>
        <div class="row" style="width: 95%">
            <div class="col-2 column items-center icon-box" v-for="(item, index) in materialIcons_key" :key="index"
                @click="copy(item)">
                <q-icon :name="item" size="xl" />
                <span>
                    {{item}}
                </span>
            </div>
        </div>

    </q-page>
</template>

<script>
import * as materialIconsSet from '@quasar/extras/material-icons'
import * as fontawesomeSet from '@quasar/extras/fontawesome-v5'
import { copyToClipboard } from 'quasar'

export default {
    name: 'Icon',
    data() {
        return {
            materialIcons_key: [],
        }
    },
    created() {
        this.initMaterial()
    },
    methods: {
        initMaterial() {
            // 获取图标 materialIcons 下划线格式命名集合
            for (const i in materialIconsSet) {
                this.materialIcons_key.push(this.toLowerLine(i))
            }
        },
        toLowerLine(str) {
            if (str.substr(0, 3) === 'mat') {
                let t = str.replace(/([A-Z]|\d+)/g, (a, l) => `_${l.toLowerCase()}`).substring(4)
                switch (t) {
                    case 'crop_32':
                        t = 'crop_3_2'
                        break
                    case 'crop_169':
                        t = 'crop_16_9'
                        break
                    case 'crop_54':
                        t = 'crop_5_4'
                        break
                    case 'crop_75':
                        t = 'crop_7_5'
                        break
                    default:
                        break
                }
                return t
            }
            if (str.substr(0, 2) === 'fa') {
                let t = str.replace(/([A-Z])/g, (a, l) => `-${l.toLowerCase()}`).replace(/-/, ' fa-')
                switch (t) {
                    case 'fab500px':
                        t = 'fab fa-500px'
                        break
                    case 'fas fa-stopwatch20':
                        t = 'fas fa-stopwatch-20'
                        break
                    case 'fab fa-font-awesome-logo-full':
                        t = 'fas fa-stopwatch-20'
                        break
                    case 'far fa-font-awesome-logo-full':
                        t = 'fas fa-stopwatch-20'
                        break
                    default:
                        break
                }
                return t
            }
        },
        copy(e) {
            copyToClipboard(e)
                .then(() => {
                    this.$q.notify({
                        message: '成功复制到剪切板',
                        color: 'green',
                    })
                })
                .catch(() => {
                    // 不支持复制
                    this.$q.notify({
                        message: '复制到剪切板失败',
                        color: 'warming',
                    })
                })
        },
    },
}
</script>

<style lang="scss" scoped>
.icon-box {
    margin-bottom: 10px;
    color: #363f45;
    cursor: pointer;
    &:hover {
        background: #edecec;
        box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    }
}
</style>