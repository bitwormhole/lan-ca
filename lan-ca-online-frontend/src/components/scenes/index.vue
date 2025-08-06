<script lang="js">

import MyToolbar from './SceneToolbar.vue'
import MyTable from './SceneTable.vue'
import MyLoader from './SceneLoader.vue'

export default {

    components: { MyToolbar, MyTable, MyLoader },

    data() {
        return {}
    },

    methods: {},

    mounted() {

    },
}
</script>

<style lang="css"></style>

<template>
    <div class="root">
        <my-loader auto-fetch></my-loader>
        <my-toolbar></my-toolbar>
        <my-table></my-table>
    </div>
</template>
