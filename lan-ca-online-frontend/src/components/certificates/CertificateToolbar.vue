<script lang="js">
import { ElButton } from 'element-plus';
import MyLoader from './CertificateLoader.vue'
import MyAddingDialog from './CertificateAddingDialog.vue'


export default {

    components: { MyLoader, MyAddingDialog },

    data() {
        return {
            displayAddingDialog: false,
        }
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
        <ElButton @click="displayAddingDialog = true"> Add </ElButton>


        <MyAddingDialog v-model="displayAddingDialog" />
    </div>
</template>
