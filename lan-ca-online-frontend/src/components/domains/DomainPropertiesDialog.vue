<script lang="js">


import { Time } from '@/js/time'


export default {
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

<style lang="css"></style>

<template>
    <ElDialog title="域名属性">
        <ElForm label-width="100">

            <ElFormItem></ElFormItem>
            <ElFormItem></ElFormItem>

            <ElFormItem label="ID"> {{ item.id }} </ElFormItem>
            <ElFormItem label="UUID"> {{ item.uuid }} </ElFormItem>
            <ElFormItem label="创建时间"> {{ formatTime(item.created_at) }} </ElFormItem>
            <ElFormItem label="更新时间"> {{ formatTime(item.updated_at) }} </ElFormItem>

            <ElFormItem></ElFormItem>
            <ElFormItem></ElFormItem>

            <ElFormItem>
                <ElButton type="danger">删除</ElButton>
                <ElButton type="default">编辑</ElButton>
                <ElButton type="default">关闭</ElButton>
            </ElFormItem>

        </ElForm>
    </ElDialog>
</template>
