export const state = () => ({
  status: false
})

export const mutations = {
  toggle(state) {
    state.status = !state.status
  }
}
