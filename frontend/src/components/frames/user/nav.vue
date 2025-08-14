<script lang="js">

import SessionInfoBox from '@/components/sessions/SessionInfoBox.vue'

export default {

    components: { SessionInfoBox },

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

        <SessionInfoBox />

        <ul>
            <li> <router-link to="/"> Home </router-link> </li>
            <li> <router-link to="/scenes"> scenes </router-link> </li>
            <li> <router-link to="/domains"> domains </router-link> </li>
            <li> <router-link to="/certificates"> certs </router-link> </li>
            <li> <router-link to="/about"> About </router-link> </li>
        </ul>


    </div>
</template>
