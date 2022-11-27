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
              <v-col>
                <SearchBar v-show="isAuthenticated" />
              </v-col>
            </v-row>
            <v-row class="ma-0">
              <v-col>
                <v-row justify="center" align="center" v-if="loading">
                  <Loader />
                </v-row>
                <transition v-else name="fade">
                  <router-view />
                </transition>
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
import Loader from '@/components/Loader'
export default {
  name: 'App',
  data: () => ({
    loading: false
  }),

  components: {
    SideMenu,
    SearchBar,
    Loader
  },
  computed: {
    ...mapGetters('auth', ['isAuthenticated']),
  },
};
</script>

<style>
fade-enter-active,
.fade-leave-active {
  transition-property: opacity;
  transition-duration: 0.25s;
}

.fade-enter-active {
  transition-delay: 0.25s;
}

.fade-enter,
.fade-leave-active {
  opacity: 0;
}
</style>