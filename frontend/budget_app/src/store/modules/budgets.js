import { getBudgets } from '../../services/api'

const state = {
    budgets: [
    ],
}

const getters = {
    budgets: state => state.budgets,
}

const actions = {
    async getBudgets({ commit }) {
        try {
            const response = await getBudgets()
            commit('setBudgets', await response.data)
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