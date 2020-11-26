<template>
  <p>redirect</p>
</template>
<script>
export default {
  mounted() {
    try {
      if (this.$auth.loggedIn) {
        this.$router.replace({ name: 'scanner' })
      } else this.$router.replace({ name: 'index' })
    } catch (err) {
      this.$router.replace({ name: 'index' })
    }
  }
}
</script>
