import { useStorageStore } from 'src/stores/storage';
import { computed } from 'vue';


export default function useFrontendBackend() {
    const storageStore = useStorageStore();
    const gqaFrontend = computed(() => storageStore.GetGqaFrontend());
    const gqaBackend = computed(() => storageStore.GetGqaBackend());
    const openLink = (url) => {
        window.open(url)
    }
    return {
        gqaFrontend,
        gqaBackend,
        openLink
    }
}