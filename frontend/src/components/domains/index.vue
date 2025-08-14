<script lang="js">

import MyLoader from './DomainLoader.vue'
import MyToolbar from './DomainToolbar.vue'
import MyTable from './DomainTable.vue'
import MyPages from './DomainPages.vue'

export default {

    components: { MyLoader, MyToolbar, MyTable, MyPages },

    data() {
        return {}
    },

    methods: {
        handleAction(p) {
            if (p.name == 'fetch') {
                this.$refs.loader.fetch()
            }
        },
    },

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
        <my-loader ref="loader" auto-fetch></my-loader>
        <my-toolbar @action="handleAction"></my-toolbar>
        <my-table class="the-table"></my-table>
        <my-pages></my-pages>
    </div>
</template>
