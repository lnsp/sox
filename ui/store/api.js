export const state = () => ({
  version: '',
  machines: [],
});

export const mutations = {
  version(state, version) {
    state.version = version
  },
  machines(state, machines) {
    state.machines = machines
  }
}

export const actions = {
  async connect({ commit }) {
    let response = await this.$axios.$get('/')
    commit('version', response.version)
  },
  async machines({ commit }) {
    let response = await this.$axios.$get('/machines')
    commit('machines', response)
  }
}