export const state = () => ({
  leftDrawer: true,
  miniVariant: false
})

export const mutations = {
  set_minivariant(state, newMiniVariant) {
    state.miniVariant = newMiniVariant
  },
  set_left_drawer(state, newLeftDrawerState) {
    state.leftDrawer = newLeftDrawerState
  }
}
