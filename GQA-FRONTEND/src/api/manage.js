import { api } from 'src/boot/axios'

// export const getAction = (url, params) => api.get(url, params)
// export const postAction = (url, params) => api.post(url, params)
// export const putAction = (url, params) => api.put(url, params)
// export const deleteAction = (url, params) => api.delete(url, params)

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
        data: params
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