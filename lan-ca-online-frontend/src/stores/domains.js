
import { defineStore } from 'pinia'
import { useAxiosStore } from './axios'


const axios_store = useAxiosStore();


export const useDomainsStore = defineStore('domains-store', {

    state() {
        return {
            inner_items: []
        }
    },

    getters: {

        items(state) {
            // const list = [];
            // list.push({ domain: 'x.y.z', ip: '1.0.2.4', cert: '888888', })
            // return list; 
            return state.inner_items;
        },

    },

    actions: {

        fetch() {
            let url = '/api/v1/domains'
            let method = 'GET';
            let self = this;
            let pro = axios_store.request({ method, url }).then((res) => {
                let vo = res.data;
                let items = vo.domains;
                self.inner_items = items;
            });
            return pro;
        },

        insert(vo) {
            let url = '/api/v1/domains'
            let method = 'POST';
            let data = vo;
            let pro = axios_store.request({ method, url, data }).then((res) => {
                // let vo = res.data;
                // let items = vo.domains;
                // self.inner_items = items;
            });
            return pro;
        },
    },

});
