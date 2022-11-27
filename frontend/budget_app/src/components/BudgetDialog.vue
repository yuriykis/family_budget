<template>
    <v-card>
        <v-card-title>
            <span class="headline">{{ title }}</span>
        </v-card-title>
        <v-card-text>
            <v-container>
                <v-row>
                    <v-col cols="12">
                        <v-text-field v-model="name" label="Name" required></v-text-field>
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols="12">
                        <v-text-field v-model="description" label="Description" required></v-text-field>
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
    name: "BudgetDialog",
    props: {
        title: {
            type: String,
            required: true,
            default: "Budget Properties"
        }
    },
    data: () => ({

        name: "",
        description: ""
    }),
    methods: {
        addBudget() {
            this.dialog = true;
        },
        close() {
            this.$emit('on-close');
        },
        save() {
            this.dialog = false;
            this.$emit("add-budget", {
                name: this.name,
                description: this.description
            });
        }
    }
};
</script>