// import { ref, computed } from 'vue'

import { defineStore } from 'pinia'

import axios from 'axios'


export const useAxiosStore = defineStore('axios', {

  state: () => ({
    foo: 0,
  }),

  getters: {},

  actions: {

    execute(config) {
      let p = axios(config);
      return p;
    }

  },

})
