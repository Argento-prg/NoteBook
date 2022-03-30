import { authorize, registrate } from './authRequests'
import { store } from '../storeApp'

export async function authorizeService(login, password) {
    const response = await authorize(login, password)
    if (!response.status) {
        alert(response.message)
        return false
    }
    const setTokenAndName = store.setToken(response.token) && store.setName(response.name)
    return setTokenAndName
}
export async function registrateService(name, login, password) {
    const response = await registrate(name, login, password)
    if (!response.status) {
        alert(response.message)
        return false
    }
    alert('Регистрация прошла успешно!')
    return true
}