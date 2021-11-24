export const state = () => ({
  error: '',
})

export const mutations = {
  error(state, error) {
    if (error.response && error.response.data) state.error = error.response.data
    else state.error = error
  }
}