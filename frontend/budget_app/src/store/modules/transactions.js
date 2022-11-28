import { createTransaction, getTransactions } from '../../services/api'

const state = {
    transactions: [],
    amountLeft: 0,
    totalTransactions: 0,
}

const getters = {
    transactions: state => state.transactions,
    amountLeft: state => state.amountLeft,
    totalTransactions: state => state.transactions.length,
}

const actions = {
    async getTransactionsAction({ commit }, budgetId) {
        try {
            const response = await getTransactions(budgetId)
            commit('setTransactions', await response.data)
        } catch (error) {
            console.log(error)
        }
    },
    async addTransactionAction({ commit, dispatch }, transaction_obj) {
        try {
            const res = await createTransaction(transaction_obj.transaction, transaction_obj.budgetId)
            commit('setAmountLeft', res.data.amount_left)
            commit('setTotalTransactions', res.data.total_transactions)
            dispatch('getTransactionsAction', transaction_obj.budgetId)
        } catch (error) {
            console.log(error)
        }
    }
}

const mutations = {
    setTransactions(state, data) {
        state.transactions = data
    },
    setAmountLeft(state, data) {
        state.amountLeft = data
    },
    setTotalTransactions(state, data) {
        state.totalTransactions = data
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
