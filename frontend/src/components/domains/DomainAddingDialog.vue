<script lang="js">
import { ElInputNumber } from 'element-plus';
import { ElInput } from 'element-plus';
import { ElButton } from 'element-plus';
import MySceneSelector from '@/components/scenes/SceneSelector.vue'

import { useDomainsStore } from '@/stores/domains'


const theDomainsStore = useDomainsStore();


export default {

    components: { MySceneSelector },

    data() {
        return {
            domainNamePrefix: '',
            domainNameSuffix: '.user.lan-ca.space',
            ipv4: '',
            scene: 0, // the scene.id
            label: '',
            comment: '',
            description: '',
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

        makePostItem() {
            let item = {}
            item.scene = this.scene;
            item.name = this.domainNameFull;
            item.ipv4_address = this.ipv4;
            item.label = this.label;
            item.comment = this.comment;
            item.description = this.description;
            return item;
        },

        commit() {
            let item = this.makePostItem();
            let vo = {
                domains: [item]
            };
            theDomainsStore.insert(vo).then(() => {
                theDomainsStore.fetch()
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
    <ElDialog title="添加新域名">
        <ElForm label-width="100px">

            <ElFormItem label="使用场景">
                <MySceneSelector v-model="scene" />
            </ElFormItem>

            <ElFormItem label-width="0px">
                <div class="domain-name-full-box"> {{ domainNameFull }} </div>
            </ElFormItem>

            <ElFormItem label="域名">
                <ElInput v-model="domainNamePrefix" placeholder="请输入域名的前缀部分" clearable>
                    <template #append> {{ domainNameSuffix }} </template>
                </ElInput>
            </ElFormItem>

            <ElFormItem label="IPv4 地址">
                <ElInput v-model="ipv4" placeholder="请输入解析到的 IPv4 地址"></ElInput>
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
