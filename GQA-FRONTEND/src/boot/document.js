import { computed, watch } from 'vue';
import { useStorageStore } from 'src/stores/storage';
import useConfig from 'src/composables/useConfig';

const { GqaFrontendDefault } = useConfig()
const storageStore = useStorageStore();
const gqaFrontend = computed(() => storageStore.GetGqaFrontend());

const createLink = () => {
    const toDelete = document.getElementsByName('gqa-link-href')
    if (toDelete && toDelete.length) {
        document.getElementsByTagName('head')[0].removeChild(toDelete[0])
    }
    const gqaLink = document.createElement('link')
    gqaLink.type = 'image/ico'
    gqaLink.rel = 'icon'
    gqaLink.setAttribute('name', 'gqa-link-href')
    if (gqaFrontend.value.favicon && gqaFrontend.value.favicon !== '') {
        const favicon = process.env.API + gqaFrontend.value.favicon.substring(11)
        gqaLink.href = favicon
    } else {
        gqaLink.href = 'favicon.ico'
    }
    document.getElementsByTagName('head')[0].appendChild(gqaLink)
}

watch(gqaFrontend, (newValue) => {
    document.title = newValue.subTitle
    createLink()
})

export default async () => {
    document.title = gqaFrontend.value.subTitle || GqaFrontendDefault.subTitle
    createLink()
}