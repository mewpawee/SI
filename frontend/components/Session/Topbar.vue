<template lang="html">
  <v-app-bar clipped-left fixed app>
    <v-tabs>
      <v-tab v-for="(item, i) in items" :key="i" :to="item.to"
        >{{ item.title }}
      </v-tab>
      <v-spacer />
      <v-btn class="ma-2" color="indigo" outlined to="login">
        Login
      </v-btn>
    </v-tabs>
  </v-app-bar>
</template>

<script>
export default {
  data() {
    return {
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
  }
}
</script>
