<script lang="js">

import MyNav from './nav.vue'

export default {

    components: { MyNav },

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
        <MyNav />
        <slot></slot>
    </div>
</template>
