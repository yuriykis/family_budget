<template>
  <v-app>
    <v-main>
      <v-row>
        <v-col cols="4">
          <SideMenu v-show="isAuthenticated" />
        </v-col>
        <v-col cols="8" md="10">
          <v-container>
            <v-row>
              <SearchBar v-show="isAuthenticated" />
            </v-row>
            <v-row>
              <v-col>
                <router-view />
              </v-col>
            </v-row>
          </v-container>
        </v-col>
      </v-row>
    </v-main>
  </v-app>
</template>

<script>
import { mapGetters } from 'vuex'
import SideMenu from '@/components/SideMenu'
import SearchBar from './SearchBar.vue';
export default {
  name: 'App',
  data: () => ({
  }),

  components: {
    SideMenu,
    SearchBar
  },
  computed: {
    ...mapGetters('auth', ['isAuthenticated']),
  },
};
</script>
