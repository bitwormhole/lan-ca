<script lang="js">

import { useScenesStore } from '@/stores/scenes'

const scenes_store = useScenesStore();

export default {

    data() {
        return {}
    },

    computed: {
        items() {
            return scenes_store.items;
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
            <el-table-column prop="name" label="名称" />
            <el-table-column prop="label" label="标签" />
            <el-table-column prop="comment" label="备注" />
        </el-table>
    </div>
</template>
