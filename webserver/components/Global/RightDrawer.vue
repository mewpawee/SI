<template>
  <v-navigation-drawer
    v-model="rightDrawer"
    :right="true"
    :clipped="false"
    :fixed="true"
    app
  >
  </v-navigation-drawer>
</template>

<script>
export default {
  computed: {
    rightDrawer: {
      get() {
        return this.$nuxt.$store.state.rightDrawer
      },
      set(val) {
        return this.$store.commit('set_right_drawer', val)
      }
    }
  },
  methods: {
    handleLoggedOut() {
      return this.$nuxt.$auth.logout()
    }
  }
}
</script>

<style lang="css" scoped></style>
