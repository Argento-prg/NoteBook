import { connInstanse } from '../../requestConfig'
import { store } from '../../storeApp'
import { listValidate, idValidate } from './listValidate'

export async function getAllLists() {
    const url = '/api/lists/'
    const headers = { 'Authorization': `Bearer ${store.getToken()}` }
    let listsRes
    try {
        listsRes = await connInstanse.get(url, { headers })
    }
    catch(err) {
        return {
            status: false,
            message: err.response.data.message
        }
    }
    let listsData = []
    if (listsRes.data.data !== null) {
        listsData = listsRes.data.data
    }
    return { 
        status: true, 
        data: listsData
    }
}

export async function getListById(id) {
    if (!idValidate(id)) {
        return {
            status: false,
            message: ''
        }
    }
    const url = `/api/lists/${id}`
    const headers = { 'Authorization': `Bearer ${store.getToken()}` }
    let listRes
    try{
        listRes = await connInstanse.get(url, { headers })
    }
    catch(err) {
        return {
            status: false,
            message: err.response.data.message
        }
    }
    return {
        status: true,
        data: listRes.data
    }
}

export async function createList(title, description) {
    if (!listValidate(title, description)) {
        return {
            status: false,
            message: 'Некорректные данные'
        }
    }
    const url = '/api/lists/'
    const headers = { 'Authorization': `Bearer ${store.getToken()}` }
    const data = {
        title,
        description
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
        status: true,
        id: createRes.data.id
    }
}

export async function updateList(id, title, description) {
    if (!listValidate(title, description)) {
        return {
            status: false,
            message: 'Некорректные данные'
        }
    }
    if (!idValidate(id)) {
        return {
            status: false,
            message: ''
        }
    }
    const url = `/api/lists/${id}`
    const headers = { 'Authorization': `Bearer ${store.getToken()}` }
    const data = {
        title,
        description
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

export async function deleteList(id) {
    if (!idValidate(id)) {
        return {
            status: false,
            message: ''
        }
    }
    const url = `/api/lists/${id}`
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