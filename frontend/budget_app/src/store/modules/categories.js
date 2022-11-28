import { createCategory, getCategories } from '../../services/api'

const state = {
    categories: [],
}

const getters = {
    categories: state => state.categories,
}

const actions = {
    async getCategoriesAction({ commit }) {
        try {
            const response = await getCategories()
            commit('setCategories', await response.data)
        } catch (error) {
            console.log(error)
        }
    },
    async addCategoryAction({ dispatch }, category) {
        try {
            await createCategory(category)
            dispatch('getCategoriesAction')
        } catch (error) {
            console.log(error)
        }
    }
}

const mutations = {
    setCategories(state, data) {
        state.categories = data
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
