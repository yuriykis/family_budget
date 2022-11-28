import Vue from 'vue'
import Vuex from 'vuex'
import auth from './modules/auth'
import budgets from './modules/budgets'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        auth,
        budgets,
    },
    strict: debug
})