<template>
  <div>
    <v-data-table
      :headers="headers"
      :items="scans"
      sort-by="Start"
      class="elevation-1"
      item-key="name"
    >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>History</v-toolbar-title>
        </v-toolbar>
      </template>
      <template #[`item.start`]="{item}">
        <span>{{
          new Date(item.start).toLocaleString('en-GB', {
            timeZone: 'Asia/Bangkok'
          })
        }}</span>
      </template>
      <template #[`item.Complete`]="{item}">
        <span>{{
          item.Complete == '0001-01-01T00:00:00Z'
            ? null
            : new Date(item.Complete).toLocaleString('en-GB', {
                timeZone: 'Asia/Bangkok'
              })
        }}</span>
      </template>
      <template v-slot:item.actions="{ item }">
        <span v-if="item.status == 'success' && item.downloading == false">
          <v-icon small class="mr-2" @click="downloadReport(item)">
            mdi-download
          </v-icon>
        </span>
        <v-progress-circular
          v-else-if="item.downloading == true"
          indeterminate
          :size="20"
          color="grey"
        />
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="initialize">
          Reset
        </v-btn>
      </template>
    </v-data-table>
  </div>
</template>

<script>
import { getAllScans, downloadReport } from '@/utils/backendAPI.js'
export default {
  middleware: 'is-admin',
  async fetch() {
    const result = await getAllScans()
    const manipulatedData = await this.manipulateData(result.data)
    this.scans = manipulatedData
  },
  fetchDelay: 1000,
  fetchOnServer: false,
  data: () => ({
    polling: null,
    dialog: false,
    dialogDelete: false,
    headers: [
      {
        text: 'Scan ID',
        align: 'start',
        sortable: false,
        value: 'scanid'
      },
      { text: 'Start', value: 'start' },
      { text: 'Complete', value: 'Complete' },
      { text: 'Status', value: 'status' },
      { text: 'Download', align: 'right', value: 'actions', sortable: false }
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
    },
    async downloadReport(item) {
      item.downloading = true
      await downloadReport(item.scanid)
      item.downloading = false
    },
    manipulateData(data) {
      for (const thisData of data) {
        thisData.downloading = false
      }
      return data
    }
  }
}
</script>
