<template>
    <v-card>
        <v-card-title>
            <span class="headline">Share budget</span>
        </v-card-title>
        <v-card-text>
            <v-container>
                <v-row>
                    <v-col cols="12">
                        <v-select v-model="selectedUser" :items="users" label="Select user" required
                            item-text="username">
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
import { mapGetters, mapActions } from 'vuex';
export default {
    name: "ShareBudgetDialog",
    data: () => ({
        selectedUser: "",
    }),
    computed: {
        ...mapGetters('users', ['users'])
    },
    async mounted() {
        await this.getAllUsersAction();
    },
    methods: {
        ...mapActions('users', ['getAllUsersAction']),
        close() {
            this.$emit('on-close');
        },
        save() {
            this.$emit("share-budget", {
                user: this.users.find(u => u.username === this.selectedUser)
            });
        },
    }
}
</script>