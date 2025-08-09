<script lang="js">

import { Time } from '@/js/time'
import MyCertUserInfoBox from './CertificateUserInfoBox.vue'


export default {

    components: { MyCertUserInfoBox },

    computed: {
        stringifyOfDomainNames() {
            let str = '';
            let list = this.item.domain_names;
            let sep = ''
            for (let i in list) {
                let dn = list[i];
                str = str + sep + dn;
                sep = ', '
            }
            return str;
        },
    },

    methods: {

        formatTime(t) {
            return Time.formatTime(t)
        },

    },

    props: {
        item: Object,
    },

}
</script>

<style lang="css">
.btn-submit-csr {
    margin-left: 20px;
}
</style>

<template>
    <ElDialog title="证书属性">
        <ElForm label-width="100">

            <ElFormItem label="域名"> {{ stringifyOfDomainNames }} </ElFormItem>

            <ElFormItem label="状态">
                {{ item.state }}
                <ElButton class="btn-submit-csr" type="primary">提交申请</ElButton>
            </ElFormItem>

            <ElFormItem label="标签"> {{ item.label }} </ElFormItem>
            <ElFormItem label="备注"> {{ item.comment }} </ElFormItem>



            <ElFormItem label="ID"> {{ item.id }} </ElFormItem>
            <ElFormItem label="UUID"> {{ item.uuid }} </ElFormItem>
            <ElFormItem label="序列号"> {{ item.sn }} </ElFormItem>
            <ElFormItem label="创建时间"> {{ formatTime(item.created_at) }} </ElFormItem>
            <ElFormItem label="更新时间"> {{ formatTime(item.updated_at) }} </ElFormItem>

            <ElFormItem label="签发者">
                <MyCertUserInfoBox v-model="item.signer" />
            </ElFormItem>
            <ElFormItem label="持有者">
                <MyCertUserInfoBox v-model="item.subject" />
            </ElFormItem>

            <ElFormItem label="证书指纹"></ElFormItem>
            <ElFormItem label="公钥指纹"></ElFormItem>

            <ElFormItem>
                <ElButton type="danger">删除</ElButton>
                <ElButton type="default">编辑</ElButton>
                <ElButton type="default">关闭</ElButton>
            </ElFormItem>

        </ElForm>
    </ElDialog>
</template>
