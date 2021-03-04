<template>
  <v-navigation-drawer
    v-model="leftDrawer"
    :mini-variant="miniVariant"
    :clipped="true"
    :fixed="true"
    app
  >
    <v-list>
      <v-list-item
        v-for="(item, i) in items"
        :key="i"
        :to="item.to"
        router
        exact
      >
        <v-list-item-action>
          <v-icon>{{ item.icon }}</v-icon>
        </v-list-item-action>
        <v-list-item-content>
          <v-list-item-title v-text="item.title" />
        </v-list-item-content>
      </v-list-item>
    </v-list>
    <template v-slot:append>
      <div class="pa-2">
        <v-btn icon @click.stop="miniVariantToggle">
          <v-icon>mdi-{{ `chevron-${miniVariant ? 'right' : 'left'}` }}</v-icon>
        </v-btn>
      </div>
    </template>
  </v-navigation-drawer>
</template>

<script>
export default {
  data() {
    return {
      items: [
        {
          icon: 'mdi-shield-search',
          title: 'Scanner',
          to: { name: 'scanner' }
        },
        {
          icon: 'mdi-history',
          title: 'History',
          to: { name: 'history' }
        }
      ]
    }
  },
  computed: {
    miniVariant: {
      get() {
        return this.$nuxt.$store.state.miniVariant
      }
    },
    leftDrawer: {
      get() {
        return this.$nuxt.$store.state.leftDrawer
      }
    }
  },
  methods: {
    miniVariantToggle() {
      this.$store.commit(
        'set_minivariant',
        !this.$nuxt.$store.state.miniVariant
      )
    }
  }
}
</script>
