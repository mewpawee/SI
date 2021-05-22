<template>
  <p>redirect</p>
</template>
<script>
export default {
  async mounted() {
    try {
      if (!this.$auth.loggedIn) {
        await this.$auth.loginWith('keycloak')
      }
    } catch (err) {
      console.log(err)
    }
  }
}
</script>
