import { useQuasar } from "quasar";
import { computed } from "vue";


export default function useDarkTheme() {
    const $q = useQuasar();
    const darkTheme = computed(() => {
        if ($q.dark.isActive) {
            return 'bg-dark text-white'
        } else {
            return 'bg-white text-grey-8'
        }
    });
    const darkThemeSelect = computed(() => {
        if ($q.dark.isActive) {
            return 'bg-grey-9 text-white'
        } else {
            return 'bg-grey-4 text-dark'
        }
    });
    const darkThemeChart = computed(() => {
        if ($q.dark.isActive) {
            return 'dark'
        } else {
            return ''
        }
    });
    return {
        darkTheme,
        darkThemeSelect,
        darkThemeChart
    }
}