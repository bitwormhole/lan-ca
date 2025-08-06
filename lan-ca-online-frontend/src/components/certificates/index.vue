<script lang="js">

import MyToolbar from './CertificateToolbar.vue'
import MyTable from './CertificateTable.vue'
import MyLoader from './CertificateLoader.vue'

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
        <my-toolbar> </my-toolbar>
        <my-table></my-table>
    </div>
</template>
