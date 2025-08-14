<script lang="js">

import UserFrameIndex from './user/index.vue'

export default {

    components: { UserFrameIndex },

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
        <UserFrameIndex>
            <slot default> </slot>
        </UserFrameIndex>
    </div>
</template>
