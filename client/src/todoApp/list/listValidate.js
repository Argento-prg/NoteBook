import { maxLengthTitle, maxLengthDescription } from './listStore'

export function listValidate(title, description) {
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