
import { defineStore } from 'pinia'
import { useAxiosStore } from './axios'
import { Pagination } from '@/js/pagination'


const axios_store = useAxiosStore();

export const useCertificatesStore = defineStore('certificates-store', {

    state() {
        return {
            inner_pagination: {
                page: 0,
                size: 0,
                total: 0,
            },
            inner_items: [],
        }
    },

    getters: {
        items(state) {
            return state.inner_items;
        },

        pagination(state) {
            return state.inner_pagination;
        },
    },

    actions: {

        fetch(params) {
            let url = '/api/v1/certificates'
            let method = 'GET';
            let self = this;
            let pro = axios_store.request({ method, url, params }).then((res) => {
                let vo = res.data;
                let items = vo.certificates;
                self.inner_items = items;
                self.inner_pagination = Pagination.normalizePagination(vo.pagination);
            })
            return pro
        },

        insert(vo) {
            let url = '/api/v1/certificates'
            let method = 'POST';
            let data = vo;
            let pro = axios_store.request({ method, url, data }).then((res) => {
                // let vo = res.data;
                // let items = vo.certificates;
                // self.inner_items = items;
            })
            return pro
        },

    },

});
