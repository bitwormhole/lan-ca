<script lang="js">
import { ElButton } from 'element-plus';
import MyLoader from './SceneLoader.vue'
import MyAddingDialog from './SceneAddingDialog.vue'


export default {

    components: { MyLoader, MyAddingDialog },

    data() {
        return {
            displayAddingDialog: false,
            addingDialogTitle: '添加新场景',
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


        <MyAddingDialog v-model="displayAddingDialog" :title="addingDialogTitle" />


        <ElButton @click="fetch">Refresh</ElButton>
        <ElButton @click="displayAddingDialog = true">Add</ElButton>

    </div>
</template>
