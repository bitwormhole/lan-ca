<script lang="js">

import { useScenesStore } from '@/stores/scenes'
import { useSessionStore } from '@/stores/sessions'

const scenes_store = useScenesStore();
const session_store = useSessionStore();

export default {

    data() {
        return {}
    },

    computed: {
        currentUserUUID() {
            let info = session_store.info;
            return info.uuid;
        },
    },

    methods: {
        fetch() {
            scenes_store.fetch()
        },

        reset() {
            scenes_store.reset()
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

    watch: {
        currentUserUUID() {
            this.reset()
            this.fetch()
        },
    },
}
</script>

<style lang="css"></style>

<template>
    <div class="root"> </div>
</template>
