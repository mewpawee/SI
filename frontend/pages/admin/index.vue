<template>
  <div>
    <v-data-table
      :headers="headers"
      :items="scans"
      sort-by="Start"
      class="elevation-1"
      item-key="name"
      :search="search"
    >
      <template v-slot:top>
        <v-text-field
          v-model="search"
          label="Search"
          class="mx-4"
        ></v-text-field>
        <v-toolbar flat>
          <v-toolbar-title>Scans</v-toolbar-title>
        </v-toolbar>
      </template>
    </v-data-table>
  </div>
</template>

<script>
import { adminGetEndpoints } from '@/utils/backendAPI.js'
export default {
  auth: false,
  layout: 'admin',
  async fetch() {
    const result = await adminGetEndpoints()
    this.scans = result.data
  },
  fetchDelay: 1000,
  fetchOnServer: false,
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
    search: '',
    polling: null,
    headers: [
      {
        text: 'Company',
        align: 'start',
        sortable: true,
        value: 'company'
      },
      { text: 'Endpoint', value: 'endpoint' }
    ],
    scans: []
  }),
  beforeDestroy() {
    clearInterval(this.polling)
  },
  created() {
    this.pollData()
  },
  methods: {
    pollData() {
      this.polling = setInterval(() => {
        this.$fetch()
      }, 3000)
    }
  }
}
</script>
