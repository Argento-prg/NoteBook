import { writable } from 'svelte/store'

export const choosedTodoStore = writable(-1)
export const todosStore = writable([])
export const modalCreateTodoVisible = writable(false)
export const modalEditTodoVisible = writable(false)
export const modalDescribeTodoVisible = writable(false)
export const currentTodoStore = writable({
    id: 0,
    title: '',
    description: '',
    importance: false,
    done: false
})
export const maxLengthTitle = 100
export const maxLengthDescription = 300