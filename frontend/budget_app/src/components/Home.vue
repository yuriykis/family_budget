<template>
    <v-container>
        <v-row>
            <v-col cols="12">
                <div id="my_budget">My Budgets</div>
            </v-col>
        </v-row>
        <v-row justify="end" class="ma-0">
            <v-dialog v-model="dialog" width="500">
                <template v-slot:activator="{ on, attrs }">
                    <v-btn dark class="white--text" @click="addBudget" v-bind="attrs" v-on="on">
                        Add New Budget
                    </v-btn>
                </template>
                <BudgetProps @on-close="closeDialog" :title="budgetTitle" />
            </v-dialog>
        </v-row>
        <v-row>
            <v-col cols="12">
                <BudgetsList @editBudget="editBudgetDialog" @deleteBudget="deleteBudget" />
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import BudgetsList from './BudgetsList.vue';
import BudgetProps from './BudgetProps.vue';
export default {
    name: "Home",
    data: () => ({
        dialog: false,
        budgetTitle: ""
    }),
    components: { BudgetsList, BudgetProps },
    methods: {
        setBudgetTitle(title) {
            this.budgetTitle = title;
        },
        addBudget() {
            this.setBudgetTitle("Add New Budget");
            this.dialog = true;
        },
        closeDialog() {
            this.dialog = false;
        },
        editBudgetDialog() {
            this.setBudgetTitle("Edit Budget");
            this.dialog = true;
        },
        deleteBudget() {
            this.dialog = false;
            console.log("delete budget");
        }
    },
    computed: {
        budgetTitle() {
            return this.budgetTitle;
        }
    }
}
</script>

<style scoped>
#my_budget {
    font-size: 30px;
    font-weight: bold;
    margin-top: 20px;
    margin-bottom: 20px;
    font-family: 'Roboto', sans-serif;
}
</style>