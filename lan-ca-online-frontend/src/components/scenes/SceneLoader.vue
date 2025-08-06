<script lang="js">

import { useScenesStore } from '@/stores/scenes'

const scenes_store = useScenesStore();

export default {

    data() {
        return {}
    },

    methods: {
        fetch() {
            scenes_store.fetch()
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
