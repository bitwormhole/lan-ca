<template>
    <div class="domain-table-view">
        <h1> domains </h1>

        <ElButton @click="fetch"> refresh </ElButton>
        <ElButton> + add domain ... </ElButton>

        <el-table :data="items">
            <el-table-column prop="name" label="Domain Name" />
            <el-table-column prop="ipv4addr" label="IP" />
            <el-table-column prop="cert" label="Certificate" />
            <el-table-column prop="comment" label="备注" />
        </el-table>
    </div>
</template>

<style>
@media (min-width: 1024px) {
    .about {
        min-height: 100vh;
        display: flex;
        align-items: center;
    }
}
</style>


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
        fetch() {
            domains_store.fetch();
        },
    },

    mounted() {
        this.fetch()
    },

}

</script>
