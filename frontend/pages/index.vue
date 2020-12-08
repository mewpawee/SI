<template lang="html">
  <v-carousel v-model="model">
    <v-carousel-item v-for="(item, i) in items" :key="i" :src="item.src">
      <v-row class="fill-height" align="center" justify="center">
        <div class="display-3">Security {{ i + 1 }}</div>
      </v-row>
    </v-carousel-item>
  </v-carousel>
</template>

<script>
import security1 from '@/static/slide/security-1.jpg'
import security2 from '@/static/slide/security-2.jpg'
export default {
  layout: 'session',
  data: () => ({
    model: 0,
    items: [
      {
        src: security1
      },
      {
        src: security2
      }
    ]
  }),
  mounted() {
    if (this.$auth.loggedIn) {
      this.$router.replace({ name: 'login' })
    }
  }
}
</script>
