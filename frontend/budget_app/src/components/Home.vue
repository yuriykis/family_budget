<template>
    <v-container>
        <v-row>
            <div id="my_budget">My Budgets</div>
        </v-row>
        <v-row>
            <v-col cols="12">
                <BudgetsList />
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import BudgetsList from './BudgetsList.vue';
export default {
    name: "Home",
    data: () => ({}),
    components: { BudgetsList }
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