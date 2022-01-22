import Vue from 'vue';

export const state = () => ({
  version: '',
  machines: [],
  sshKeys: [],
  images: [],
  networks: [],
  machineDetails: {},
  activities: [],
});

export const mutations = {
  version(state, version) {
    state.version = version
  },
  machines(state, machines) {
    state.machines = machines
  },
  machineDetails(state, { id, details }) {
    Vue.set(state.machineDetails, id, details)
  },
  sshKeys(state, sshKeys) {
    state.sshKeys = sshKeys
  },
  images(state, images) {
    state.images = images
  },
  networks(state, networks) {
    state.networks = networks
  },
  activities(state, activities) {
    state.activities = activities
  }
}

export const getters = {
  reversedActivities: (state) => {
    return state.activities.slice().reverse()
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
  },
  async sshKeys({ commit }) {
    let response = await this.$axios.$get('/ssh-keys')
    commit('sshKeys', response)
  },
  async images({ commit }) {
    let response = await this.$axios.$get('/images')
    commit('images', response)
  },
  async networks({ commit }) {
    let response = await this.$axios.$get('/networks')
    commit('networks', response)
  },
  async machineDetails({ commit }, id) {
    let details = await this.$axios.$get('/machines/' + id)
    commit('machineDetails', { id, details })
  },
  async activities({ commit }) {
    let response = await this.$axios.$get('/activities')
    commit('activities', response.activities)
  }
}