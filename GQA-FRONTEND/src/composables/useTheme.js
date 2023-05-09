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
            if (themeStyle.value === 'Element Plus') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Ant Design Vue') {
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
            if (themeStyle.value === 'Element Plus') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Ant Design Vue') {
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
            if (themeStyle.value === 'Element Plus') {
                return 'bg-primary text-white'
            }
            if (themeStyle.value === 'Ant Design Vue') {
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
            if (themeStyle.value === 'Element Plus') {
                return 'bg-white text-primary'
            }
            if (themeStyle.value === 'Ant Design Vue') {
                return 'bg-white text-primary'
            }
        }
    });
    const darkThemeSideBar = computed(() => {
        if ($q.dark.isActive) {
            return ['class', 'bg-grey-10 text-white']
        } else {
            if (themeStyle.value === 'Gin-Quasar-Admin') {
                return ['class', 'bg-white text-black']
            }
            if (themeStyle.value === 'Quasar') {
                return ['class', 'bg-white text-dark']
            }
            if (themeStyle.value === 'Element Plus') {
                return ['style', { backgroundColor: '#545c64', color: 'white' }]
            }
            if (themeStyle.value === 'Ant Design Vue') {
                return ['style', { backgroundColor: '#001529', color: 'white' }]
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
            if (themeStyle.value === 'Element Plus') {
                return 'bg-#545c64 text-yellow'
            }
            if (themeStyle.value === 'Ant Design Vue') {
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