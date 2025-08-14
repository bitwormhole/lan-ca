<script lang="js">

import { useScenesStore } from '@/stores/scenes'
import MyPropertiesDialog from './ScenePropertiesDialog.vue'



const scenes_store = useScenesStore();

export default {

    components: { MyPropertiesDialog },

    data() {
        return {
            displayPropertiesDialog: false,
            currentItem: {},
        }
    },

    computed: {
        items() {
            return scenes_store.items;
        },
    },


    methods: {
        handleClickItemProps(item) {
            // console.log("open props-dialog for item : " + item.id)
            this.currentItem = item;
            this.displayPropertiesDialog = true;
        },
    },

    mounted() {

    },
}
</script>

<style lang="css"></style>

<template>
    <div class="root">
        <el-table :data="items">
            <el-table-column prop="id" label="ID" />
            <el-table-column prop="name" label="名称" />
            <el-table-column prop="label" label="标签" />
            <el-table-column prop="comment" label="备注" />

            <el-table-column label="">
                <template #default="scope"> <el-button type="primary" link @click="handleClickItemProps(scope.row)">
                        属性 </el-button>
                </template>
            </el-table-column>

        </el-table>

        <MyPropertiesDialog v-model="displayPropertiesDialog" :item="currentItem" />
    </div>
</template>
