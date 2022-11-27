<template>
  <v-container>
    <v-row aling="center" justify="center">
      <v-col cols=" 12" sm="8" md="4">
        <v-card v-if="loading" flat>
          <Loader />
        </v-card>
        <v-card v-else class="elevation-12">
          <v-toolbar dark flat>
            <v-spacer />
            <v-toolbar-title>FamilyBudget</v-toolbar-title>
            <v-spacer />
          </v-toolbar>
          <v-card-text>
            <v-form>
              <v-text-field prepend-icon="mdi-account" name="username" label="Username" type="text" color="#363636"
                v-model="userData.username">
              </v-text-field>
              <v-text-field id="password" prepend-icon="mdi-lock" name="password" label="Password" type="password"
                color="#363636" v-model="userData.password">
              </v-text-field>
            </v-form>
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn dark @click="login">Login</v-btn>
            <v-spacer></v-spacer>
          </v-card-actions>
          <v-card-actions>
            <v-btn color="#363636" text style="text-transform:none !important;" @click="register">
              No account? Register
            </v-btn>
            <v-spacer />
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Loader from '@/components/Loader'
import { mapActions, mapGetters } from 'vuex'
export default {
  name: 'Login',

  data: () => ({
    userData: {
      username: '',
      password: '',
    },
    loading: false
  }),
  components: {
    Loader
  },
  methods: {
    ...mapActions('auth', ['loginUser']),
    ...mapGetters('auth', ['isAuthenticated']),
    register() {
      this.$router.push('/register')
    },
    async login() {
      if (this.ifFieldsAreEmpty()) {
        alert('Please fill all fields')
        return
      }
      this.loading = true
      await this.loginUser(this.userData)
      if (this.isAuthenticated()) {
        this.$router.push('/')
      } else {
        alert('Wrong username or password')
      }
      this.loading = false
    },
    ifFieldsAreEmpty() {
      return (
        this.userData.username === '' ||
        this.userData.password === ''
      )
    }
  }
}
</script>
