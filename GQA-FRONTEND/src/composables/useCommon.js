import { useStorageStore } from 'src/stores/storage';
import { computed, onMounted, ref } from 'vue';
import { FormatDateTime } from 'src/utils/date'
import { useI18n } from 'vue-i18n';
import GqaDictShow from 'src/components/GqaDictShow'
import GqaShowName from 'src/components/GqaShowName'
import GqaAvatar from 'src/components/GqaAvatar'

export default function useCommon() {
    const { t } = useI18n
    const storageStore = useStorageStore();
    const gqaFrontend = computed(() => storageStore.GetGqaFrontend());
    const gqaBackend = computed(() => storageStore.GetGqaBackend());

    const openLink = (url) => {
        window.open(url)
    }
    const showDateTime = computed(() => {
        return (datetime) => {
            return FormatDateTime(datetime)
        }
    })

    /* 配置项 */
    // 浏览器console打印内容
    const gqaLogo = () => {
        console.info('欢迎使用Gin-Quasar-Admin!')
        console.info('项目地址: https://github.com/Junvary/gin-quasar-admin ')
        console.info('欢迎交流, 感谢Star!')
    }
    // 首页等允许无token的白名单
    const AllowList = ['/login',]
    // 没有用户名的时候使用这个名字
    const GqaDefaultUsername = () => t('GqaDefaultUsername')
    // 没有头像配置的时候使用这个头像
    const GqaDefaultAvatar = "gqa128.png"
    // 没有网站前台配置的时候用这个配置
    const GqaFrontendDefault = {
        mainTitle: "Gin-Quasar-Admin",
        subTitle: "Gin-Quasar-Admin",
        webDescribe: "Be the change you want to see in the world.",
        showGit: "yes"
    }
    // 没有网站后台配置的时候用这个配置
    const GqaBackendDefault = {}
    /* 配置项 */
    return {
        gqaLogo,
        GqaDictShow,
        GqaShowName,
        GqaAvatar,
        showDateTime,
        gqaFrontend,
        gqaBackend,
        openLink,
        AllowList,
        GqaDefaultUsername,
        GqaDefaultAvatar,
        GqaFrontendDefault,
        GqaBackendDefault
    }
}