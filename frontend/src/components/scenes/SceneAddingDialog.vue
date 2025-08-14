<script lang="js">
import { ElFormItem } from 'element-plus';
import { ElButton } from 'element-plus';
import { ElInput } from 'element-plus';
import { ElForm } from 'element-plus';
import { ElDialog } from 'element-plus';
import { useScenesStore } from '@/stores/scenes'

const scenes_store = useScenesStore()

export default {


    data() {
        return {
            name: '',
            label: '',
            comment: '',
        }
    },

    methods: {

        cancel() {
            this.$emit('update:modelValue', false);
        },

        commit() {
            let item = {}
            item.name = this.name;
            item.label = this.label;
            item.comment = this.comment;
            let vo = {
                scenes: [item]
            }
            scenes_store.insert(vo).then(() => {
                scenes_store.fetch()
                this.cancel()
            }).catch((err) => {
                let msg = err;
                alert(msg)
            })
        },
    }
}

</script>

<style lang="css"></style>

<template>
    <ElDialog>
        <ElForm>

            <ElFormItem label="名称">
                <ElInput v-model="name"></ElInput>
            </ElFormItem>
            <ElFormItem label="标签">
                <ElInput v-model="label"></ElInput>
            </ElFormItem>
            <ElFormItem label="备注">
                <ElInput v-model="comment"></ElInput>
            </ElFormItem>

            <ElFormItem>
                <ElButton type="primary" @click="commit">提交</ElButton>
                <ElButton @click="cancel">取消</ElButton>
            </ElFormItem>

        </ElForm>

    </ElDialog>
</template>
