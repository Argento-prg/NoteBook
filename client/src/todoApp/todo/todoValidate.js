import { maxLengthTitle, maxLengthDescription } from './todoStore'

export function todoValidate(title, description) {
    if (title === -1 && description === -1) {
        return true
    }
    if (title === '' || title.length > maxLengthTitle) {
        return false
    }
    if (description.length > maxLengthDescription) {
        return false
    }
    return true
}
export function idValidate(id) {
    if (id === -1 || !Number.isInteger(id)) {
        return false
    }
    return true
}