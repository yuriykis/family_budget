<template>
    <v-container>
        <v-row>
            <!-- cards with family budget -->
            <v-card>
                <v-card-title>
                    <h3 class="headline mb-0">Family Budget</h3>
                </v-card-title>
                <v-card-text>
                    <v-row>
                        <v-col cols="12" sm="6" md="4">
                            <v-card class="elevation-12">
                                <v-toolbar dark flat>
                                    <v-spacer />
                                    <v-toolbar-title>Income</v-toolbar-title>
                                    <v-spacer />
                                </v-toolbar>
                                <v-card-text>
                                    <v-form>
                                        <v-text-field prepend-icon="mdi-cash" name="income" label="Income" type="text">
                                        </v-text-field>
                                    </v-form>
                                </v-card-text>
                                <v-card-actions>
                                    <v-spacer></v-spacer>
                                    <v-btn dark>Save</v-btn>
                                </v-card-actions>
                            </v-card>
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                            <v-card class="elevation-12">
                                <v-toolbar dark flat>
                                    <v-spacer />
                                    <v-toolbar-title>Expenses</v-toolbar-title>
                                    <v-spacer />
                                </v-toolbar>
                                <v-card-text>
                                    <v-form>
                                        <v-text-field prepend-icon="mdi-cash" name="expenses" label="Expenses"
                                            type="text"></v-text-field>
                                    </v-form>
                                </v-card-text>
                                <v-card-actions>
                                    <v-spacer></v-spacer>
                                    <v-btn dark>Save</v-btn>
                                </v-card-actions>
                            </v-card>
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                            <v-card class="elevation-12">
                                <v-toolbar dark flat>
                                    <v-spacer />
                                    <v-toolbar-title>Balance</v-toolbar-title>
                                    <v-spacer />
                                </v-toolbar>
                                <v-card-text>
                                    <v-form>
                                        <v-text-field prepend-icon="mdi-cash" name="balance" label="Balance"
                                            type="text"></v-text-field>
                                    </v-form>
                                </v-card-text>
                                <v-card-actions>
                                    <v-spacer></v-spacer>
                                    <v-btn dark>Save</v-btn>
                                </v-card-actions>
                            </v-card>
                        </v-col>
                    </v-row>
                </v-card-text>
            </v-card>
        </v-row>
    </v-container>
</template>
