<template lang="html">
  <div>
    <div v-for="item in companies" :key="item.name">
      <v-card>
        <v-card-title>{{ item.name }}</v-card-title>
        <v-card-actions>
          <v-btn color="primary lighten-2" text>
            Endpoints
          </v-btn>

          <v-spacer></v-spacer>

          <v-btn icon @click="toggle(item)">
            <v-icon>{{ show ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
          </v-btn>
        </v-card-actions>
        <v-expand-transition>
          <div v-show="item.show">
            <v-divider></v-divider>

            <v-card-text>
              <v-list>
                <v-list-item v-for="endpoint in item.endpoints" :key="endpoint">
                  {{ endpoint }}
                </v-list-item>
              </v-list>
            </v-card-text>
          </div>
        </v-expand-transition>
      </v-card>
    </div>
  </div>
</template>

<script>
export default {
  auth: false,
  layout: 'admin',
  async middleware({ $auth, redirect }) {
    if ($auth.loggedIn) {
      const roles = await $auth.user.realm_access.roles
      if (!roles.includes('admin')) {
        return redirect('/')
      }
    } else {
      return redirect('/')
    }
  },
  data: () => ({
    companies: [
      { name: 'Start', endpoints: [1, 2, 3, 4], show: false },
      { name: 'Start', endpoints: [1, 2, 3, 4], show: false },
      { name: 'Start', endpoints: [1, 2, 3, 4], show: false }
    ]
  }),
  methods: {
    toggle(item) {
      item.show = !item.show
    }
  }
}
</script>
