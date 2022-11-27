<template>
    <v-container>
        <v-row align="center" justify="center">
            <v-col md="12" lg="6" align="center" justify="center">
                <v-card v-if="loading" flat>
                    <Loader />
                </v-card>
                <v-card v-else>
                    <v-toolbar dark flat>
                        <v-spacer />
                        <v-toolbar-title>Register in FamilyBudget App!</v-toolbar-title>
                        <v-spacer />
                    </v-toolbar>
                    <v-card-text>
                        <v-container>
                            <v-row>
                                <v-col cols="12">
                                    <v-text-field color="#363636" v-model="user.username" label="Username" required>
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col cols="12">
                                    <v-text-field color="#363636" v-model="user.first_name" label="First Name" required>
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col cols="12">
                                    <v-text-field color="#363636" v-model="user.last_name" label="Last Name" required>
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col cols="12">
                                    <v-text-field color="#363636" v-model="user.email" label="Email" required>
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col cols="12">
                                    <v-text-field color="#363636" v-model="user.password" label="Password" required
                                        type="password">
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col cols="12">
                                    <v-text-field color="#363636" v-model="password2" label="Confirm password" required
                                        type="password">
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <v-row align="center" justify="center">
                                <v-col cols="3">
                                    <v-btn dark @click="registerUser">Register</v-btn>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col cols="3">
                                    <v-btn color="#363636" text style="text-transform:none !important;" @click="login">
                                        Already have an account? Login
                                    </v-btn>
                                </v-col>
                            </v-row>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import { register } from '@/services/api'
import Loader from '@/components/Loader'
export default {
    name: "Register",
    data: () => ({
        user: {
            username: "",
            first_name: "",
            last_name: "",
            email: "",
            password: ""
        },
        password2: "",
        loading: false
    }),
    components: {
        Loader
    },
    methods: {
        async registerUser() {
            if (this.ifFieldsAreEmpty()) {
                alert("Please fill all fields")
                return
            }
            if (this.user.password !== this.password2) {
                alert("Passwords do not match");
                return;
            }
            this.loading = true;

            try {
                await register(this.user);
                this.$router.push("/login");
            } catch (error) {
                alert("Error while registering user");
            } finally {
                this.loading = false;
            }
        },
        ifFieldsAreEmpty() {
            if (this.user.username === "" || this.user.first_name === "" || this.user.last_name === "" || this.user.email === "" || this.user.password === "") {
                return true;
            }
            return false;
        },
        login() {
            this.$router.push("/login");
        }
    }
};
</script>