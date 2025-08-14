<script lang="js">

import { ElButton } from 'element-plus';

import { useSessionStore } from '@/stores/sessions'
import pinia from '@/stores/pinia'
import MyLoginDialog from '@/components/authx/LoginDialog.vue'


const session_store = useSessionStore(pinia);


export default {

    components: { MyLoginDialog },

    data() {
        return {
            displayLoginDialog: false,
        }
    },

    computed: {
        info() {
            return session_store.info;
        },

        authenticated() {
            return session_store.info.authenticated;
        },
    },

    methods: {
        fetch() {
            session_store.fetch()
        },

        handleClickLogin() {
            this.displayLoginDialog = true;
        },

        handleClickLogout() {
            session_store.exit();
        },
    },

    mounted() {
        this.fetch()
    },
}
</script>

<style lang="css"></style>

<template>
    <div class="session-info-box">
        <div class="session-info-1" v-show="!authenticated">
            <ElButton @click="handleClickLogin">登录</ElButton>
            <ElButton>注册</ElButton>
        </div>
        <div class="session-info-2" v-show="authenticated">
            <div> user_avatar: {{ info.avatar }} </div>
            <div> user_nick_name: {{ info.nickname }} </div>
            <div> user_id: {{ info.user }} </div>
            <div> user_uuid: {{ info.uuid }} </div>
            <div> user_email: {{ info.email }} </div>

            <ElButton @click="handleClickLogout">退登</ElButton>

        </div>

        <MyLoginDialog ref="theLoginDialog" v-model="displayLoginDialog" />
    </div>
</template>
