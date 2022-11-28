<template>
    <v-container>
        <v-overlay :value="loading">
            <Loader />
        </v-overlay>
        <v-dialog v-model="shareBudgetDialog" max-width="400">
            <ShareBudgetDialog @share-budget="(selectedUser) => shareBudget(selectedUser)"
                @on-close="shareBudgetDialog = false" />
        </v-dialog>
        <v-row>
            <v-col cols="12">
                <div id="my_budget">My Budgets</div>
            </v-col>
        </v-row>
        <v-row justify="end" class="ma-0">
            <v-dialog v-model="dialog" width="500">
                <template v-slot:activator="{ on, attrs }">
                    <v-btn dark class="white--text" @click="addBudgetDialog" v-bind="attrs" v-on="on">
                        Add New Budget
                    </v-btn>
                </template>
                <BudgetDialog @on-close="dialog = false"
                    @budget-action="(budget, action_title) => executeBudgetAction(budget, action_title)"
                    :title="budgetTitle" />
            </v-dialog>
        </v-row>
        <v-row>
            <v-col cols="12">
                <BudgetsList @editBudget="(budget) => editBudgetDialog(budget)"
                    @deleteBudget="(budget) => deleteBudget(budget)"
                    @shareBudget="(budget) => openShareBudgetDialog(budget)" />
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import BudgetsList from './BudgetsList.vue';
import BudgetDialog from './BudgetDialog.vue';
import Loader from '@/components/Loader'
import { mapActions } from 'vuex';
import { shareBudget } from '../services/api';
import ShareBudgetDialog from './ShareBudgetDialog.vue';
export default {
    name: "Home",
    data: () => ({
        dialog: false,
        budgetTitle: "",
        loading: false,
        selectedBudget: {},
        shareBudgetDialog: false,
    }),
    components: {
        BudgetsList, BudgetDialog, Loader,
        ShareBudgetDialog
    },
    methods: {
        ...mapActions('budgets', ['addBudgetAction', 'editBudgetAction', 'deleteBudgetAction']),
        setBudgetTitle(title) {
            this.budgetTitle = title;
        },
        addBudgetDialog() {
            this.setBudgetTitle("Add New Budget");
            this.dialog = true;
        },
        editBudgetDialog(budget) {
            this.selectedBudget = budget;
            this.setBudgetTitle("Edit Budget");
            this.dialog = true;
        },

        async deleteBudget(budget) {
            this.dialog = false;
            this.loading = true;
            try {
                await this.deleteBudgetAction(budget.width);
            } catch (error) {
                console.log(error);
            }
            this.loading = false;
        },
        openShareBudgetDialog(budget) {
            this.selectedBudget = budget;
            this.shareBudgetDialog = true;
        },
        async shareBudget(selectedUser) {
            this.shareBudgetDialog = false;
            this.loading = true;
            try {
                await shareBudget(this.selectedBudget.id, selectedUser.user.id);
            } catch (error) {
                console.log(error);
            }
            this.loading = false;
        },
        async executeBudgetAction(budget, action_title) {
            this.loading = true;
            switch (action_title) {
                case "Add New Budget":
                    await this.addBudgetAction(budget);
                    this.loading = false;
                    this.dialog = false;
                    break;
                case "Edit Budget":
                    budget.id = this.selectedBudget.id;
                    this.editBudgetAction(budget);
                    this.loading = false;
                    this.dialog = false;
                    break;
                default:
                    this.loading = false;
                    this.dialog = false;
            }
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