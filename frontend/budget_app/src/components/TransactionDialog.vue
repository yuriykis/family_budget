<template>
    <v-card>
        <v-card-title>
            <span class="headline">New transaction</span>
        </v-card-title>
        <v-card-text>
            <v-container>
                <v-row>
                    <v-col cols="12">
                        <v-text-field v-model="title" label="Title" required></v-text-field>
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols="12">
                        <v-text-field v-model="amount" label="Amount" required></v-text-field>
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols="12">
                        <v-select v-model="type" :items="types" label="Type" required></v-select>
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols="12">
                        <v-select v-model="selectedCategory" :items="category_names" label="Category" required>
                        </v-select>
                    </v-col>
                </v-row>
            </v-container>
        </v-card-text>
        <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
            <v-btn color="blue darken-1" text @click="save">Save</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
export default {
    name: "TransactionDialog",
    props: {
        categories: {
            type: Array,
            required: true
        }
    },
    data: () => ({
        title: "",
        amount: 0,
        type: "",
        types: ["INCOME", "EXPENSE"],
        selectedCategory: "",
        category_names: []
    }),
    methods: {
        close() {
            this.$emit('on-close');
        },
        save() {
            this.$emit("transaction-action", {
                type: this.type,
                amount: this.amount,
                title: this.title,
                category: this.mapCategoryNameToId(this.selectedCategory)
            });
        },
        mapCategoryNameToId(categoryName) {
            return this.categories.find(category => category.name === categoryName).id
        }
    },
    created() {
        this.category_names = this.categories.map(category => category.name);
    }
} 
</script>