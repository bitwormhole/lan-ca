<script lang="js">

import { useDomainsStore } from '@/stores/domains'

const domains_store = useDomainsStore();

export default {

    data() {
        return {}
    },

    methods: {
        fetch() {
            domains_store.fetch();
        },
    },

    mounted() {
        if (this.autoFetch) {
            this.fetch();
        }
    },

    props: {
        autoFetch: Boolean
    },
}
</script>

<style lang="css"></style>

<template>
    <div class="root"> </div>
</template>
