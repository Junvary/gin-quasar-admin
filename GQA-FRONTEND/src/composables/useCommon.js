import { useStorageStore } from 'src/stores/storage';
import { computed } from 'vue';
import { FormatDateTime } from 'src/utils/date'
import { useI18n } from 'vue-i18n';


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
    // 首页等允许无token的白名单
    const AllowList = ['/login', '/init-db']
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
    return {
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