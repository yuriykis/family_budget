<template>
    <v-card flat>
        <v-tabs v-model="tab" color="#363636">
            <v-tab v-for="item in items" :key="item.tab">
                {{ item.tab }}
            </v-tab>
        </v-tabs>

        <v-tabs-items v-model="tab">
            <v-tab-item v-for="item in items" :key="item.tab">
                <v-card flat>
                    <v-list>
                        <v-list-item-group>
                            <v-list-item v-for="transaction in item.transactions" :key="transaction.id"
                                @click="selectBudget(budget)">
                                <v-list-item-content>
                                    <v-list-item-title>{{ transaction.title }}</v-list-item-title>
                                    <v-list-item-subtitle>{{ transaction.datetime }}</v-list-item-subtitle>
                                </v-list-item-content>
                                <v-list-item-content>
                                    <v-list-item-title>Amount</v-list-item-title>
                                    <v-list-item-subtitle>{{ transaction.amount }}</v-list-item-subtitle>
                                </v-list-item-content>
                                <v-list-item-action v-if="transaction.category">
                                    <v-chip>
                                        {{ transaction.category.name }}
                                    </v-chip>
                                </v-list-item-action>
                            </v-list-item>
                        </v-list-item-group>
                    </v-list>
                </v-card>
            </v-tab-item>
        </v-tabs-items>
    </v-card>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
export default {
    name: 'TransationsList',
    props: {
        budget: {
            type: Object,
            required: true
        }
    },
    data: () => ({
        tab: null,
        items: [
            {
                tab: 'Income', transactions: []
            },
            {
                tab: 'Expenses', transactions: []
            }
        ]
    }),
    computed: {
        ...mapGetters('transactions', ['transactions'])
    },
    async created() {
        await this.getTransactionsAction(this.budget.id)
        this.items[0].transactions = this.transactions.filter(transaction => transaction.type === 'INCOME')
        this.items[1].transactions = this.transactions.filter(transaction => transaction.type === 'EXPENSE')
    },
    methods: {
        ...mapActions('transactions', ['getTransactionsAction']),
        selectTransaction(transaction) {
            this.selectTransaction(transaction)
        }
    },
    watch: {
        transactions() {
            this.items[0].transactions = this.transactions.filter(transaction => transaction.type === 'INCOME')
            this.items[1].transactions = this.transactions.filter(transaction => transaction.type === 'EXPENSE')
        }
    }
}
</script>