import { useQuasar } from "quasar";
import { computed } from "vue";
import { useSettingStore } from "src/stores/setting";


export default function useTheme() {
    const $q = useQuasar();
    const settingStore = useSettingStore()
    const themeStyle = computed(() => settingStore.GetThemeStyle())

    const darkTheme = computed(() => {
        if ($q.dark.isActive) {
            return 'bg-dark text-white'
        } else {
            if (themeStyle.value === 'Gin-Quasar-Admin') {
                return 'bg-white text-black'
            }
            if (themeStyle.value === 'Quasar') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Element') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Ant Design') {
                return 'bg-primary text-white'
            }
        }
    });
    const darkThemeLoginHelp = computed(() => {
        if ($q.dark.isActive) {
            return 'bg-dark text-white'
        } else {
            if (themeStyle.value === 'Gin-Quasar-Admin') {
                return 'text-dark'
            }
            if (themeStyle.value === 'Quasar') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Element') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Ant Design') {
                return 'bg-primary text-white'
            }
        }
    });
    const darkThemeLoginCard = computed(() => {
        if ($q.dark.isActive) {
            return 'bg-dark text-white'
        } else {
            if (themeStyle.value === 'Gin-Quasar-Admin') {
                return 'bg-white text-dark'
            }
            if (themeStyle.value === 'Quasar') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Element') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Ant Design') {
                return 'bg-primary text-white'
            }
        }
    });
    const darkThemeTab = computed(() => {
        if ($q.dark.isActive) {
            return 'bg-dark text-white'
        } else {
            if (themeStyle.value === 'Gin-Quasar-Admin') {
                return 'bg-white text-dark'
            }
            if (themeStyle.value === 'Quasar') {
                return 'bg-white text-primary'
            }
            if (themeStyle.value === 'Element') {
                return 'bg-white text-primary'
            }
            if (themeStyle.value === 'Ant Design') {
                return 'bg-white text-primary'
            }
        }
    });
    const darkThemeSideBar = computed(() => {
        if ($q.dark.isActive) {
            return 'bg-grey-10 text-white'
        } else {
            if (themeStyle.value === 'Gin-Quasar-Admin') {
                return 'bg-white text-black'
            }
            if (themeStyle.value === 'Quasar') {
                return 'bg-white text-dark'
            }
            if (themeStyle.value === 'Element') {
                return 'bg-#545c64 text-white'
            }
            if (themeStyle.value === 'Ant Design') {
                return 'bg-#000c17 text-white'
            }
        }
    })
    const darkThemeSelect = computed(() => {
        if ($q.dark.isActive) {
            return 'bg-grey-8 text-white'
        } else {
            if (themeStyle.value === 'Gin-Quasar-Admin') {
                return 'bg-grey-5 text-dark'
            }
            if (themeStyle.value === 'Quasar') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Element') {
                return 'bg-#545c64 text-yellow'
            }
            if (themeStyle.value === 'Ant Design') {
                return 'bg-primary text-white'
            }
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
        darkThemeLoginHelp,
        darkThemeLoginCard,
        darkThemeTab,
        darkThemeSideBar,
        darkThemeSelect,
        darkThemeChart
    }
}