import { writable } from 'svelte/store'

export const listsStore = writable([])
export const modalCreateListVisible = writable(false)
export const modalEditListVisible = writable(false)
export const modalDescribeListVisible = writable(false)
export const currentListStore = writable({id: 0, title: '', description: ''})
export const maxLengthTitle = 50
export const maxLengthDescription = 200