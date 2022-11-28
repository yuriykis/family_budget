import { getAllUsers } from '../../services/api'

const state = {
    users: [],
}

const getters = {
    users: state => state.users,
}

const actions = {
    async getAllUsersAction({ commit }) {
        try {
            const response = await getAllUsers()
            commit('setUsers', await response.data)
        } catch (error) {
            console.log(error)
        }
    }
}

const mutations = {
    setUsers(state, data) {
        state.users = data
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}