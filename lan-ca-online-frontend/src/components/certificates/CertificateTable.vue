<script lang="js">

import { useCertificatesStore } from '@/stores/certificates'
import { Time } from '@/js/time'
import MyPropertiesDialog from './CertificatePropertiesDialog.vue'


const certs_store = useCertificatesStore();


const the_cert_state_table = {

    'initialled': '等待提交申请',
    'sub': '',
    'signed': '已签发',
    'deployed': '已部署',
    'expired': '已过期',

}

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
            return certs_store.items;
        },
    },

    methods: {

        handleClickItemProps(item) {
            this.currentItem = item;
            this.displayPropertiesDialog = true;
        },

        formatterForTime(row, column, cellValue, index) {
            return Time.formatTime(cellValue);
        },

        formatterForState(row, column, cellValue, index) {
            let table = the_cert_state_table;
            let text = table[cellValue];
            if (text == null) {
                text = 'N/A'
            }
            return text;
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

            <!-- <el-table-column prop="id" label="ID" /> -->
            <!-- <el-table-column prop="label" label="标签" /> -->

            <el-table-column prop="domain_names" label="目标域名" />

            <el-table-column prop="state" label="状态" :formatter="formatterForState" />

            <el-table-column prop="not_before" label="生效日期" :formatter="formatterForTime" />
            <el-table-column prop="not_after" label="失效日期" :formatter="formatterForTime" />

            <el-table-column prop="comment" label="备注" />

            <el-table-column label="">
                <template #default="scope"> <el-button type="primary" link @click="handleClickItemProps(scope.row)"> 属性
                    </el-button> </template>
            </el-table-column>
        </el-table>
        <MyPropertiesDialog v-model="displayPropertiesDialog" :item="currentItem" />
    </div>
</template>
