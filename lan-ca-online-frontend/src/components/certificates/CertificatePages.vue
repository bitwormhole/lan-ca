<script lang="js">
import { ElPagination } from 'element-plus';
import { useCertificatesStore } from '@/stores/certificates'
import { Pagination } from '@/js/pagination'

const theCertificatesStore = useCertificatesStore();

export default {

    data() {
        return {}
    },

    computed: {
        pagination() {
            let p = theCertificatesStore.pagination;
            return Pagination.normalizePagination(p)
        },
    },

    methods: {

        reload_with_pagination(pagination) {

            let size = pagination.size;
            let page = pagination.page;

            if (size == null) {
                size = this.pagination.size;
            }
            if (page == null) {
                page = this.pagination.page;
            }

            let params = { size, page }
            theCertificatesStore.fetch(params);
        },

        onCurrentPageChange(value) {
            this.reload_with_pagination({ page: value });
        },

        onPageSizeChange(value) {
            this.reload_with_pagination({ size: value });
        },

    },
}
</script>

<style lang="css"></style>

<template>
    <ElPagination background layout="total, prev, pager, next, sizes" :current-page="pagination.page"
        :page-size="pagination.size" :total="pagination.total" :page-sizes="[10, 20, 50, 100]"
        @update:current-page="onCurrentPageChange" @update:page-size="onPageSizeChange">
    </ElPagination>
</template>
