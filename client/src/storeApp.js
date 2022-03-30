import { writable } from 'svelte/store'

class Store {
    //ключ параметра сессии Токена
    _keyToken = 'token'
    //установка значения токена
    setToken(value) {
        if (value === '') {
            return false
        }
        sessionStorage.setItem(this._keyToken, value)
        isLoggedStore.update(value => true)
        return true
    }
    //получение токена
    getToken() {
        return sessionStorage.getItem(this._keyToken)
    }
    //удаление токена
    deleteToken() {
        sessionStorage.removeItem(this._keyToken)
    }
    //ключ параметра сессии Имени
    _keyName = 'name'
    //установка значения имени
    setName(value) {
        if (value === '') {
            return false
        }
        sessionStorage.setItem(this._keyName, value)
        return true
    }
    //получение имени
    getName() {
        return sessionStorage.getItem(this._keyName)
    }
    //удаление имени
    deleteName() {
        sessionStorage.removeItem(this._keyName)
    }
}

export const store = new Store()
export const isLoggedStore = function() {
    let value = false
    if (store.getToken()) {
        value = true
    }
    return writable(value)
}()