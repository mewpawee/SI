<template>
  <v-app dark>
    <Topbar />
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Topbar from '@/components/Global/Topbar'
export default {
  components: {
    Topbar
  }
}
</script>
