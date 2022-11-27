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
                                    <v-list-item-title>{{ transaction.name }}</v-list-item-title>
                                    <v-list-item-subtitle>{{ transaction.description }}</v-list-item-subtitle>
                                </v-list-item-content>
                                <v-list-item-action v-if="transaction.category">
                                    <v-chip>
                                        {{ transaction.category }}
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
export default {
    name: 'TransationsList',
    data: () => ({
        tab: null,
        items: [
            {
                tab: 'Income', transactions: [
                    {
                        id: 1,
                        name: 'Income 1',
                        description: 'Description 1',
                        category: 'category 1',
                    },
                    {
                        id: 2,
                        name: 'Income 2',
                        description: 'Description 2',
                        category: 'category 2',
                    },
                    {
                        id: 3,
                        name: 'Income 3',
                        description: 'Description 3',
                        category: 'category 3',
                    },
                ]
            },
            {
                tab: 'Expenses', transactions: [
                    {
                        id: 1,
                        name: 'Expense 1',
                        description: 'Description 1',
                        category: 'category 1',
                    },
                    {
                        id: 2,
                        name: 'Expense 2',
                        description: 'Description 2'
                    },
                    {
                        id: 3,
                        name: 'Expense 3',
                        description: 'Description 3'
                    },
                ]
            }
        ],
    }),
}
</script>