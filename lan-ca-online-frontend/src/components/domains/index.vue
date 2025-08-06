<script lang="js">

import MyLoader from './DomainLoader.vue'
import MyToolbar from './DomainToolbar.vue'
import MyTable from './DomainTable.vue'

export default {

    components: { MyLoader, MyToolbar, MyTable, },

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

<style lang="css"></style>

<template>
    <div class="root">
        <my-loader ref="loader" auto-fetch></my-loader>
        <my-toolbar @action="handleAction"></my-toolbar>
        <my-table></my-table>
    </div>
</template>
