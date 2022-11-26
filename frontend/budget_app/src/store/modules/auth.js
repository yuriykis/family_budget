import axios from 'axios'

const state = {
    isAuthorized: true,
}

const getters = {
    isAuthenticated: state => state.isAuthorized,
}

const actions = {
    async login({ commit }, user) {
        try {
            const response = await axios.post('/api/auth/login', user)
            commit('setAuth', response.data)
        } catch (error) {
            console.log(error)
        }
    }
}

const mutations = {
    setAuth(state, data) {
        state.isAuthorized = true
        localStorage.setItem('user', JSON.stringify(data))
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}


