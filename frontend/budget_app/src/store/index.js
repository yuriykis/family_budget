import Vue from 'vue'
import Vuex from 'vuex'
import auth from './modules/auth'
import budgets from './modules/budgets'
import transactions from './modules/transactions'
import categories from './modules/categories'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        auth,
        budgets,
        transactions,
        categories
    },
    strict: debug
})