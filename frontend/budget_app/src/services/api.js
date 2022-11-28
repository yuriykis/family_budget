import * as Axios from 'axios'
import { getAccessToken } from '../store/api_auth'

var apiHost = 'https://future_budget_app.com'

if (process.env.NODE_ENV === 'development') {
    apiHost = 'http://localhost:8000'
}

const authenticationHeader = () => {
    return {
        Authorization: `Bearer ${getAccessToken()}`
    }
}

const api = Axios.create({
    baseURL: apiHost,
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
    }
})

export async function register(user) {
    return api.post('/app/register/', user)
}

export async function loginToTheApplication(username, password) {
    return api.post('api/token/', { username: username, password: password })
}

export async function getCategories() {
    return api.get('/app/category/', { headers: authenticationHeader() })
}

export async function getTransactions() {
    return api.get('/app/transaction/', { headers: authenticationHeader() })
}

export async function getTransactionsByCategory(categoryId) {
    return api.get(`/app/transaction/?category=${categoryId}`, { headers: authenticationHeader() })
}

export async function getBudgets() {
    return api.get('/app/budget/', { headers: authenticationHeader() })
}


