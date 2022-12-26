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
    },
})

export async function register(user) {
    return api.post('/api/register', user)
}

export async function loginToTheApplication(email, password) {
    return api.post('/api/login', { email: email, password: password })
}

export async function getTransactions(budgetId) {
    return api.get(`/api/auth/budget/${budgetId}/transaction`, { headers: authenticationHeader() })
}

export async function createTransaction(transaction, budgetId) {
    return api.post(`/api/auth/budget/${budgetId}/transaction`, transaction, { headers: authenticationHeader() })
}

export async function getTransactionsByCategory(categoryId) {
    return api.get(`/api/auth/transaction/?category=${categoryId}`, { headers: authenticationHeader() })
}

export async function getBudgets() {
    return api.get('/api/auth/budget', { headers: authenticationHeader() })
}

export async function createBudget(budget) {
    return api.post('/api/auth/budget', budget, { headers: authenticationHeader() })
}

export async function editBudget(budget) {
    return api.put(`/api/auth/budget/${budget.id}`, budget, { headers: authenticationHeader() })
}

export async function deleteBudget(budgetId) {
    return api.delete(`/api/auth/budget/${budgetId}`, { headers: authenticationHeader() })
}

export async function createCategory(category) {
    return api.post('/api/auth/category', category, { headers: authenticationHeader() })
}

export async function getCategories() {
    return api.get('/api/auth/category', { headers: authenticationHeader() })
}

export async function getMyUser() {
    return api.get('/api/auth/user', { headers: authenticationHeader() })
}

export async function getUser(userId) {
    return api.get(`/api/auth/user/${userId}`, { headers: authenticationHeader() })
}

export async function getAllUsers() {
    return api.get('/api/auth/users', { headers: authenticationHeader() })
}

export async function shareBudget(budgetId, userId) {
    return api.post(`/api/auth/budget/${budgetId}/share`, { user_id: userId }, { headers: authenticationHeader() })
}