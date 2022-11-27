<template>
    <v-navigation-drawer permanent dark app>
        <v-list>
            <v-list-item class="px-2">
                <v-row justify="center">
                    <v-col cols="12" class="text-center">
                        <v-list-item-avatar size="100">
                            <v-img src="https://randomuser.me/api/portraits/men/1.jpg"></v-img>
                        </v-list-item-avatar>
                    </v-col>
                </v-row>
            </v-list-item>

            <v-list-item>
                <v-list-item-content>
                    <v-list-item-title class="text-h6">
                        Name Surname
                    </v-list-item-title>
                    <v-list-item-subtitle>name.surname@gmail.com</v-list-item-subtitle>
                </v-list-item-content>
            </v-list-item>
        </v-list>

        <v-divider></v-divider>

        <v-list nav dense>
            <v-list-item link>
                <v-list-item-icon>
                    <v-icon>mdi-view-dashboard</v-icon>
                </v-list-item-icon>
                <v-list-item-title>Dashboard</v-list-item-title>
            </v-list-item>
            <v-list-item link @click="onMyDashboardClick">
                <v-list-item-icon>
                    <v-icon>mdi-cash-multiple</v-icon>
                </v-list-item-icon>
                <v-list-item-title>My budgets</v-list-item-title>
            </v-list-item>
            <v-list-item link>
                <v-list-item-icon>
                    <v-icon>mdi-account-multiple</v-icon>
                </v-list-item-icon>
                <v-list-item-title>Shared with me</v-list-item-title>
            </v-list-item>
        </v-list>
        <v-spacer></v-spacer>
        <template v-slot:append>
            <v-row justify="center" class="mb-5">
                <v-col cols="12" class="text-center">
                    <v-btn text @click="logout">
                        <span>Logout</span>
                        <v-icon>mdi-logout-variant</v-icon>
                    </v-btn>
                </v-col>
            </v-row>
        </template>

    </v-navigation-drawer>
</template>

<script>
import { mapActions } from 'vuex'
export default {
    name: 'SideMenu',
    data: () => ({
    }),
    methods: {
        ...mapActions('auth', ['logoutUser']),
        onMyDashboardClick() {
            this.$router.push({ name: 'Home' })
        },
        logout() {
            this.logoutUser()
            this.$router.push({ name: 'Login' })
        }
    }
}
</script>