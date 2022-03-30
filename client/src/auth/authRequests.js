import { connInstanse } from '../requestConfig'
import { validate } from './authValidate'

export async function authorize(login, password) {
    const check = validate(login, password)
    if (!check.status) {
        return check
    }
    const url = '/auth/sign-in'
    const data = {
        login,
        password
    }
    let authRes
    try {
        authRes = await connInstanse.post(url, data)
    }
    catch(err) {
        return {
            status: false,
            message: 'Проверьте введённые данные!' //err.response.data.message
        }
    }
    return {
        status: true,
        token: authRes.data.token,
        name: authRes.data.name
    }
}

export async function registrate(name, login, password) {
    const check = validate(login, password, name)
    if (!check.status) {
        return check
    }
    const url = '/auth/sign-up'
    const data = {
        name,
        login,
        password
    }
    let regRes
    try {
        regRes = await connInstanse.post(url, data)
    }
    catch(err) {
        return {
            status: false,
            message: err.response.data.message
        }
    }
    if (!regRes.data.id) {
        return {
            status: false,
            message: 'Неизвестная ошибка'
        }
    }
    return {
        status: true
    }
}