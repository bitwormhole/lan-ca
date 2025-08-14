
import { defineStore } from 'pinia'
import { useAxiosStore } from './axios'

const axios_store = useAxiosStore();

export const useExampleStore = defineStore('example-store', {

    state() {
        return {
            state1: 1
        }
    },

    getters: {

        getter1(state) {
            return state.state1;
        },

    },

    actions: {

        action1() {
            let url = '/api/v1/example'
            let method = 'GET';
            let self = this;
            axios_store.request({ method, url }).then((res) => {
                // let vo = res.data;
                // let items = vo.domains;
                console.log(res.statusText)
            })
        },
    },

});
