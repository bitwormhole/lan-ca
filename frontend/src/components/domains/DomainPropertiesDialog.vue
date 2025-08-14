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

            <ElFormItem label="域名"> {{ item.name }} </ElFormItem>
            <ElFormItem label="使用场景"> {{ item.scene }} </ElFormItem>

            <ElFormItem label="解析到">
                <div>
                    <div> IPv4 地址: {{ item.ipv4_address }} </div>
                    <div> IPv6 地址: {{ item.ipv6_address }} </div>
                </div>
            </ElFormItem>

            <ElFormItem label="标签"> {{ item.label }} </ElFormItem>
            <ElFormItem label="备注"> {{ item.comment }} </ElFormItem>
            <ElFormItem label="详细描述"> {{ item.description }} </ElFormItem>



            <ElFormItem label="ID"> {{ item.id }} </ElFormItem>
            <ElFormItem label="UUID"> {{ item.uuid }} </ElFormItem>
            <ElFormItem label="创建时间"> {{ formatTime(item.created_at) }} </ElFormItem>
            <ElFormItem label="更新时间"> {{ formatTime(item.updated_at) }} </ElFormItem>


            <ElFormItem>
                <ElButton type="danger">删除</ElButton>
                <ElButton type="default">编辑</ElButton>
                <ElButton type="default">关闭</ElButton>
            </ElFormItem>

        </ElForm>
    </ElDialog>
</template>
