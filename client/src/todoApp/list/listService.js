import { listsStore, currentListStore } from './listStore'
import { choosedListStore } from '../todoAppStore'
import { getAllLists, createList, getListById, updateList, deleteList } from './listRequests'

export async function getAllListsService() {
    const response = await getAllLists()
    if (!response.status) {
        alert(response.message)
        return false
    }
    listsStore.update(() => response.data)
    return true
}
export async function createListService(title, description) {
    const response = await createList(title, description)
    if (!response.status) {
        alert(response.message)
        return false
    }
    const updateCurrentListofListsStatus = await getAllListsService()
    return updateCurrentListofListsStatus
}
export async function getListByIdService(id) {
    const response = await getListById(id)
    if (!response.status) {
        if (response.message !== '') {
            alert(response.message)
        }
        return false
    }
    currentListStore.update(() => {
        return {
            id: response.data.id,
            title: response.data.title,
            description: response.data.description
        }
    })
    return true
}
export async function editListService(title, description) {
    let id
    currentListStore.subscribe((value) => {id = value.id})
    const response = await updateList(id, title, description)
    if (!response.status) {
        alert(response.message)
        return false
    }
    const updateCurrent = await getListByIdService(id)
    const updateAllCurrentLists = await getAllListsService()
    return updateCurrent && updateAllCurrentLists
}
export async function delListByIdService() {
    let id
    currentListStore.subscribe((value) => {id = value.id})
    const response = await deleteList(id)
    if (!response.status) {
        alert(response.message)
        return false
    }
    choosedListStore.update(() => -1)
    currentListStore.update(() => {
        return {id: 0, title: '', description: ''}
    })
    const updateAllCurrentLists = await getAllListsService()
    return updateAllCurrentLists
}