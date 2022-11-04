import { useStorageStore } from 'src/stores/storage';
import { computed } from 'vue';
import { FormatDateTime } from 'src/utils/date'
import GqaDictShow from 'src/components/GqaDictShow'
import GqaShowName from 'src/components/GqaShowName'
import GqaAvatar from 'src/components/GqaAvatar'

export default function useCommon() {
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

    return {
        GqaDictShow,
        GqaShowName,
        GqaAvatar,
        showDateTime,
        gqaFrontend,
        gqaBackend,
        openLink,
    }
}