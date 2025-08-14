<script lang="js">

import MyToolbar from './SceneToolbar.vue'
import MyTable from './SceneTable.vue'
import MyLoader from './SceneLoader.vue'
import MyPages from './ScenePages.vue'

export default {

    components: { MyToolbar, MyTable, MyLoader, MyPages },

    data() {
        return {}
    },

    methods: {},

    mounted() {

    },
}
</script>

<style lang="css">
.the-table {
    margin-top: 10px;
    margin-bottom: 10px;
}
</style>

<template>
    <div class="root">
        <my-loader auto-fetch></my-loader>
        <my-toolbar></my-toolbar>
        <my-table class="the-table"></my-table>
        <my-pages></my-pages>
    </div>
</template>
