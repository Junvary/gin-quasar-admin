export function ArrayOrObject(value) {
    if (Object.prototype.toString.call(value) === "[object Array]") {
        return "Array"
    } else if (Object.prototype.toString.call(value) === '[object Object]') {
        return "Object"
    } else {
        return ""
    }
}