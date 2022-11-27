<template>
    <v-container>
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
                        <span class="headline">Total budget value</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-row>
                                <v-col cols="4">
                                    <h1 style="color: black; font-size: 40px;">${{ budget.totalValue }}</h1>
                                </v-col>
                            </v-row>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-col>
            <v-col cols="4">
                <v-card>
                    <v-card-title>
                        <span class="headline">Budget left</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-row>
                                <v-col cols="4">
                                    <h1 style="color: black; font-size: 40px;">${{ budget.valueLeft }}</h1>
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
                                <v-col cols="4">
                                    <h1 style="color: black; font-size: 40px;">{{ budget.totalTransactions }}</h1>
                                </v-col>
                            </v-row>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="12">
                <TransationsList :budget="budget" />
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import TransationsList from './TransationsList.vue';

export default {
    name: "BudgetDetails",
    data: () => ({
        budget: {
            name: "Budget 1",
            description: "klsdcnlwvmnerlkmvlwemc;lkwem",
            totalValue: 5000,
            valueLeft: 5000,
            totalTransactions: 0
        }
    }),
    components: {
        TransationsList
    },
    methods: {

    },
    computed: {

    }
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