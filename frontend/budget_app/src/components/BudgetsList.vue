<template>
    <v-card flat>
        <v-dialog v-model="deleteBudgetDialog" persistent max-width="400">
            <v-card dark>
                <v-card-title class="headline">
                    Are you sure you want to delete this budget?
                </v-card-title>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn text @click="deleteBudget()">
                        Yes
                    </v-btn>
                    <v-btn text @click="closeDeleteBudgetDialog">
                        No
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-list>
            <v-list-item-group>
                <v-list-item v-for="budget in budgets()" :key="budget.id" @click="selectBudget(budget)">
                    <v-list-item-content>
                        <v-list-item-title>{{ budget.name }}</v-list-item-title>
                        <v-list-item-subtitle>{{ budget.description }}</v-list-item-subtitle>
                    </v-list-item-content>
                    <v-list-item-content>
                        <v-list-item-title>Amount</v-list-item-title>
                        <v-list-item-subtitle>{{ budget.amount }}</v-list-item-subtitle>
                    </v-list-item-content>
                    <v-list-item-action>
                        <v-chip v-if="budget.ownership" color="green">
                            Owner
                        </v-chip>
                        <v-chip v-else color="yellow">
                            Shared
                        </v-chip>
                    </v-list-item-action>
                    <v-list-item-action>
                        <v-menu right offset-x rounded>
                            <template v-slot:activator="{ on, attrs }">
                                <v-btn icon v-bind="attrs" v-on="on">
                                    <v-icon>mdi-dots-vertical</v-icon>
                                </v-btn>
                            </template>

                            <v-list dense height="150px">
                                <v-list-item v-for="(budgetAction, index) in budgetActions" :key="index">
                                    <v-list-item-title>
                                        <v-btn @click="executeBudgetAction(budget, budgetAction)" plain
                                            :disabled="!budget.ownership">
                                            {{ budgetAction.title }}
                                        </v-btn>
                                    </v-list-item-title>
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
import { mapGetters, mapActions } from 'vuex'
export default {
    name: 'BudgetsList',
    data: () => ({
        deleteBudgetDialog: false,
        budgetToDelete: null,
        budgetActions: [
            { title: 'Edit', icon: 'mdi-pencil' },
            { title: 'Share', icon: 'mdi-share' },
            { title: 'Delete', icon: 'mdi-delete' },
        ],

    }),
    async created() {
        await this.getBudgetsAction()
    },
    methods: {
        ...mapGetters('budgets', ['budgets']),
        ...mapActions('budgets', ['getBudgetsAction']),
        selectBudget(budget) {
            this.$router.push({ path: `/budget/${budget.id}` });
        },
        executeBudgetAction(budget, budgetAction) {
            if (budgetAction.title === 'Edit') {
                this.$emit('editBudget', budget)
            } else if (budgetAction.title === 'Delete') {
                this.budgetToDelete = budget
                this.openDeleteBudgetDialog()
            } else if (budgetAction.title === 'Share') {
                this.$emit('shareBudget', budget)
            }
        },
        openDeleteBudgetDialog() {
            this.deleteBudgetDialog = true
        },
        deleteBudget() {
            this.$emit('deleteBudget', this.budgetToDelete)
            this.deleteBudgetDialog = false
        },
        closeDeleteBudgetDialog() {
            this.deleteBudgetDialog = false
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


