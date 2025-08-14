<script lang="js">

import { useAuthxStore } from '@/stores/authx'
import { useSessionStore } from '@/stores/sessions'


const theAuthxStore = useAuthxStore();
const theSessionStore = useSessionStore();


export default {

    data() {
        return {
            title: "Login",
            username: 'user@example.com',
            password: '666666',
        }
    },

    methods: {
        handleClickLogin() {
            let username = this.username;
            let password = this.password;
            theAuthxStore.login({ username, password }).then(() => {
                theSessionStore.fetch()
                this.$emit('login-ok')
            })
        },

        handleClickCreateNewUser() {
            let to = '/signup';
            let location = this.$router.resolve(to);
            let url = location.href;
            window.open(url, '_blank')
        },
    }

}

</script>

<style lang="css">
.view-root {
    border-width: 1px;
    border-color: gray;
    border-style: none;
}

.btn-login {
    width: 100%;
    margin-left: 0px;
    margin-right: 0px;
}

.btn-new-user {
    margin-top: 20px;
    margin-bottom: 30px;
    width: 100%;
    margin-left: auto;
    margin-right: auto;
}

.input-box {
    margin-bottom: 10px;
}
</style>

<template>
    <div class="view-root">
        <ElForm>

            <!-- <SessionInfoBox /> -->
            <!-- <h1>Login</h1> -->

            <ElInput class="input-box" v-model="username" placeholder="用户名 (Email 地址)" clearable></ElInput>
            <ElInput class="input-box" v-model="password" type="password" placeholder="密码" clearable></ElInput>

            <div>
                <ElButton class="btn-new-user" @click="handleClickCreateNewUser" link>注册新账号 >> </ElButton>
            </div>

            <div>
                <ElButton class="btn-login" type="primary" @click="handleClickLogin">登录</ElButton>
            </div>
        </ElForm>
    </div>
</template>
