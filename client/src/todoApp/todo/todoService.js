import { 
    getAllTodos,
    getTodoById,
    createTodo,
    updateTodo,
    deleteTodo
} from './todoRequests'
import { todosStore, choosedTodoStore , currentTodoStore} from './todoStore'
import { choosedListStore } from '../todoAppStore'

let listId
choosedListStore.subscribe(value => listId = value)
let todoId
choosedTodoStore.subscribe(value => todoId = value)

export async function getAllTodosService() {
    const response = await getAllTodos(listId)
    if(!response.status) {
        alert(response.message)
        return false
    }
    todosStore.update(() => response.data)
    return true
}
export async function getTodoByIdService() {
    const response = await getTodoById(listId, todoId)
    if(!response.status) {
        alert(response.message)
        return false
    }
    currentTodoStore.update(() => {
        return {
            id: response.data.id,
            title: response.data.title,
            description: response.data.description,
            importance: response.data.importance,
            done: response.data.done
        }
    })
    return true
}
export async function createTodoService(title, description, importance) {
    const response = await createTodo(listId, title, description, importance)
    if(!response.status) {
        alert(response.message)
        return false
    }
    const updateAllTodos = await getAllTodosService()
    return updateAllTodos
}
export async function updateTodoService(title, description, importance, done) {
    const response = await updateTodo(listId, todoId, title, description, importance, done)
    if(!response.status) {
        alert(response.message)
        return false
    }
    const updateAllTodos = await getAllTodosService()
    return updateAllTodos
}
export async function updateImportanceTodoService(importance) {
    const response = await updateTodo(listId, todoId, -1, -1, importance, -1)
    if(!response.status) {
        alert(response.message)
        return false
    }
    const updateAllTodos = await getAllTodosService()
    return updateAllTodos
}
export async function updateDoneTodoService(done) {
    const response = await updateTodo(listId, todoId, -1, -1, -1, done)
    if(!response.status) {
        alert(response.message)
        return false
    }
    const updateAllTodos = await getAllTodosService()
    return updateAllTodos
}
export async function deleteTodoService() {
    const response = await deleteTodo(listId, todoId)
    if(!response.status) {
        alert(response.message)
        return false
    }
    const updateAllTodos = await getAllTodosService()
    return updateAllTodos
}