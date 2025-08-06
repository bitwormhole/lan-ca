<script lang="js">

import { useCertificatesStore } from '@/stores/certificates'

const certs_store = useCertificatesStore();

export default {

    data() {
        return {}
    },

    computed: {
        items() {
            return certs_store.items;
        },
    },

    methods: {},

    mounted() {

    },
}
</script>

<style lang="css"></style>

<template>
    <div class="root">
        <el-table :data="items">
            <el-table-column prop="id" label="ID" />
            <el-table-column prop="label" label="标签" />
            <el-table-column prop="domain_names" label="目标域名" />

            <el-table-column prop="not_before" label="生效日期" />
            <el-table-column prop="not_after" label="失效日期" />

            <el-table-column prop="comment" label="备注" />
        </el-table>
    </div>
</template>
