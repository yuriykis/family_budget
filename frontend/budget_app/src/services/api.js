import * as Axios from 'axios'
import { getAccessToken } from '../store/api_auth'

var apiHost = 'http://localhost:52000'

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
    return api.post('/api/register/', user)
}

export async function loginToTheApplication(username, password) {
    return api.post('api/token/', { username: username, password: password })
}

export async function getTransactions(budgetId) {
    return api.get(`/api/budget/${budgetId}/transaction/`, { headers: authenticationHeader() })
}

export async function createTransaction(transaction, budgetId) {
    return api.post(`/api/budget/${budgetId}/transaction/`, transaction, { headers: authenticationHeader() })
}

export async function getTransactionsByCategory(categoryId) {
    return api.get(`/api/transaction/?category=${categoryId}`, { headers: authenticationHeader() })
}

export async function getBudgets() {
    return api.get('/api/budget/', { headers: authenticationHeader() })
}

export async function createBudget(budget) {
    return api.post('/api/budget/', budget, { headers: authenticationHeader() })
}

export async function editBudget(budget) {
    return api.put(`/api/budget/${budget.id}/`, budget, { headers: authenticationHeader() })
}

export async function deleteBudget(budgetId) {
    return api.delete(`/api/budget/${budgetId}/`, { headers: authenticationHeader() })
}

export async function createCategory(category) {
    return api.post('/api/category/', category, { headers: authenticationHeader() })
}

export async function getCategories() {
    return api.get('/api/category/', { headers: authenticationHeader() })
}

export async function getMyUser() {
    return api.get('/api/user/', { headers: authenticationHeader() })
}

export async function getUser(userId) {
    return api.get(`/api/user/${userId}/`, { headers: authenticationHeader() })
}

export async function getAllUsers() {
    return api.get('/api/users/', { headers: authenticationHeader() })
}

export async function shareBudget(budgetId, userId) {
    return api.post(`/api/budget/${budgetId}/share/`, { user_id: userId }, { headers: authenticationHeader() })
}