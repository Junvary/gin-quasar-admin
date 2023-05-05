import GqaAvatar from 'src/components/common/GqaAvatar/index.vue'
import GqaTreeTd from 'src/components/GqaTreeTd/index.vue'
import GqaFormTop from 'src/components/GqaFormTop/index.vue'

// we globally register our component with Vue
export default ({ app }) => {
    app.component('gqa-avatar', GqaAvatar)
    app.component('gqa-tree-td', GqaTreeTd)
    app.component('gqa-form-top', GqaFormTop)
}