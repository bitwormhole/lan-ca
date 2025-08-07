<script lang="js">
import { ElDialog } from 'element-plus';
import MyLoginBox from './LoginBox.vue'

export default {
    components: { MyLoginBox },

    methods: {
        handleLoginOK() {
            this.close();
        },

        close() {
            this.$emit('update:modelValue', false)
        },
    },

}

</script>

<style lang="css">
.login-dialog {
    max-width: 380px;
}
</style>

<template>
    <ElDialog class="login-dialog" title="登录">
        <MyLoginBox @login-ok="handleLoginOK" />
    </ElDialog>
</template>
