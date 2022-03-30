import { store } from '../../storeApp'
import { connInstanse } from '../../requestConfig'
import { todoValidate, idValidate } from './todoValidate'

export async function getAllTodos(listId) {
    if (!idValidate(listId)) {
        return {
            status: false,
            message: 'Ошибка валидации id списка'
        }
    }
    const url = `/api/lists/${listId}/todos/`
    const headers = { 'Authorization': `Bearer ${store.getToken()}` }
    let todosRes
    try {
        todosRes = await connInstanse.get(url, { headers })
    }
    catch(err) {
        return {
            status: false,
            message: err.response.data.message
        }
    }
    let todosData = []
    if (todosRes.data !== null) {
        todosData = todosRes.data
    }
    return { 
        status: true, 
        data: todosData
    }
}
export async function getTodoById(listId, id) {
    if (!idValidate(listId)) {
        return {
            status: false,
            message: 'Ошибка валидации id списка'
        }
    }
    if (!idValidate(id)) {
        return {
            status: false,
            message: 'Ошибка валидации id задачи'
        }
    }
    const url = `/api/lists/${listId}/todos/${id}`
    const headers = { 'Authorization': `Bearer ${store.getToken()}` }
    let todoRes
    try {
        todoRes = await connInstanse.get(url, { headers })
    }
    catch(err) {
        return {
            status: false,
            message: err.response.data.message
        }
    }
    return {
        status: true,
        data: todoRes.data
    }
}
export async function createTodo(listId, title, description, importance) {
    if(!todoValidate(title, description)) {
        return {
            status: false,
            message: 'Некорректные данные'
        }
    }
    if (!idValidate(listId)) {
        return {
            status: false,
            message: 'Ошибка валидации id списка'
        }
    }
    const url = `/api/lists/${listId}/todos/`
    const headers = { 'Authorization': `Bearer ${store.getToken()}` }
    const data = {
        title,
        description,
        importance,
        done: false,
    }
    let createRes
    try {
        createRes = await connInstanse.post(url, data, { headers })
    }
    catch(err) {
        return {
            status: false,
            message: err.response.data.message
        }
    }
    return {
        status: true
    }
}
export async function updateTodo(listId, id, title, description, importance, done) {
    if(!todoValidate(title, description)) {
        return {
            status: false,
            message: 'Некорректные данные'
        }
    }
    if (!idValidate(listId)) {
        return {
            status: false,
            message: 'Ошибка валидации id списка'
        }
    }
    if (!idValidate(id)) {
        return {
            status: false,
            message: 'Ошибка валидации id задачи'
        }
    }
    const url = `/api/lists/${listId}/todos/${id}`
    const headers = { 'Authorization': `Bearer ${store.getToken()}` }
    const data = {}
    if (title !== -1) {
        data.title = title
    }
    if (description !== -1) {
        data.description = description
    }
    if (importance !== -1) {
        data.importance = importance
    }
    if (done !== -1) {
        data.done = done
    }
    let updateRes
    try {
        updateRes = await connInstanse.put(url, data, { headers })
    }
    catch(err) {
        return {
            status: false,
            message: err.response.data.message
        }
    }
    return {
        status: updateRes.data.status
    }
}
export async function deleteTodo(listId, id) {
    if (!idValidate(listId)) {
        return {
            status: false,
            message: 'Ошибка валидации id списка'
        }
    }
    if (!idValidate(id)) {
        return {
            status: false,
            message: 'Ошибка валидации id задачи'
        }
    }
    const url = `/api/lists/${listId}/todos/${id}`
    const headers = { 'Authorization': `Bearer ${store.getToken()}` }
    let delRes
    try {
        delRes = await connInstanse.delete(url, { headers })
    }
    catch(err) {
        return {
            status: false,
            message: err.response.data.message
        }
    }
    return {
       status: delRes.data.status
    }
}