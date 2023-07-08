export const UpperFirst = (str) => {
    str = str.toLowerCase()
    return str.replace(/\b\w|\s\w/g, fw => {
        return fw.toUpperCase()
    })
}

