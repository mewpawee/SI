<template lang="html">
  <v-app-bar clipped-left fixed app>
    <v-tabs>
      <v-tab v-for="(item, i) in items" :key="i" :to="item.to"
        >{{ item.title }}
      </v-tab>
      <v-spacer />
      <v-btn
        class="ma-2"
        color="indigo"
        outlined
        @click.stop="dialogLogin = true"
      >
        Login
      </v-btn>
    </v-tabs>

    <v-dialog v-model="dialogLogin" max-width="500">
      <v-card color="accent" class="justify-center">
        <v-card-title>Login</v-card-title>
        <hr />
        <v-btn large block @click="handleLoginClicked">
          <v-icon class="mx-5">mdi-google</v-icon>
          Sign in with Google
        </v-btn>
      </v-card>
    </v-dialog>
  </v-app-bar>
</template>

<script>
export default {
  data() {
    return {
      dialogLogin: false,
      usernameLogin: '',
      passwordLogin: '',
      items: [
        {
          icon: 'mdi-monitor-dashboard',
          title: 'Cyber Security Index',
          to: { name: 'index' }
        },
        {
          icon: 'mdi-monitor-dashboard',
          title: 'Overall Security Index',
          to: { name: 'overall' }
        }
      ]
    }
  },
  mounted() {
    if (this.$auth.loggedIn) {
      this.$router.push({ name: 'dashboard' })
    }
  },
  methods: {
    async handleLoginClicked() {
      try {
        await this.$auth.loginWith('google')
      } catch (err) {
        this.errorLogin = true
        console.log(err)
      }
    }
  }
}
</script>
