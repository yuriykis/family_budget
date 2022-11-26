<template>
    <v-card flat>
        <v-list>
            <v-list-item-group>
                <v-list-item v-for="budget in budgets" :key="budget.id" @click="selectBudget(budget)">
                    <v-list-item-content>
                        <v-list-item-title>{{ budget.name }}</v-list-item-title>
                        <v-list-item-subtitle>{{ budget.description }}</v-list-item-subtitle>
                    </v-list-item-content>
                    <v-list-item-action v-if="budget.label">
                        <v-chip :color=budget.labelColor>
                            {{ budget.label }}
                        </v-chip>
                    </v-list-item-action>
                    <v-list-item-action>
                        <v-menu right offset-x=100 rounded>
                            <template v-slot:activator="{ on, attrs }">
                                <v-btn icon v-bind="attrs" v-on="on">
                                    <v-icon>mdi-dots-vertical</v-icon>
                                </v-btn>
                            </template>

                            <v-list dense height="150px">
                                <v-list-item v-for="(budgetAction, index) in budgetActions" :key="index">
                                    <v-list-item-title>{{ budgetAction.title }}</v-list-item-title>
                                </v-list-item>
                            </v-list>
                        </v-menu>
                    </v-list-item-action>
                </v-list-item>
            </v-list-item-group>
        </v-list>
    </v-card>
</template>

<script>
export default {
    name: 'BudgetsList',
    data: () => ({
        budgetActions: [
            { title: 'Edit', icon: 'mdi-pencil' },
            { title: 'Delete', icon: 'mdi-delete' },
            { title: 'Share', icon: 'mdi-share' },
        ],
        budgets: [
            {
                id: 1,
                name: 'Budget 1',
                description: 'Description 1'
            },
            {
                id: 2,
                name: 'Budget 2',
                description: 'Description 2'
            },
            {
                id: 3,
                name: 'Budget 3',
                description: 'Description 3',
                label: 'Label 3',
                labelColor: 'green'
            },
            {
                id: 4,
                name: 'Budget 4',
                description: 'Description 4'
            },
            {
                id: 5,
                name: 'Budget 5',
                description: 'Description 5',
                label: 'Label 5',
                labelColor: 'red'
            },
            {
                id: 6,
                name: 'Budget 6',
                description: 'Description 6'
            },
            {
                id: 7,
                name: 'Budget 7',
                description: 'Description 7'
            },
            {
                id: 8,
                name: 'Budget 8',
                description: 'Description 8'
            },
            {
                id: 9,
                name: 'Budget 9',
                description: 'Description 9'
            },
            {
                id: 10,
                name: 'Budget 10',
                description: 'Description 10'
            },
            {
                id: 11,
                name: 'Budget 11',
                description: 'Description 11'
            },
            {
                id: 12,
                name: 'Budget 12',
                description: 'Description 12'
            },
            {
                id: 13,
                name: 'Budget 13',
                description: 'Description 13'
            },
            {
                id: 14,
                name: 'Budget 14',
                description: 'Description 14'
            },
            {
                id: 15,
                name: 'Budget 15',
                description: 'Description 15'
            },
            {
                id: 16,
                name: 'Budget 16',
                description: 'Description 16'
            },
            {
                id: 17,
                name: 'Budget 17',
                description: 'Description 17'
            },
            {
                id: 18,
                name: 'Budget 18',
                description: 'Description 18'
            },
            {
                id: 19,
                name: 'Budget 19',
                description: 'Description 19'
            },
        ]
    }),
    methods: {
        selectBudget(budget) {
            this.$emit('selectBudget', budget)
        }
    }
}
</script>

<style scoped>
.v-list {
    height: 600px;
    overflow-y: auto
}
</style>


