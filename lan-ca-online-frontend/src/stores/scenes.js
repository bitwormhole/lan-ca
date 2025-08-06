
import { defineStore } from 'pinia'
import { useAxiosStore } from './axios'

const axios_store = useAxiosStore();

export const useScenesStore = defineStore('scenes-store', {

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
            let url = '/api/v1/scenes'
            let method = 'GET';
            let self = this;
            let pro = axios_store.request({ method, url }).then((res) => {
                let vo = res.data;
                let items = vo.scenes;
                self.inner_items = items;
            })
            return pro;
        },

        insert(vo) {
            let url = '/api/v1/scenes'
            let method = 'POST';
            let self = this;
            let data = vo
            let pro = axios_store.request({ method, url, data }).then((res) => {
                // let vo = res.data;
                // let items = vo.scenes;
                // self.inner_items = items;
            })
            return pro;
        },

    },

});
