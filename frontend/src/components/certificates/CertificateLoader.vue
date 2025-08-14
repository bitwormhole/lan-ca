<script lang="js">

import { useCertificatesStore } from '@/stores/certificates'

const certs_store = useCertificatesStore();

export default {

    data() {
        return {}
    },

    methods: {
        fetch() {
            certs_store.fetch()
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
