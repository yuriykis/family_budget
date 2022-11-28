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

export async function getTransactions(budgetId) {
    return api.get(`/app/budget/${budgetId}/transaction/`, { headers: authenticationHeader() })
}

export async function createTransaction(transaction, budgetId) {
    return api.post(`/app/budget/${budgetId}/transaction/`, transaction, { headers: authenticationHeader() })
}

export async function getTransactionsByCategory(categoryId) {
    return api.get(`/app/transaction/?category=${categoryId}`, { headers: authenticationHeader() })
}

export async function getBudgets() {
    return api.get('/app/budget/', { headers: authenticationHeader() })
}

export async function createBudget(budget) {
    return api.post('/app/budget/', budget, { headers: authenticationHeader() })
}

export async function editBudget(budget) {
    return api.put(`/app/budget/${budget.id}/`, budget, { headers: authenticationHeader() })
}

export async function deleteBudget(budgetId) {
    return api.delete(`/app/budget/${budgetId}/`, { headers: authenticationHeader() })
}

export async function createCategory(category) {
    return api.post('/app/category/', category, { headers: authenticationHeader() })
}

export async function getCategories() {
    return api.get('/app/category/', { headers: authenticationHeader() })
}

export async function getMyUser() {
    return api.get('/app/user/', { headers: authenticationHeader() })
}

export async function getUser(userId) {
    return api.get(`/app/user/${userId}/`, { headers: authenticationHeader() })
}

export async function getAllUsers() {
    return api.get('/app/users/', { headers: authenticationHeader() })
}

export async function shareBudget(budgetId, userId) {
    return api.post(`/app/budget/${budgetId}/share/`, { user_id: userId }, { headers: authenticationHeader() })
}