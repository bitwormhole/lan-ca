<template>
    <div class="view-root">


        {{ title }}

        <ElForm>

            <ElInput></ElInput>
            <ElInput></ElInput>

            <ElButton>创建新账号</ElButton>
        </ElForm>


    </div>
</template>

<style>
.view-root {
    border-width: 1px;
    border-color: gray;
    border-style: solid;
}
</style>

<script>
import { ElButton } from 'element-plus';
import { ElInput } from 'element-plus';
import { ElForm } from 'element-plus';


// import TheWelcome from '../components/TheWelcome.vue'

export default {

    data() {
        return {
            title: "Sign-Up"
        }
    }

}

</script>
