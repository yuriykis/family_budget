<template>
  <v-container>
    <v-row aling="center" justify="center">
      <v-col cols=" 12" sm="8" md="4">
        <v-card class="elevation-12">
          <v-toolbar dark flat>
            <v-spacer />
            <v-toolbar-title>FamilyBudget</v-toolbar-title>
            <v-spacer />
          </v-toolbar>
          <v-card-text>
            <v-form>
              <v-text-field prepend-icon="mdi-account" name="login" label="Login" type="text"></v-text-field>
              <v-text-field id="password" prepend-icon="mdi-lock" name="password" label="Password" type="password">
              </v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn dark>Login</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Login',

  data: () => ({

  }),
}
</script>
