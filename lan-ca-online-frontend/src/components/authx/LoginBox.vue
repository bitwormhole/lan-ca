<template>
    <div class="view-root">
        <ElForm>

            <SessionInfoBox />
            <h1>Login</h1>

            <ElInput v-model="username" placeholder="用户名 (Email 地址)"></ElInput>
            <ElInput v-model="password" type="password" placeholder="密码"></ElInput>
            <ElButton @click="handleClickLogin">登录</ElButton>
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

import { useAuthxStore } from '@/stores/authx'
import SessionInfoBox from '@/components/sessions/SessionInfoBox.vue'


const theAuthxStore = useAuthxStore();


export default {

    components: { SessionInfoBox },

    data() {
        return {
            title: "Login",
            username: 'user@example.com',
            password: '666666',
        }
    },

    methods: {
        handleClickLogin() {
            // useAuthxStore.login({}); 
            let username = this.username;
            let password = this.password;
            theAuthxStore.login({ username, password })
        },
    }

}

</script>
