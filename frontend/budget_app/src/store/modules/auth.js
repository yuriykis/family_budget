import { loginToTheApplication, getMyUser } from '../../services/api'
import * as auth from '../api_auth'

const state = {
    isAuthorized: auth.ifValidToken(),
    userData: {},
}

const getters = {
    isAuthenticated: state => state.isAuthorized,
    userData: state => state.userData,
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
    },
    async getMyUserAction({ commit }) {
        try {
            const response = await getMyUser()
            commit('setUserData', await response.data)
        } catch (error) {
            console.log(error)
        }
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
    },
    setUserData(state, data) {
        state.userData = data
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}


