import { getBudgets, createBudget, editBudget, deleteBudget } from '../../services/api'

const state = {
    budgets: [
    ],
}

const getters = {
    budgets: state => state.budgets,
}

const actions = {
    async getBudgetsAction({ commit }) {
        try {
            const response = await getBudgets()
            commit('setBudgets', await response.data)
        } catch (error) {
            console.log(error)
        }
    },
    async addBudgetAction({ dispatch }, budget) {
        try {
            await createBudget(budget)
            dispatch('getBudgetsAction')
        } catch (error) {
            console.log(error)
        }
    },
    async editBudgetAction({ dispatch }, budget) {
        try {
            await editBudget(budget)
            dispatch('getBudgetsAction')
        } catch (error) {
            console.log(error)
        }
    },
    async deleteBudgetAction({ dispatch }, budgetId) {
        try {
            await deleteBudget(budgetId)
            dispatch('getBudgetsAction')
        } catch (error) {
            console.log(error)
        }
    }
}

const mutations = {
    setBudgets(state, data) {
        state.budgets = data
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}