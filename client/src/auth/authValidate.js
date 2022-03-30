import { 
    maxLengthName, 
    maxLengthLogin, 
    minLengthPassword, 
    maxLengthPassword 
} from './authStore'

export function validate(login, password, name = 0) {
    if (name !== 0) {
        if (name === '' || name.length > maxLengthName) {
            return {
                status: false,
                message: `Имя пользователя является обязательным 
                и не должно превышать ${maxLengthName} символов`
            }
        }
    }
    if (login === '' || login.length > maxLengthLogin) {
        return {
            status: false,
            message: `Логин является обязательным, должен быть уникальным
            и не должен превышать ${maxLengthLogin} символов`
        }
    }
    if (password.length < minLengthPassword || password.length > maxLengthPassword) {
        return {
            status: false,
            message: `Пароль является обязательным, должен быть 
            не менее ${minLengthPassword} символов и не более ${maxLengthPassword} символов`
        }
    }
    return {status: true}
}