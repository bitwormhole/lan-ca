
import { defineStore } from 'pinia'
import { useAxiosStore } from './axios'
import { Pagination } from '@/js/pagination'

const axios_store = useAxiosStore();


function prepare_fetch_params(ctx, params) {
    if (params == null) {
        let src = ctx.pagination;
        let dst = {}
        dst.page = src.page;
        dst.size = src.size;
        params = dst;
    }
    return params;
}


export const useScenesStore = defineStore('scenes-store', {

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

        reset() {
            this.inner_items = [];
        },

        fetch(params) {
            params = prepare_fetch_params(this, params);
            let url = '/api/v1/scenes'
            let method = 'GET';
            let self = this;
            let pro = axios_store.request({ method, url, params }).then((res) => {
                let vo = res.data;
                let items = vo.scenes;
                self.inner_pagination = Pagination.normalizePagination(vo.pagination);
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
