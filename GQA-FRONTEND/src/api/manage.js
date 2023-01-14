import { api } from 'src/boot/axios'
import { Notify } from 'quasar'
import { Cookies } from 'quasar'

export function getAction(url, params) {
    return api({
        url: url,
        method: 'get',
        params: params
    })
}

export function postAction(url, params) {
    return api({
        url: url,
        method: 'post',
        data: params,
    })
}

export function postBlobAction(url, params) {
    return api({
        url: url,
        method: 'post',
        data: params,
        responseType: 'blob'
    })
}

export function putAction(url, params) {
    return api({
        url: url,
        method: 'put',
        data: params
    })
}

export function deleteAction(url, params) {
    return api({
        url: url,
        method: 'delete',
        data: params
    })
}

export async function downloadAction(url, params, fileName) {
    const data = await postBlobAction(url, params)
    if (!data || data.size === 0) {
        Notify.create({
            type: 'negative',
            message: 'download error!',
        })
        return
    }
    if (typeof window.navigator.msSaveBlob !== 'undefined') {
        window.navigator.msSaveBlob(new Blob([data]), fileName)
    } else {
        let urlHref = window.URL.createObjectURL(new Blob([data]))
        let link = document.createElement('a')
        link.style.display = 'none'
        link.href = urlHref
        link.setAttribute('download', fileName)
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        window.URL.revokeObjectURL(urlHref)
    }
}