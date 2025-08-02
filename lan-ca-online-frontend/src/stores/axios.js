
import { defineStore } from 'pinia'
import axios from 'axios'


export const useAxiosStore = defineStore('axios-store', {

    state() {
        return {}
    },

    getters: {

    },

    actions: {

        request(config) {
            return axios(config)
        },

    },

});
