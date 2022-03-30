import { writable } from 'svelte/store'

export const choiceUserStore = writable(false)
export const maxLengthName = 50
export const maxLengthLogin = 50
export const minLengthPassword = 6
export const maxLengthPassword = 50
