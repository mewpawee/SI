<template>
  <v-app-bar clipped-left fixed app>
    <v-btn icon @click.stop="leftDrawerToggle">
      <v-icon>mdi-menu</v-icon>
    </v-btn>
    <v-toolbar-title v-text="title" />
    <v-spacer />
    <v-toolbar-title v-text="user" />
    <v-btn icon @click.stop="signOut">
      <v-icon>mdi-export</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
export default {
  data() {
    return {
      title: 'Cyber Security Index',
      user: ''
    }
  },
  mounted() {
    if (this.$nuxt.$auth.loggedIn) {
      this.user = 'Welcome, ' + this.$nuxt.$auth.user.preferred_username + '!'
    }
  },
  methods: {
    signOut() {
      this.$auth.logout('local')
    },
    leftDrawerToggle() {
      this.$store.commit('leftDrawer/toggle')
    }
  }
}
</script>
>
