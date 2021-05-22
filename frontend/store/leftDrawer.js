export const state = () => ({
  status: true
})

export const mutations = {
  toggle(state) {
    state.status = !state.status
  }
}
