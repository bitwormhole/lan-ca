<script lang="js">
import { ElButton } from 'element-plus';
import MyLoader from './DomainLoader.vue'


export default {

    components: { MyLoader },

    data() {
        return {}
    },

    methods: {
        fetch() {
            this.$refs.loader.fetch();
        },
    },

    mounted() {

    },
}
</script>

<style lang="css"></style>

<template>
    <div class="root">
        <MyLoader ref="loader" />
        <ElButton @click="fetch"> Refresh </ElButton>
        <ElButton> Add </ElButton>
    </div>
</template>
