
import { defineStore } from 'pinia'
import { useAxiosStore } from './axios'
import { Pagination } from '@/js/pagination'


const axios_store = useAxiosStore();


export const useDomainsStore = defineStore('domains-store', {

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
        pagination(state) {
            return state.inner_pagination;
        },

        items(state) {
            return state.inner_items;
        },
    },

    actions: {

        fetch( params ) {
            let url = '/api/v1/domains'
            let method = 'GET';
            let self = this;
            let pro = axios_store.request({ method, url , params   }).then((res) => {
                let vo = res.data;
                let items = vo.domains;
                self.inner_pagination = Pagination.normalizePagination(vo.pagination);
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
