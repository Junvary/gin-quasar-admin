import { useStorageStore } from 'src/stores/storage';
import { computed } from 'vue';
import { FormatDateTime } from 'src/utils/date'
import GqaDictShow from 'src/components/GqaDictShow'
import GqaShowName from 'src/components/GqaShowName'
import GqaAvatar from 'src/components/GqaAvatar'
import { useI18n } from 'vue-i18n';

export default function useCommon() {
    const { t } = useI18n()
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
    const selectOptionLabel = (opt) => {
        if (opt.name === 'system' || opt.parent_code === 'system' || opt.parent_code === 'log') {
            return t(opt.title)
        }
        return opt.title
    }
    const selectRouteLabel = (opt) => {
        if (opt.name === 'system' || opt.meta.parent_code === 'system' || opt.meta.parent_code === 'log') {
            return t(opt.meta.title)
        }
        return opt.meta.title
    }

    return {
        GqaDictShow,
        GqaShowName,
        GqaAvatar,
        showDateTime,
        gqaFrontend,
        gqaBackend,
        openLink,
        selectOptionLabel,
        selectRouteLabel,
    }
}