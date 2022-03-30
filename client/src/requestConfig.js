import axios from 'axios'

export const connInstanse = axios.create(
    {
        baseURL: 'http://localhost:8000'
    }
)