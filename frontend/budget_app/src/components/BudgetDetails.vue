<template>
    <v-container>
        <v-overlay :value="loading">
            <Loader />
        </v-overlay>
        <v-row>
            <v-col cols="12">
                <div id="budget_details">{{ budget.name }}</div>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="12">
                <span>{{ budget.description }}</span>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="4">
                <v-card>
                    <v-card-title>
                        <span class="headline">Amount</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-row>
                                <v-col cols="12">
                                    <h1 style="color: black; font-size: 40px;">${{ budget.amount }}</h1>
                                </v-col>
                            </v-row>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-col>
            <v-col cols="4">
                <v-card>
                    <v-card-title>
                        <span class="headline">Amount left</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-row>
                                <v-col cols="12">
                                    <h1 style="color: black; font-size: 40px;">${{ amountLeft }}</h1>
                                </v-col>
                            </v-row>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-col>
            <v-col cols="4">
                <v-card>
                    <v-card-title>
                        <span class="headline">Total transactions</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-row>
                                <v-col cols="12">
                                    <h1 style="color: black; font-size: 40px;">{{ totalTransactions }}</h1>
                                </v-col>
                            </v-row>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
        <v-row justify="end" class="mt-5">
            <v-dialog v-model="dialog" width="500">
                <template v-slot:activator="{ on, attrs }">
                    <v-btn dark class="white--text" v-bind="attrs" v-on="on" :disabled="!budget.ownership">
                        New transaction
                    </v-btn>
                </template>
                <TransactionDialog @on-close="closeTransactionDialog"
                    @transaction-action="(transaction) => createTransaction(transaction)" :categories="categories" />
            </v-dialog>
        </v-row>
        <v-row>
            <v-col cols="12">
                <TransationsList :budget="budget" ref="transactionsList" />
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import TransationsList from './TransationsList.vue';
import TransactionDialog from './TransactionDialog.vue';
import Loader from '@/components/Loader'
import { mapGetters, mapActions, mapMutations } from 'vuex';
export default {
    name: "BudgetDetails",
    data: () => ({
        budget: {
        },
        dialog: false,
        loading: false,
    }),
    components: {
        TransationsList,
        TransactionDialog,
        Loader
    },
    computed: {
        ...mapGetters('budgets', ['budgets']),
        ...mapGetters('categories', ['categories']),
        ...mapGetters('transactions', ['amountLeft', 'totalTransactions']),
    },
    methods: {
        ...mapActions('transactions', ['addTransactionAction']),
        ...mapActions('categories', ['getCategoriesAction']),
        ...mapMutations('transactions', ['setAmountLeft']),
        openTransactionDialog() {
            this.dialog = true;
        },
        closeTransactionDialog() {
            this.dialog = false;
        },
        async createTransaction(transaction) {
            this.dialog = false;
            this.loading = true;
            const budgetId = this.budget.id
            await this.addTransactionAction({ transaction, budgetId });
            this.loading = false;
        }
    },
    async created() {
        const budgetId = this.$route.params.budgetId;
        this.budget = this.budgets.find(b => b.id == budgetId);
        this.setAmountLeft(this.budget.amount_left);
        await this.getCategoriesAction();
    },
};
</script>

<style scoped>
#budget_details {
    font-size: 30px;
    font-weight: bold;
    margin-top: 20px;
    margin-bottom: 20px;
    font-family: 'Roboto', sans-serif;
}
</style>