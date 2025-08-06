<script>
import { ElTable } from 'element-plus';
import { ElButton } from 'element-plus';

import { useDomainsStore } from '@/stores/domains'

const domains_store = useDomainsStore();

export default {

    data() {
        return {

        }
    },

    computed: {
        items() {
            return domains_store.items;
        },
    },

    methods: {
    },

    mounted() {
    },

}

</script>

<style>
@media (min-width: 1024px) {
    .about {
        min-height: 100vh;
        display: flex;
        align-items: center;
    }
}
</style>

<template>
    <div class="domain-table-view">

        <!-- <ElButton @click="fetch"> refresh </ElButton>
        <ElButton> + add domain ... </ElButton> -->

        <el-table :data="items">
            <el-table-column prop="id" label="ID" />
            <el-table-column prop="name" label="域名" />
            <el-table-column prop="ipv4_address" label="IP 地址" />
            <el-table-column prop="comment" label="备注" />
            <el-table-column prop="scene" label="应用场景" />
            <el-table-column prop="cert" label="Cert state (todo)" />
        </el-table>
    </div>
</template>
