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
    <v-avatar color="accent">
      <span class="white--text headline">{{ avartar }}</span>
    </v-avatar>
  </v-app-bar>
</template>

<script>
export default {
  data() {
    return {
      title: 'Cyber Security Index',
      user: '',
      avatar: ''
    }
  },
  computed: {
    miniVariant: {
      get() {
        return this.$nuxt.$store.state.miniVariant
      }
    }
  },
  mounted() {
    if (this.$nuxt.$auth.loggedIn) {
      this.user = 'Welcome, ' + this.$nuxt.$auth.user.given_name + '!'
      this.avartar =
        this.$nuxt.$auth.user.given_name[0].toUpperCase() +
        this.$nuxt.$auth.user.family_name[0].toUpperCase()
    }
  },
  methods: {
    signOut() {
      this.$auth.logout('local')
    },
    miniVariantToggle() {
      this.$store.commit(
        'set_minivariant',
        !this.$nuxt.$store.state.miniVariant
      )
    },
    leftDrawerToggle() {
      this.$store.commit('set_left_drawer', !this.$nuxt.$store.state.leftDrawer)
    }
  }
}
</script>
>
