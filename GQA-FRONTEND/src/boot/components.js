import GqaAvatar from 'src/components/common/GqaAvatar/index.vue'
import GqaFormTop from 'src/components/common/GqaFormTop/index.vue'
import GqaTreeTd from 'src/components/common/GqaTreeTd/index.vue'
import GqaVersionMenu from 'src/components/common/GqaVersionMenu/index.vue'

// we globally register our component with Vue
export default ({ app }) => {
    app.component('gqa-avatar', GqaAvatar)
    app.component('gqa-form-top', GqaFormTop)
    app.component('gqa-tree-td', GqaTreeTd)
    app.component('gqa-version-menu', GqaVersionMenu)
}