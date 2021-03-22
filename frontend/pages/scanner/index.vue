<template>
  <div>
    <v-data-table
      v-model="selected"
      :headers="headers"
      :items="endpoints"
      sort-by="calories"
      class="elevation-1"
      item-key="endpoint"
      show-select
    >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>Endpoints</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ on, attrs }">
              <v-btn color="primary" dark class="mb-2" v-bind="attrs" v-on="on">
                <v-icon left>mdi-plus-thick</v-icon>
                Add Endpoint
              </v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">{{ formTitle }}</span>
              </v-card-title>

              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12" sm="12" md="12">
                      <v-text-field
                        v-model="editedItem.endpoint"
                        label="Endpoint"
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="close">
                  Cancel
                </v-btn>
                <v-btn color="blue darken-1" text @click="save">
                  Save
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
          <v-dialog v-model="dialogDelete" max-width="500px">
            <v-card>
              <v-card-title class="headline"
                >Are you sure you want to delete this item?</v-card-title
              >
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="closeDelete"
                  >Cancel</v-btn
                >
                <v-btn color="blue darken-1" text @click="deleteItemConfirm"
                  >OK</v-btn
                >
                <v-spacer></v-spacer>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-toolbar>
      </template>
      <template v-slot:[`item.actions`]="{ item }">
        <v-icon small class="mr-2" @click="editItem(item)">
          mdi-pencil
        </v-icon>
        <v-icon small @click="deleteItem(item)">
          mdi-delete
        </v-icon>
      </template>
      <template v-slot:no-data>
        <v-title> Add Endpoint</v-title>
      </template>
    </v-data-table>
  </div>
</template>

<script>
import {
  getEndpoints,
  addEndpoint,
  deleteEndpoint
} from '@/utils/backendAPI.js'
export default {
  async fetch() {
    const result = await getEndpoints()
    this.endpoints = result.data
  },
  data: () => ({
    dialog: false,
    dialogDelete: false,
    selected: [],
    headers: [
      {
        text: 'Name',
        align: 'start',
        sortable: false,
        value: 'endpoint'
      },
      { text: 'Actions', align: 'right', value: 'actions', sortable: false }
    ],
    endpoints: [],
    editedIndex: -1,
    editedItem: {
      endpoint: ''
    },
    defaultItem: {
      endpoint: ''
    }
  }),

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'Add Endpoint' : 'Edit Endpoint'
    }
  },

  watch: {
    dialog(val) {
      val || this.close()
    },
    dialogDelete(val) {
      val || this.closeDelete()
    }
  },

  created() {
    this.$fetch()
  },

  methods: {
    editItem(item) {
      this.editedIndex = this.endpoints.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },

    deleteItem(item) {
      this.editedIndex = this.endpoints.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
    },

    deleteItemConfirm() {
      deleteEndpoint(this.endpoints[this.editedIndex].endpoint)
      this.endpoints.splice(this.editedIndex, 1)
      this.closeDelete()
    },

    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    closeDelete() {
      this.dialogDelete = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    async save() {
      if (this.editedIndex > -1) {
        Object.assign(this.endpoints[this.editedIndex], this.editedItem)
      } else {
        this.endpoints.push(this.editedItem)
        await addEndpoint(this.editedItem.endpoint)
      }
      this.close()
    }
  }
}
</script>
