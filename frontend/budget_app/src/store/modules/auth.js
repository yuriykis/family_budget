import { loginToTheApplication } from '../../services/api'
import * as auth from '../api_auth'

const state = {
    isAuthorized: auth.ifValidToken(),
}

const getters = {
    isAuthenticated: state => state.isAuthorized,
}

const actions = {
    async loginUser({ commit }, userData) {
        try {
            const response = await loginToTheApplication(userData.username, userData.password)
            commit('setAuth', await response.data)
        } catch (error) {
            console.log(error)
        }
    },
    logoutUser({ commit }) {
        commit('removeAuth')
    }
}

const mutations = {
    setAuth(state, data) {
        state.isAuthorized = true
        auth.setLocalStorage(data.access, data.refresh)
    },
    removeAuth(state) {
        state.isAuthorized = false
        auth.removeLocalStorage()
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}


