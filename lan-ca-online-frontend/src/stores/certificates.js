
import { defineStore } from 'pinia'
import { useAxiosStore } from './axios'

const axios_store = useAxiosStore();

export const useCertificatesStore = defineStore('certificates-store', {

    state() {
        return {
            inner_items: []
        }
    },

    getters: {

        items(state) {
            return state.inner_items;
        },

    },

    actions: {

        fetch() {
            let url = '/api/v1/certificates'
            let method = 'GET';
            let self = this;
            axios_store.request({ method, url }).then((res) => {
                let vo = res.data;
                let items = vo.certificates;
                self.inner_items = items;
            })
        },

    },

});
