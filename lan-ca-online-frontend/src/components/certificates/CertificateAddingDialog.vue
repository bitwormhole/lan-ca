<script lang="js">
import { ElInputNumber } from 'element-plus';
import { ElInput } from 'element-plus';
import { ElButton } from 'element-plus';

import { useDomainsStore } from '@/stores/domains'
import { useCertificatesStore } from '@/stores/certificates'

import MyDomainSelector from '@/components/domains/DomainSelector.vue'


const theDomainsStore = useDomainsStore();
const theCertificatesStore = useCertificatesStore();


export default {

    components: { MyDomainSelector },

    data() {
        return {
            domain_name: '',
            label: '',
            comment: '',
            description: '',

            max_age: 12,

            subject_cn: '', // common-name
            subject_c: '', // 国家或地区
            subject_o: '', // 组织
            subject_ou: '', // 组织-部门
            subject_l: '', // 城市
            subject_st: '', // 省
            subject_email: '', // 邮件地址 
        }
    },

    computed: {
        domainNameFull() {
            let p1 = this.domainNamePrefix;
            let p2 = this.domainNameSuffix;
            return p1 + p2;
        },
    },

    methods: {

        computeNotAfter(not_before) {
            let ma = this.max_age; // max-age in month
            let ms = ma * (30 * 24 * 3600 * 1000);
            return not_before + ms;
        },

        makePostItem() {
            let dn = this.domain_name;
            let subject = {};
            let item = {} // a item as cert

            subject.cn = this.subject_cn;
            subject.c = this.subject_c;
            subject.st = this.subject_st;
            subject.l = this.subject_l;
            subject.o = this.subject_o;
            subject.ou = this.subject_ou;
            subject.email = this.subject_email;

            item.scene = this.scene;
            item.name = this.domainNameFull;
            item.ipv4_address = this.ipv4;
            item.label = this.label;
            item.comment = this.comment;
            item.description = this.description;
            item.domain_names = [dn]
            item.subject = subject;
            item.not_before = 0;
            item.not_after = this.computeNotAfter(0);

            return item;
        },

        commit() {
            let item = this.makePostItem();
            let vo = {
                certificates: [item]
            };
            theCertificatesStore.insert(vo).then(() => {
                theCertificatesStore.fetch()
                this.close();
            }).catch((err) => {
                let msg = err;
                alert(msg);
            });
        },

        cancel() {
            this.close();
        },

        close() {
            this.$emit('update:modelValue', false)
        },
    }

}
</script>

<style lang="css">
.domain-name-full-box {
    width: 100%;
    padding: 20px;
    background-color: gray;
    color: whitesmoke;
    text-align: center;
    margin-bottom: 0px;
}
</style>

<template>
    <ElDialog title="添加新证书">
        <ElForm label-width="100px">

            <ElFormItem label="域名">
                <MyDomainSelector v-model="domain_name" />
            </ElFormItem>

            <ElFormItem label="有效期">
                <el-input-number v-model="max_age" :min="1" :max="120" placeholder="请输入有效期">
                </el-input-number> <span style="margin-left: 10px ;"> 月 </span>
            </ElFormItem>

            <ElFormItem label="常用名">
                <ElInput v-model="subject_cn" placeholder="CN"></ElInput>
            </ElFormItem>
            <ElFormItem label="国家/地区">
                <ElInput v-model="subject_c" placeholder="C"></ElInput>
            </ElFormItem>
            <ElFormItem label="省/自治区">
                <ElInput v-model="subject_st" placeholder="ST"></ElInput>
            </ElFormItem>
            <ElFormItem label="市/县">
                <ElInput v-model="subject_l" placeholder="L"></ElInput>
            </ElFormItem>
            <ElFormItem label="组织">
                <ElInput v-model="subject_o" placeholder="O"></ElInput>
            </ElFormItem>
            <ElFormItem label="组织部门">
                <ElInput v-model="subject_ou" placeholder="OU"></ElInput>
            </ElFormItem>
            <ElFormItem label="邮件地址">
                <ElInput v-model="subject_email" placeholder="Email Address"></ElInput>
            </ElFormItem>

            <ElFormItem label="标签">
                <ElInput v-model="label" placeholder=""></ElInput>
            </ElFormItem>
            <ElFormItem label="备注">
                <ElInput v-model="comment" placeholder=""></ElInput>
            </ElFormItem>
            <ElFormItem label="详细描述">
                <ElInput v-model="description" type="textarea" placeholder=""></ElInput>
            </ElFormItem>

            <ElFormItem>
                <ElButton type="primary" @click="commit">提交</ElButton>
                <ElButton type="default" @click="cancel">取消</ElButton>
            </ElFormItem>

        </ElForm>
    </ElDialog>
</template>
